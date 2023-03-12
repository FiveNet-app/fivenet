package docstore

import (
	context "context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/pkg/perms"
	database "github.com/galexrt/arpanet/proto/resources/common/database"
	"github.com/galexrt/arpanet/proto/resources/documents"
	"github.com/galexrt/arpanet/query"
	"github.com/galexrt/arpanet/query/arpanet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/microcosm-cc/bluemonday"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		{Key: "documents", Name: "View"},
		{Key: "documents", Name: "FindDocuments"},
		{Key: "documents", Name: "GetDocument"},
		{Key: "documents", Name: "CreateDocument"},
		{Key: "documents", Name: "UpdateDocument"},
		{Key: "documents", Name: "CompleteCategories", PerJob: true},
	})
}

var (
	u   = table.Users
	d   = table.ArpanetDocuments.AS("document")
	dua = table.ArpanetDocumentsUserAccess
	dja = table.ArpanetDocumentsJobAccess
	dt  = table.ArpanetDocumentsTemplates
)

type Server struct {
	DocStoreServiceServer

	p *bluemonday.Policy
}

func NewServer() *Server {
	return &Server{
		p: bluemonday.UGCPolicy(),
	}
}

func (s *Server) getDocumentsQuery(where jet.BoolExpression, onlyColumns jet.ProjectionList, additionalColumns jet.ProjectionList, userID int32, job string, jobGrade int32) jet.SelectStatement {
	wheres := []jet.BoolExpression{jet.OR(
		jet.OR(
			d.Public.IS_TRUE(),
			d.CreatorID.EQ(jet.Int32(userID)),
		),
		jet.OR(
			jet.AND(
				dua.Access.IS_NOT_NULL(),
				dua.Access.NOT_EQ(jet.Int32(int32(documents.DOCUMENT_ACCESS_BLOCKED))),
			),
			jet.AND(
				dua.Access.IS_NULL(),
				dja.Access.IS_NOT_NULL(),
				dja.Access.NOT_EQ(jet.Int32(int32(documents.DOCUMENT_ACCESS_BLOCKED))),
			),
		),
	)}

	if where != nil {
		wheres = append(wheres, where)
	}

	u := u.AS("creator")
	var q jet.SelectStatement
	if onlyColumns != nil {
		q = d.SELECT(
			onlyColumns,
		)
	} else {
		if additionalColumns == nil {
			q = d.SELECT(
				d.AllColumns,
				u.ID,
				u.Identifier,
				u.Job,
				u.JobGrade,
				u.Firstname,
				u.Lastname,
			)
		} else {
			additionalColumns = append(jet.ProjectionList{
				u.ID,
				u.Identifier,
				u.Job,
				u.JobGrade,
				u.Firstname,
				u.Lastname,
			}, additionalColumns)
			q = d.SELECT(
				d.AllColumns,
				additionalColumns...,
			)
		}
	}

	return q.
		FROM(
			d.LEFT_JOIN(dua,
				dua.DocumentID.EQ(d.ID).
					AND(dua.UserID.EQ(jet.Int32(userID)))).
				LEFT_JOIN(dja,
					dja.DocumentID.EQ(d.ID).
						AND(dja.Job.EQ(jet.String(job))).
						AND(dja.MinimumGrade.LT_EQ(jet.Int32(jobGrade))),
				).
				LEFT_JOIN(u, u.ID.EQ(d.CreatorID)),
		).WHERE(
		jet.AND(
			wheres...,
		),
	).
		ORDER_BY(d.CreatedAt.DESC()).
		LIMIT(database.DefaultPageLimit)

}

func (s *Server) FindDocuments(ctx context.Context, req *FindDocumentsRequest) (*FindDocumentsResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "FindDocuments") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to list documents!")
	}
	condition := d.ResponseID.IS_NULL()
	if req.Search != "" {
		condition = condition.AND(jet.BoolExp(jet.Raw("MATCH(title) AGAINST ($search IN NATURAL LANGUAGE MODE)", jet.RawArgs{"$search": req.Search})))
	}

	countStmt := s.getDocumentsQuery(condition, jet.ProjectionList{jet.COUNT(d.ID).AS("total_count")}, nil, userID, job, jobGrade)
	var count struct{ TotalCount int64 }
	if err := countStmt.QueryContext(ctx, query.DB, &count); err != nil {
		return nil, err
	}

	resp := &FindDocumentsResponse{
		Offset:     req.Offset,
		TotalCount: count.TotalCount,
		End:        0,
	}

	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := s.getDocumentsQuery(condition, nil, nil, userID, job, jobGrade)
	if err := stmt.QueryContext(ctx, query.DB, &resp.Documents); err != nil {
		return nil, err
	}

	resp.TotalCount = count.TotalCount
	if req.Offset >= resp.TotalCount {
		resp.Offset = 0
	} else {
		resp.Offset = req.Offset
	}
	resp.End = resp.Offset + int64(len(resp.Documents))

	return resp, nil
}

