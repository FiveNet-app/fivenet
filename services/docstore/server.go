package docstore

import (
	"database/sql"

	"github.com/fivenet-app/fivenet/gen/go/proto/resources/documents"
	pbdocstore "github.com/fivenet-app/fivenet/gen/go/proto/services/docstore"
	"github.com/fivenet-app/fivenet/pkg/access"
	"github.com/fivenet-app/fivenet/pkg/grpc/auth/userinfo"
	"github.com/fivenet-app/fivenet/pkg/housekeeper"
	"github.com/fivenet-app/fivenet/pkg/html/htmldiffer"
	"github.com/fivenet-app/fivenet/pkg/mstlystcdata"
	"github.com/fivenet-app/fivenet/pkg/notifi"
	"github.com/fivenet-app/fivenet/pkg/perms"
	"github.com/fivenet-app/fivenet/pkg/server/audit"
	"github.com/fivenet-app/fivenet/query/fivenet/table"
	"go.uber.org/fx"
	grpc "google.golang.org/grpc"
)

func init() {
	housekeeper.AddTable(&housekeeper.Table{
		Table:           table.FivenetDocuments,
		TimestampColumn: table.FivenetDocuments.DeletedAt,
		MinDays:         housekeeperMinDays,
	})

	housekeeper.AddTable(&housekeeper.Table{
		Table:           table.FivenetDocumentsTemplates,
		TimestampColumn: table.FivenetDocumentsTemplates.DeletedAt,
		MinDays:         housekeeperMinDays,
	})

	housekeeper.AddTable(&housekeeper.Table{
		Table:           table.FivenetDocumentsComments,
		TimestampColumn: table.FivenetDocumentsComments.DeletedAt,
		MinDays:         housekeeperMinDays,
	})

	housekeeper.AddTable(&housekeeper.Table{
		Table:           table.FivenetDocumentsReferences,
		TimestampColumn: table.FivenetDocumentsReferences.DeletedAt,
		MinDays:         housekeeperMinDays,
	})
	housekeeper.AddTable(&housekeeper.Table{
		Table:           table.FivenetDocumentsRelations,
		TimestampColumn: table.FivenetDocumentsRelations.DeletedAt,
		MinDays:         housekeeperMinDays,
	})
}

type Server struct {
	pbdocstore.DocStoreServiceServer

	db       *sql.DB
	ps       perms.Permissions
	cache    *mstlystcdata.Cache
	enricher *mstlystcdata.UserAwareEnricher
	aud      audit.IAuditer
	ui       userinfo.UserInfoRetriever
	notif    notifi.INotifi
	htmlDiff *htmldiffer.Differ

	access         *access.Grouped[documents.DocumentJobAccess, *documents.DocumentJobAccess, documents.DocumentUserAccess, *documents.DocumentUserAccess, access.DummyQualificationAccess[documents.AccessLevel], *access.DummyQualificationAccess[documents.AccessLevel], documents.AccessLevel]
	templateAccess *access.Grouped[documents.TemplateJobAccess, *documents.TemplateJobAccess, documents.TemplateUserAccess, *documents.TemplateUserAccess, access.DummyQualificationAccess[documents.AccessLevel], *access.DummyQualificationAccess[documents.AccessLevel], documents.AccessLevel]
}

type Params struct {
	fx.In

	DB         *sql.DB
	Perms      perms.Permissions
	Cache      *mstlystcdata.Cache
	Enricher   *mstlystcdata.UserAwareEnricher
	Aud        audit.IAuditer
	Ui         userinfo.UserInfoRetriever
	Notif      notifi.INotifi
	HTMLDiffer *htmldiffer.Differ
}

func NewServer(p Params) *Server {
	return &Server{
		db:       p.DB,
		ps:       p.Perms,
		cache:    p.Cache,
		enricher: p.Enricher,
		aud:      p.Aud,
		ui:       p.Ui,
		notif:    p.Notif,
		htmlDiff: p.HTMLDiffer,

		access: newAccess(p.DB),

		templateAccess: access.NewGrouped[documents.TemplateJobAccess, *documents.TemplateJobAccess, documents.TemplateUserAccess, *documents.TemplateUserAccess, access.DummyQualificationAccess[documents.AccessLevel], *access.DummyQualificationAccess[documents.AccessLevel], documents.AccessLevel](
			p.DB,
			table.FivenetDocumentsTemplates,
			&access.TargetTableColumns{
				ID:         table.FivenetDocumentsTemplates.ID,
				DeletedAt:  table.FivenetDocumentsTemplates.DeletedAt,
				CreatorID:  nil,
				CreatorJob: table.FivenetDocumentsTemplates.CreatorJob,
			},
			access.NewJobs[documents.TemplateJobAccess, *documents.TemplateJobAccess, documents.AccessLevel](
				table.FivenetDocumentsTemplatesJobAccess,
				&access.JobAccessColumns{
					BaseAccessColumns: access.BaseAccessColumns{
						ID:        table.FivenetDocumentsTemplatesJobAccess.ID,
						CreatedAt: table.FivenetDocumentsTemplatesJobAccess.CreatedAt,
						TargetID:  table.FivenetDocumentsTemplatesJobAccess.TemplateID,
						Access:    table.FivenetDocumentsTemplatesJobAccess.Access,
					},
					Job:          table.FivenetDocumentsTemplatesJobAccess.Job,
					MinimumGrade: table.FivenetDocumentsTemplatesJobAccess.MinimumGrade,
				},
				table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access"),
				&access.JobAccessColumns{
					BaseAccessColumns: access.BaseAccessColumns{
						ID:        table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").ID,
						CreatedAt: table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").CreatedAt,
						TargetID:  table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").TemplateID,
						Access:    table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").Access,
					},
					Job:          table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").Job,
					MinimumGrade: table.FivenetDocumentsTemplatesJobAccess.AS("template_job_access").MinimumGrade,
				},
			),
			nil,
			nil,
		),
	}
}

