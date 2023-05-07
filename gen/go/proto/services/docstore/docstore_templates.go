package docstore

import (
	"bytes"
	context "context"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/galexrt/fivenet/gen/go/proto/resources/documents"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	dTemplates          = table.FivenetDocumentsTemplates.AS("templateshort")
	dTemplatesJobAccess = table.FivenetDocumentsTemplatesJobAccess.AS("templatejobaccess")
)

func (s *Server) ListTemplates(ctx context.Context, req *ListTemplatesRequest) (*ListTemplatesResponse, error) {
	_, job, jobGrade := auth.GetUserInfoFromContext(ctx)

	stmt := dTemplates.
		SELECT(
			dTemplates.ID,
			dCategory.ID,
			dCategory.Name,
			dCategory.Description,
			dCategory.Job,
			dTemplates.Title,
			dTemplates.Description,
			dTemplates.Schema,
			dTemplates.CreatorID,
		).
		FROM(
			dTemplates.
				INNER_JOIN(dTemplatesJobAccess,
					dTemplatesJobAccess.TemplateID.EQ(dTemplates.ID).
						AND(dTemplatesJobAccess.Job.EQ(jet.String(job))).
						AND(dTemplatesJobAccess.MinimumGrade.LT_EQ(jet.Int32(jobGrade))),
				).
				LEFT_JOIN(dCategory,
					dCategory.ID.EQ(dTemplates.CategoryID),
				),
		).
		WHERE(
			dTemplates.DeletedAt.IS_NULL(),
		).
		GROUP_BY(dTemplates.ID)

	fmt.Println(stmt.DebugSql())

	resp := &ListTemplatesResponse{}
	if err := stmt.QueryContext(ctx, s.db, &resp.Templates); err != nil {
		if !errors.Is(err, qrm.ErrNoRows) {
			return nil, err
		}
	}

	return resp, nil
}

func (s *Server) GetTemplate(ctx context.Context, req *GetTemplateRequest) (*GetTemplateResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)

	check, err := s.checkIfUserHasAccessToTemplate(ctx, req.TemplateId, userId, job, jobGrade, false, documents.ACCESS_LEVEL_VIEW)
	if err != nil {
		return nil, FailedQueryErr
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to view this template!")
	}

	dTemplates := dTemplates.AS("template")
	stmt := dTemplates.
		SELECT(
			dTemplates.ID,
			dTemplates.CreatedAt,
			dTemplates.UpdatedAt,
			dCategory.ID,
			dCategory.Name,
			dCategory.Description,
			dCategory.Job,
			dTemplates.Title,
			dTemplates.Description,
			dTemplates.ContentTitle,
			dTemplates.Content,
			dTemplates.Access,
			dTemplates.Schema,
			dTemplates.CreatorID,
			dTemplates.CreatorJob,
		).
		FROM(
			dTemplates.
				LEFT_JOIN(dCategory,
					dCategory.ID.EQ(dTemplates.CategoryID),
				),
		).
		WHERE(
			dTemplates.ID.EQ(jet.Uint64(req.TemplateId)),
		).
		LIMIT(1)

	fmt.Println(stmt.DebugSql())

	resp := &GetTemplateResponse{}
	if err := stmt.QueryContext(ctx, s.db, resp); err != nil {
		return nil, FailedQueryErr
	}

	if req.Render == nil || !*req.Render {
		resp.Template.JobAccess, err = s.getTemplateJobAccess(ctx, req.TemplateId)
		if err != nil {
			return nil, FailedQueryErr
		}
	} else if req.Render != nil && *req.Render {
		// Parse data as json for the templating process
		var data map[string]interface{}
		err := json.Unmarshal([]byte(req.Data), &data)
		if err != nil {
			return nil, err
		}

		resp.Template.ContentTitle, resp.Template.Content, err = s.renderTemplate(resp.Template, data)
		if err != nil {
			return nil, err
		}

		resp.Rendered = true
	}

	return resp, nil
}

func (s *Server) renderTemplate(docTmpl *documents.Template, data map[string]interface{}) (outTile string, out string, err error) {
	// Render Title template
	tmplTitle, err := template.
		New("title").
		Funcs(sprig.FuncMap()).
		Parse(docTmpl.ContentTitle)
	if err != nil {
		return
	}
	buf := &bytes.Buffer{}
	err = tmplTitle.Execute(buf, data)
	if err != nil {
		return
	}
	outTile = buf.String()

	// Render Content Template
	tmpl, err := template.
		New("content").
		Funcs(sprig.FuncMap()).
		Parse(docTmpl.Content)
	if err != nil {
		return
	}

	buf.Reset()
	err = tmpl.Execute(buf, data)
	if err != nil {
		return
	}
	out = buf.String()

	return
}