func (s *Server) GetDocument(ctx context.Context, req *GetDocumentRequest) (*GetDocumentResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "GetDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to get a document!")
	}

	stmt := s.getDocumentsQuery(jet.AND(
		d.ResponseID.IS_NULL(),
		d.ID.EQ(jet.Uint64(req.Id)),
	), nil, nil, userID, job, jobGrade).
		LIMIT(1)

	resp := &GetDocumentResponse{
		Document:  &documents.Document{},
		Responses: []*documents.Document{},
	}
	if err := stmt.QueryContext(ctx, query.DB, resp.Document); err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	return resp, nil
}

func (s *Server) CreateDocument(ctx context.Context, req *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	userID, job, _ := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "CreateDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to create documents!")
	}

	d := table.ArpanetDocuments
	stmt := d.INSERT(
		d.Title,
		d.Content,
		d.ContentType,
		d.Closed,
		d.State,
		d.Public,
		d.CreatorID,
		d.CreatorJob,
	).VALUES(
		req.Title,
		s.p.Sanitize(req.Content),
		documents.DOCUMENT_CONTENT_TYPE_HTML,
		req.Closed,
		req.State,
		req.Public,
		userID,
		job,
	)

	result, err := stmt.ExecContext(ctx, query.DB)
	if err != nil {
		return nil, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	if err := s.handleDocumentAccess(ctx, DOCUMENT_ACCESS_UPDATE_MODE_REPLACE, uint64(lastID), req.JobsAccess, req.UsersAccess); err != nil {
		return nil, err
	}

	return &CreateDocumentResponse{
		Id: uint64(lastID),
	}, nil
}

func (s *Server) checkIfUserCanEditDocument(ctx context.Context, userID int32, job string, jobGrade int32) (bool, error) {
	checkStmt := d.SELECT(
		d.ID,
	).
		FROM(
			d.LEFT_JOIN(dua,
				dua.DocumentID.EQ(d.ID).
					AND(dua.UserID.EQ(jet.Int32(userID)))).
				LEFT_JOIN(dja,
					dja.DocumentID.EQ(d.ID).
						AND(dja.Job.EQ(jet.String(job))).
						AND(dja.MinimumGrade.LT_EQ(jet.Int32(jobGrade))),
				),
		).WHERE(
		jet.OR(
			d.CreatorID.EQ(jet.Int32(userID)),
			jet.AND(
				dua.Access.IS_NOT_NULL(),
				dua.Access.IN(
					jet.Int32(int32(documents.DOCUMENT_ACCESS_EDIT)),
					jet.Int32(int32(documents.DOCUMENT_ACCESS_ADMIN)),
				),
			),
			jet.AND(
				dua.Access.IS_NULL(),
				dja.Access.IS_NOT_NULL(),
				dja.Access.IN(
					jet.Int32(int32(documents.DOCUMENT_ACCESS_EDIT)),
					jet.Int32(int32(documents.DOCUMENT_ACCESS_LEADER)),
					jet.Int32(int32(documents.DOCUMENT_ACCESS_ADMIN)),
				),
			),
		),
	).
		LIMIT(1)

	var dest struct {
		ID uint64 `alias:"document.id"`
	}
	if err := checkStmt.QueryContext(ctx, query.DB, &dest); err != nil {
		return false, err
	}

	return dest.ID > 0, nil
}

func (s *Server) UpdateDocument(ctx context.Context, req *UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "UpdateDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to edit documents!")
	}

	check, err := s.checkIfUserCanEditDocument(ctx, userID, job, jobGrade)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to edit this document!")
	}

	stmt := d.UPDATE(
		d.Title,
		d.Content,
		d.Closed,
		d.State,
		d.Public,
	).
		SET(
			req.Title,
			req.Content,
			req.Closed,
			req.State,
			req.Public,
		).
		WHERE(d.ID.EQ(jet.Uint64(req.Id)))

	if _, err := stmt.ExecContext(ctx, query.DB); err != nil {
		return nil, err
	}

	return &UpdateDocumentResponse{}, nil
}

