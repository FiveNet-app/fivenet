package mstlystcdata

import (
	"context"
	"database/sql"
	"sort"
	"strings"
	"time"

	cache "github.com/Code-Hex/go-generics-cache"
	"github.com/galexrt/fivenet/gen/go/proto/resources/documents"
	"github.com/galexrt/fivenet/gen/go/proto/resources/laws"
	"github.com/galexrt/fivenet/gen/go/proto/resources/users"
	"github.com/galexrt/fivenet/pkg/config"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/puzpuzpuz/xsync/v2"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var (
	tJobs      = table.Jobs.AS("job")
	tJGrades   = table.JobGrades.AS("jobgrade")
	tDCategory = table.FivenetDocumentsCategories.AS("category")
	tLawBooks  = table.FivenetLawbooks.AS("lawbook")
	tLaws      = table.FivenetLawbooksLaws.AS("law")
)

type Cache struct {
	logger *zap.Logger
	db     *sql.DB

	refreshTime time.Duration

	tracer             trace.Tracer
	ctx                context.Context
	jobs               *cache.Cache[string, *users.Job]
	docCategories      *cache.Cache[uint64, *documents.Category]
	docCategoriesByJob *cache.Cache[string, []*documents.Category]
	lawBooks           *xsync.MapOf[uint64, *laws.LawBook]

	searcher *Searcher
}

type Params struct {
	fx.In

	LC fx.Lifecycle

	Logger *zap.Logger
	TP     *tracesdk.TracerProvider
	DB     *sql.DB
	Config *config.Config
}

func NewCache(p Params) (*Cache, error) {
	ctx, cancel := context.WithCancel(context.Background())

	c := &Cache{
		logger: p.Logger,
		db:     p.DB,

		refreshTime: p.Config.Cache.RefreshTime,

		tracer:             p.TP.Tracer("mstlystcdata-cache"),
		ctx:                ctx,
		jobs:               cache.NewContext[string, *users.Job](ctx),
		docCategories:      cache.NewContext[uint64, *documents.Category](ctx),
		docCategoriesByJob: cache.NewContext[string, []*documents.Category](ctx),
		lawBooks:           xsync.NewIntegerMapOf[uint64, *laws.LawBook](),
	}

	var err error
	c.searcher, err = NewSearcher(c)
	c.searcher.addDataToIndex()

	p.LC.Append(fx.StartHook(func(_ context.Context) error {
		if err := c.refreshCache(); err != nil {
			return err
		}

		go c.start()
		return nil
	}))

	p.LC.Append(fx.StopHook(func(_ context.Context) error {
		cancel()
		return nil
	}))

	return c, err
}

func (c *Cache) start() {
	for {
		if err := c.refreshCache(); err != nil {
			c.logger.Error("failed to refresh mostly static data cache", zap.Error(err))
		}

		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.refreshTime):
		}
	}
}

func (c *Cache) GetSearcher() *Searcher {
	return c.searcher
}

func (c *Cache) refreshCache() error {
	ctx, span := c.tracer.Start(c.ctx, "mstlystcdata-refresh-cache")
	defer span.End()

	if err := c.refreshCategories(ctx); err != nil {
		return err
	}

	if err := c.refreshJobs(ctx); err != nil {
		return err
	}

	if err := c.RefreshLaws(ctx, 0); err != nil {
		return err
	}

	if c.searcher != nil {
		c.searcher.addDataToIndex()
	}

	return nil
}