func (s *Server) CreateTemplate(ctx context.Context, req *CreateTemplateRequest) (*CreateTemplateResponse, error) {
	userId, job, _ := auth.GetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "CreateTemplate",
		UserID:  userId,
		UserJob: job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FailedQueryErr
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	req.Template.Job = job

	categoryId := jet.NULL
	if req.Template.Category != nil {
		cat, err := s.getDocumentCategory(ctx, req.Template.Category.Id)
		if err != nil {
			return nil, err
		}
		if cat != nil {
			categoryId = jet.Uint64(cat.Id)
		}
	}

	dTemplates := table.FivenetDocumentsTemplates
	stmt := dTemplates.
		INSERT(
			dTemplates.CategoryID,
			dTemplates.Title,
			dTemplates.Description,
			dTemplates.ContentTitle,
			dTemplates.Content,
			dTemplates.Access,
			dTemplates.Schema,
			dTemplates.CreatorID,
			dTemplates.CreatorJob,
		).
		VALUES(
			categoryId,
			req.Template.Title,
			req.Template.Description,
			req.Template.ContentTitle,
			req.Template.Content,
			req.Template.ContentAccess,
			req.Template.Schema,
			userId,
			job,
		)

	res, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err := s.handleTemplateAccessChanges(ctx, tx, uint64(lastId), req.Template.JobAccess); err != nil {
		return nil, FailedQueryErr
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, FailedQueryErr
	}

	auditEntry.State = int16(rector.EVENT_TYPE_CREATED)

	return &CreateTemplateResponse{
		Id: uint64(lastId),
	}, nil
}

func (s *Server) UpdateTemplate(ctx context.Context, req *UpdateTemplateRequest) (*UpdateTemplateResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "UpdateTemplate",
		UserID:  userId,
		UserJob: job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	check, err := s.checkIfUserHasAccessToTemplate(ctx, req.Template.Id, userId, job, jobGrade, false, documents.ACCESS_LEVEL_EDIT)
	if err != nil {
		return nil, FailedQueryErr
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to edit this template!")
	}

	categoryId := jet.NULL
	if req.Template.Category != nil {
		cat, err := s.getDocumentCategory(ctx, req.Template.Category.Id)
		if err != nil {
			return nil, err
		}
		if cat != nil {
			categoryId = jet.Uint64(cat.Id)
		}
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, FailedQueryErr
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	dTemplates := table.FivenetDocumentsTemplates
	stmt := dTemplates.
		UPDATE(
			dTemplates.CategoryID,
			dTemplates.Title,
			dTemplates.Description,
			dTemplates.ContentTitle,
			dTemplates.Content,
			dTemplates.Access,
			dTemplates.Schema,
		).
		SET(
			categoryId,
			req.Template.Title,
			req.Template.Description,
			req.Template.ContentTitle,
			req.Template.Content,
			req.Template.ContentAccess,
			req.Template.Schema,
		).
		WHERE(
			dTemplates.ID.EQ(jet.Uint64(req.Template.Id)),
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return nil, FailedQueryErr
	}

	if err := s.handleTemplateAccessChanges(ctx, tx, uint64(req.Template.Id), req.Template.JobAccess); err != nil {
		return nil, FailedQueryErr
	}

	// Commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, FailedQueryErr
	}

	auditEntry.State = int16(rector.EVENT_TYPE_UPDATED)

	return &UpdateTemplateResponse{
		Id: req.Template.Id,
	}, nil
}

func (s *Server) DeleteTemplate(ctx context.Context, req *DeleteTemplateRequest) (*DeleteTemplateResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "DeleteTemplate",
		UserID:  userId,
		UserJob: job,
		State:   int16(rector.EVENT_TYPE_ERRORED),
	}
	defer s.a.AddEntryWithData(auditEntry, req)

	check, err := s.checkIfUserHasAccessToTemplate(ctx, req.Id, userId, job, jobGrade, false, documents.ACCESS_LEVEL_EDIT)
	if err != nil {
		return nil, FailedQueryErr
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to delete this template!")
	}

	dTemplates := table.FivenetDocumentsTemplates
	stmt := dTemplates.
		DELETE().
		WHERE(
			dTemplates.CreatorJob.EQ(jet.String(job)).AND(
				dTemplates.ID.EQ(jet.Uint64(req.Id))),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	auditEntry.State = int16(rector.EVENT_TYPE_DELETED)

	return &DeleteTemplateResponse{}, nil
}