func (s *Server) ListTemplates(ctx context.Context, req *ListTemplatesRequest) (*ListTemplatesResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "CreateDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to list/ get document templates!")
	}

	stmt := dt.SELECT(
		dt.ID,
		dt.Job,
		dt.JobGrade,
		dt.Title,
		dt.Description,
		dt.CreatorID,
	).
		FROM(dt).
		WHERE(
			jet.AND(
				dt.Job.EQ(jet.String(job)),
				dt.JobGrade.LT_EQ(jet.Int32(jobGrade)),
			),
		)

	resp := &ListTemplatesResponse{}
	if err := stmt.QueryContext(ctx, query.DB, &resp.Templates); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) GetTemplate(ctx context.Context, req *GetTemplateRequest) (*GetTemplateResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "CreateDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to list/ get document templates!")
	}

	stmt := dt.SELECT(
		dt.AllColumns,
	).
		FROM(dt).
		WHERE(
			jet.AND(
				dt.ID.EQ(jet.Uint64(req.Id)),
				dt.Job.EQ(jet.String(job)),
				dt.JobGrade.LT_EQ(jet.Int32(jobGrade)),
			),
		)

	resp := &GetTemplateResponse{}
	if err := stmt.QueryContext(ctx, query.DB, &resp.Template); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) GetDocumentAccess(ctx context.Context, req *GetDocumentAccessRequest) (*GetDocumentAccessResponse, error) {
	userID, job, jobGrade := auth.GetUserInfoFromContext(ctx)
	if !perms.P.CanID(userID, "documents", "GetDocument") {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to get document access!")
	}

	check, err := s.checkIfUserCanEditDocument(ctx, userID, job, jobGrade)
	if err != nil {
		return nil, err
	}
	if !check {
		return nil, status.Error(codes.PermissionDenied, "You don't have permission to get document access!")
	}

	stmt := d.SELECT(
		dua.AllColumns,
		dja.AllColumns,
	).
		FROM(d).
		WHERE(
			jet.AND(
				d.ID.EQ(jet.Uint64(req.Id)),
				jet.OR(
					jet.AND(
						dua.Access.IS_NOT_NULL(),
						dua.Access.NOT_EQ(jet.Int32(int32(documents.DOCUMENT_ACCESS_BLOCKED))),
					),
					jet.AND(
						dua.Access.IS_NULL(),
						dja.Access.IS_NOT_NULL(),
						dja.Access.NOT_EQ(jet.Int32(int32(documents.DOCUMENT_ACCESS_BLOCKED))),
					),
				),
			),
		)

	resp := &GetDocumentAccessResponse{}
	if err := stmt.QueryContext(ctx, query.DB, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) SetDocumentAccess(ctx context.Context, req *SetDocumentAccessRequest) (*SetDocumentAccessResponse, error) {
	resp := &SetDocumentAccessResponse{}

	if err := s.handleDocumentAccess(ctx, req.Mode, req.DocumentID, req.Jobs, req.Users); err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *Server) handleDocumentAccess(ctx context.Context, mode DOCUMENT_ACCESS_UPDATE_MODE, documentID uint64, ja []*documents.DocumentJobAccess, ua []*documents.DocumentUserAccess) error {
	// Select existing job and user accesses
	var dest struct {
		jobs  []*documents.DocumentJobAccess
		users []*documents.DocumentUserAccess
	}
	selectStmt := jet.SELECT(
		dja.AllColumns,
		dua.AllColumns,
	).
		FROM(
			dja,
			dua,
		).
		WHERE(
			dja.DocumentID.EQ(jet.Uint64(documentID)),
		)

	fmt.Println(selectStmt.DebugSql())

	if err := selectStmt.QueryContext(ctx, query.DB, &dest); err != nil && errors.Is(err, sql.ErrNoRows) {
		return err
	}

	// TODO add/update/remove for document access based on the current access in the database

	// Create accesses
	if len(ja) > 0 {
		for k := 0; k < len(ja); k++ {
			ja[k].DocumentID = documentID
		}

		// Create document job access
		stmt := dja.INSERT(
			dja.DocumentID,
			dja.Job,
			dja.MinimumGrade,
			dja.Access,
		).
			MODELS(ja)
		fmt.Println(stmt.DebugSql())
		if _, err := stmt.ExecContext(ctx, query.DB); err != nil {
			return err
		}
	}

	if len(ua) > 0 {
		for k := 0; k < len(ua); k++ {
			ua[k].DocumentID = documentID
		}
		// Create document user access
		stmt := dua.INSERT(
			dua.DocumentID,
			dua.UserID,
			dua.Access,
		).
			MODELS(ua)
		if _, err := stmt.ExecContext(ctx, query.DB); err != nil {
			return err
		}
	}

	return nil
}
