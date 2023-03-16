package livemapper

import (
	"database/sql"
	"math/rand"
	"time"

	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/pkg/perms"
	"github.com/galexrt/arpanet/proto/resources/livemap"
	"github.com/galexrt/arpanet/query/arpanet/model"
	"github.com/galexrt/arpanet/query/arpanet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"go.uber.org/zap"
)

var (
	l = table.ArpanetUserLocations
	u = table.Users.AS("user")
)

type Server struct {
	LivemapperServiceServer

	logger *zap.Logger
	db     *sql.DB
	p      perms.Permissions
}

func NewServer(logger *zap.Logger, db *sql.DB, p perms.Permissions) *Server {
	return &Server{
		logger: logger,
		db:     db,
		p:      p,
	}
}

func (s *Server) SetLogger(logger *zap.Logger) {
	s.logger = logger
}

func (s *Server) Stream(req *StreamRequest, srv LivemapperService_StreamServer) error {
	userId := auth.GetUserIDFromContext(srv.Context())

	jobs, err := s.p.GetSuffixOfPermissionsByPrefixOfUser(userId, "LivemapperService.Stream")
	if err != nil {
		return err
	}

	if len(jobs) == 0 {
		return nil
	}

	sqlJobs := make([]jet.Expression, len(jobs))
	for k := 0; k < len(jobs); k++ {
		sqlJobs[k] = jet.String(jobs[k])
	}

	l := l.AS("usermarker")
	stmt := l.SELECT(
		l.UserID,
		l.Job,
		l.X,
		l.Y,
		l.UpdatedAt,
		u.ID,
		u.Identifier,
		u.Job,
		u.JobGrade,
		u.Firstname,
		u.Lastname,
	).
		FROM(
			l.
				LEFT_JOIN(u,
					l.UserID.EQ(u.ID),
				),
		).
		WHERE(
			l.Job.IN(sqlJobs...).
				AND(
					l.Hidden.IS_FALSE(),
				).
				AND(
					l.UpdatedAt.GT_EQ(jet.CURRENT_TIMESTAMP().SUB(jet.INTERVAL(5, jet.MINUTE))),
				),
		)

	for {
		resp := &ServerStreamResponse{
			Dispatches: []*livemap.DispatchMarker{},
		}

		if err := stmt.QueryContext(srv.Context(), s.db, &resp.Users); err != nil {
			return err
		}

		if err := srv.Send(resp); err != nil {
			return err
		}
		time.Sleep(3 * time.Second)
	}
}

func (s *Server) GenerateRandomUserMarker() {
	userIds := []int32{
		// ambulance
		26061,
		4650,
		29225,
		29205,
		931,
		16235,
		6173,
		20232,
		1634,
		17800,
		27434,
		15706,
		3046,
	}

	for {
		markers := make([]*model.ArpanetUserLocations, len(userIds))
		for i := 0; i < len(userIds); i++ {
			xMin := -3500
			xMax := 4300
			x := float64(rand.Intn(xMax-xMin+1) + xMin)

			yMin := -3600
			yMax := 7800
			y := float64(rand.Intn(yMax-yMin+1) + yMin)

			job := "ambulance"
			hidden := false
			markers[i] = &model.ArpanetUserLocations{
				UserID: userIds[i],
				Job:    &job,
				Hidden: &hidden,

				X: &x,
				Y: &y,
			}
		}

		stmt := l.INSERT(
			l.UserID,
			l.Job,
			l.X,
			l.Y,
			l.Hidden,
		).
			MODELS(markers).
			ON_DUPLICATE_KEY_UPDATE(
				l.X.SET(jet.RawFloat("VALUES(x)")),
				l.Y.SET(jet.RawFloat("VALUES(y)")),
			)

		_, err := stmt.Exec(s.db)
		if err != nil {
			s.logger.Error("failed to insert/ update random location to locations table", zap.Error(err))
		}

		time.Sleep(2 * time.Second)
	}
}