func newAccess(db *sql.DB) *access.Grouped[documents.DocumentJobAccess, *documents.DocumentJobAccess, documents.DocumentUserAccess, *documents.DocumentUserAccess, access.DummyQualificationAccess[documents.AccessLevel], *access.DummyQualificationAccess[documents.AccessLevel], documents.AccessLevel] {
	return access.NewGrouped[documents.DocumentJobAccess, *documents.DocumentJobAccess, documents.DocumentUserAccess, *documents.DocumentUserAccess, access.DummyQualificationAccess[documents.AccessLevel], *access.DummyQualificationAccess[documents.AccessLevel], documents.AccessLevel](
		db,
		table.FivenetDocuments,
		&access.TargetTableColumns{
			ID:         table.FivenetDocuments.ID,
			DeletedAt:  table.FivenetDocuments.DeletedAt,
			CreatorID:  table.FivenetDocuments.CreatorID,
			CreatorJob: table.FivenetDocuments.CreatorJob,
		},
		access.NewJobs[documents.DocumentJobAccess, *documents.DocumentJobAccess, documents.AccessLevel](
			table.FivenetDocumentsJobAccess,
			&access.JobAccessColumns{
				BaseAccessColumns: access.BaseAccessColumns{
					ID:        table.FivenetDocumentsJobAccess.ID,
					CreatedAt: table.FivenetDocumentsJobAccess.CreatedAt,
					TargetID:  table.FivenetDocumentsJobAccess.DocumentID,
					Access:    table.FivenetDocumentsJobAccess.Access,
				},
				Job:          table.FivenetDocumentsJobAccess.Job,
				MinimumGrade: table.FivenetDocumentsJobAccess.MinimumGrade,
			},
			table.FivenetDocumentsJobAccess.AS("document_job_access"),
			&access.JobAccessColumns{
				BaseAccessColumns: access.BaseAccessColumns{
					ID:        table.FivenetDocumentsJobAccess.AS("document_job_access").ID,
					CreatedAt: table.FivenetDocumentsJobAccess.AS("document_job_access").CreatedAt,
					TargetID:  table.FivenetDocumentsJobAccess.AS("document_job_access").DocumentID,
					Access:    table.FivenetDocumentsJobAccess.AS("document_job_access").Access,
				},
				Job:          table.FivenetDocumentsJobAccess.AS("document_job_access").Job,
				MinimumGrade: table.FivenetDocumentsJobAccess.AS("document_job_access").MinimumGrade,
			},
		),
		access.NewUsers[documents.DocumentUserAccess, *documents.DocumentUserAccess, documents.AccessLevel](
			table.FivenetDocumentsUserAccess,
			&access.UserAccessColumns{
				BaseAccessColumns: access.BaseAccessColumns{
					ID:        table.FivenetDocumentsUserAccess.ID,
					CreatedAt: table.FivenetDocumentsUserAccess.CreatedAt,
					TargetID:  table.FivenetDocumentsUserAccess.DocumentID,
					Access:    table.FivenetDocumentsUserAccess.Access,
				},
				UserId: table.FivenetDocumentsUserAccess.UserID,
			},
			table.FivenetDocumentsUserAccess.AS("document_user_access"),
			&access.UserAccessColumns{
				BaseAccessColumns: access.BaseAccessColumns{
					ID:        table.FivenetDocumentsUserAccess.AS("document_user_access").ID,
					CreatedAt: table.FivenetDocumentsUserAccess.AS("document_user_access").CreatedAt,
					TargetID:  table.FivenetDocumentsUserAccess.AS("document_user_access").DocumentID,
					Access:    table.FivenetDocumentsUserAccess.AS("document_user_access").Access,
				},
				UserId: table.FivenetDocumentsUserAccess.AS("document_user_access").UserID,
			},
		),
		nil,
	)
}

func (s *Server) RegisterServer(srv *grpc.Server) {
	pbdocstore.RegisterDocStoreServiceServer(srv, s)
}

func (s *Server) GetPermsRemap() map[string]string {
	return pbdocstore.PermsRemap
}