func (c *Cache) refreshCategories(ctx context.Context) error {
	stmt := tDCategory.
		SELECT(
			tDCategory.ID,
			tDCategory.Name,
			tDCategory.Description,
			tDCategory.Job,
		).
		FROM(tDCategory).
		ORDER_BY(
			tDCategory.Job.ASC(),
			tDCategory.Name.ASC(),
		)

	var dest []*documents.Category
	if err := stmt.QueryContext(ctx, c.db, &dest); err != nil {
		return err
	}

	categoriesPerJob := map[string][]*documents.Category{}
	for _, d := range dest {
		c.docCategories.Set(d.Id, d)

		if _, ok := categoriesPerJob[*d.Job]; !ok {
			categoriesPerJob[*d.Job] = []*documents.Category{}
		}
		categoriesPerJob[*d.Job] = append(categoriesPerJob[*d.Job], d)
	}

	// Update cache
	for job, cs := range categoriesPerJob {
		c.docCategoriesByJob.Set(job, cs)
	}

	return nil
}

func (c *Cache) refreshJobs(ctx context.Context) error {
	stmt := tJobs.
		SELECT(
			tJobs.Name,
			tJobs.Label,
			tJGrades.JobName.AS("job_grade.job_name"),
			tJGrades.Grade,
			tJGrades.Label,
		).
		FROM(tJobs.
			INNER_JOIN(tJGrades,
				tJGrades.JobName.EQ(tJobs.Name),
			),
		).
		ORDER_BY(
			tJobs.Name.ASC(),
			tJGrades.Grade.ASC(),
		)

	var dest []*users.Job
	if err := stmt.QueryContext(ctx, c.db, &dest); err != nil {
		return err
	}

	// Update cache
	for _, job := range dest {
		c.jobs.Set(strings.ToLower(job.Name), job)
	}

	return nil
}

func (c *Cache) RefreshLaws(ctx context.Context, lawBookId uint64) error {
	stmt := tLawBooks.
		SELECT(
			tLawBooks.ID,
			tLawBooks.CreatedAt,
			tLawBooks.UpdatedAt,
			tLawBooks.Name,
			tLawBooks.Description,
			tLaws.ID,
			tLaws.LawbookID,
			tLaws.CreatedAt,
			tLaws.UpdatedAt,
			tLaws.Name,
			tLaws.Description,
			tLaws.Fine,
			tLaws.DetentionTime,
			tLaws.StvoPoints,
		).
		FROM(tLawBooks.
			INNER_JOIN(tLaws,
				tLaws.LawbookID.EQ(tLawBooks.ID),
			),
		).
		ORDER_BY(
			tLawBooks.Name.ASC(),
			tLaws.Name.ASC(),
		)

	if lawBookId != 0 {
		stmt = stmt.WHERE(
			tLawBooks.ID.EQ(jet.Uint64(lawBookId)),
		)
	}

	var dest []*laws.LawBook
	if err := stmt.QueryContext(ctx, c.db, &dest); err != nil {
		return err
	}

	// Lawbook not found, need to remove it
	if lawBookId != 0 && len(dest) == 0 {
		c.lawBooks.Delete(lawBookId)
	} else {
		// Update cache
		found := []uint64{}
		for _, lawbook := range dest {
			c.lawBooks.Store(lawbook.Id, lawbook)
			found = append(found, lawbook.Id)
		}

		// Delete non-existing law books, based on which are in the database
		c.lawBooks.Range(func(key uint64, value *laws.LawBook) bool {
			if !utils.InSliceFunc(found, func(in uint64) bool {
				return in == key
			}) {
				c.lawBooks.Delete(key)
			}
			return true
		})
	}

	return nil
}

func (c *Cache) GetLawBooks() []*laws.LawBook {
	lawBooks := []*laws.LawBook{}
	c.lawBooks.Range(func(key uint64, value *laws.LawBook) bool {
		lawBooks = append(lawBooks, value)
		return true
	})

	sort.Slice(lawBooks, func(i, j int) bool {
		return lawBooks[i].Name < lawBooks[j].Name
	})

	return lawBooks
}

func (c *Cache) GetHighestJobGrade(job string) *users.JobGrade {
	j, ok := c.jobs.Get(job)
	if !ok {
		return nil
	}

	if len(j.Grades) == 0 {
		return nil
	}

	return j.Grades[len(j.Grades)-1]
}
