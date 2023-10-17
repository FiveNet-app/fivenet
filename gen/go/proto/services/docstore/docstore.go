package docstore

import (
	context "context"
	"database/sql"
	"errors"

	database "github.com/galexrt/fivenet/gen/go/proto/resources/common/database"
	"github.com/galexrt/fivenet/gen/go/proto/resources/documents"
	"github.com/galexrt/fivenet/gen/go/proto/resources/rector"
	permsdocstore "github.com/galexrt/fivenet/gen/go/proto/services/docstore/perms"
	"github.com/galexrt/fivenet/pkg/grpc/auth"
	"github.com/galexrt/fivenet/pkg/grpc/auth/userinfo"
	"github.com/galexrt/fivenet/pkg/htmlsanitizer"
	"github.com/galexrt/fivenet/pkg/mstlystcdata"
	"github.com/galexrt/fivenet/pkg/notifi"
	"github.com/galexrt/fivenet/pkg/perms"
	"github.com/galexrt/fivenet/pkg/server/audit"
	"github.com/galexrt/fivenet/pkg/utils"
	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	DocsDefaultPageLimit  = 15
	DocShortContentLength = 128
)

var (
	tUsers       = table.Users
	tCreator     = tUsers.AS("creator")
	tDocs        = table.FivenetDocuments.AS("document")
	tDUserAccess = table.FivenetDocumentsUserAccess.AS("user_access")
	tDJobAccess  = table.FivenetDocumentsJobAccess.AS("job_access")
)

var (
	ErrFailedQuery       = status.Error(codes.Internal, "errors.DocStoreService.ErrFailedQuery")
	ErrNotFoundOrNoPerms = status.Error(codes.NotFound, "errors.DocStoreService.ErrNotFoundOrNoPerms")
	ErrTemplateNoPerms   = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrTemplateNoPerms")
	ErrPermissionDenied  = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrPermissionDenied")
	ErrClosedDoc         = status.Error(codes.Canceled, "errors.DocStoreService.ErrClosedDoc")
	ErrDocViewDenied     = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrDocViewDenied")
	ErrDocUpdateDenied   = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrDocUpdateDenied")
	ErrDocDeleteDenied   = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrDocDeleteDenied")
	ErrDocToggleDenied   = status.Error(codes.PermissionDenied, "errors.DocStoreService.ErrDocToggleDenied")
)

type Server struct {
	DocStoreServiceServer

	db       *sql.DB
	ps       perms.Permissions
	cache    *mstlystcdata.Cache
	enricher *mstlystcdata.Enricher
	auditer  audit.IAuditer
	ui       userinfo.UserInfoRetriever
	notif    notifi.INotifi
}

func NewServer(db *sql.DB, ps perms.Permissions, cache *mstlystcdata.Cache, enricher *mstlystcdata.Enricher, auditer audit.IAuditer, ui userinfo.UserInfoRetriever, notif notifi.INotifi) *Server {
	return &Server{
		db:       db,
		ps:       ps,
		cache:    cache,
		enricher: enricher,
		auditer:  auditer,
		ui:       ui,
		notif:    notif,
	}
}

func (s *Server) ListDocuments(ctx context.Context, req *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	tDocs := tDocs.AS("documentshort")
	condition := jet.Bool(true)
	if req.Search != nil && *req.Search != "" {
		condition = jet.BoolExp(
			jet.Raw("MATCH(`title`) AGAINST ($search IN BOOLEAN MODE)",
				jet.RawArgs{"$search": *req.Search}),
		)
	}
	if len(req.CategoryIds) > 0 {
		ids := make([]jet.Expression, len(req.CategoryIds))
		for i := 0; i < len(req.CategoryIds); i++ {
			ids[i] = jet.Uint64(req.CategoryIds[i])
		}

		condition = condition.AND(
			tDocs.CategoryID.IN(ids...),
		)
	}
	if len(req.CreatorIds) > 0 {
		ids := make([]jet.Expression, len(req.CreatorIds))
		for i := 0; i < len(req.CreatorIds); i++ {
			ids[i] = jet.Int32(req.CreatorIds[i])
		}

		condition = condition.AND(
			tDocs.CreatorID.IN(ids...),
		)
	}
	if req.From != nil {
		condition = condition.AND(tDocs.CreatedAt.GT_EQ(
			jet.TimestampT(req.From.AsTime()),
		))
	}
	if req.To != nil {
		condition = condition.AND(tDocs.CreatedAt.LT_EQ(
			jet.TimestampT(req.To.AsTime()),
		))
	}
	if req.Closed != nil {
		condition = condition.AND(tDocs.Closed.EQ(
			jet.Bool(*req.Closed),
		))
	}

	countStmt := s.listDocumentsQuery(
		condition, jet.ProjectionList{jet.COUNT(jet.DISTINCT(tDocs.ID)).AS("datacount.totalcount")},
		1, userInfo)

	var count database.DataCount
	if err := countStmt.QueryContext(ctx, s.db, &count); err != nil {
		return nil, ErrFailedQuery
	}

	pag, limit := req.Pagination.GetResponseWithPageSize(DocsDefaultPageLimit)
	resp := &ListDocumentsResponse{
		Pagination: pag,
	}
	if count.TotalCount <= 0 {
		return resp, nil
	}

	stmt := s.listDocumentsQuery(condition, nil,
		DocShortContentLength, userInfo).
		OFFSET(req.Pagination.Offset).
		GROUP_BY(tDocs.ID).
		LIMIT(limit)

	if err := stmt.QueryContext(ctx, s.db, &resp.Documents); err != nil {
		return nil, ErrFailedQuery
	}

	jobInfoFn := s.enricher.EnrichJobInfoFunc(userInfo)
	for i := 0; i < len(resp.Documents); i++ {
		if resp.Documents[i].Creator != nil {
			jobInfoFn(resp.Documents[i].Creator)
		}
	}

	resp.Pagination.Update(count.TotalCount, len(resp.Documents))

	return resp, nil
}

func (s *Server) GetDocument(ctx context.Context, req *GetDocumentRequest) (*GetDocumentResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "GetDocument",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userInfo, documents.AccessLevel_ACCESS_LEVEL_VIEW)
	if err != nil {
		return nil, ErrNotFoundOrNoPerms
	}
	if !check && !userInfo.SuperUser {
		return nil, ErrDocViewDenied
	}

	resp := &GetDocumentResponse{}
	resp.Document, err = s.getDocument(ctx,
		tDocs.ID.EQ(jet.Uint64(req.DocumentId)), userInfo)
	if err != nil {
		return nil, ErrFailedQuery
	}

	if resp.Document == nil || resp.Document.Id <= 0 {
		return nil, ErrNotFoundOrNoPerms
	}

	docAccess, err := s.GetDocumentAccess(ctx, &GetDocumentAccessRequest{
		DocumentId: resp.Document.Id,
	})
	if err != nil {
		if st, ok := status.FromError(err); !ok {
			return nil, ErrFailedQuery
		} else {
			// Ignore permission denied as we are simply getting the document
			if st.Code() != codes.PermissionDenied {
				return nil, ErrFailedQuery
			}
		}
	}
	if docAccess != nil {
		resp.Access = docAccess.Access
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_VIEWED)

	return resp, nil
}

func (s *Server) getDocument(ctx context.Context, condition jet.BoolExpression, userInfo *userinfo.UserInfo) (*documents.Document, error) {
	var doc documents.Document

	stmt := s.getDocumentsQuery(condition, nil, userInfo).
		LIMIT(1)

	if err := stmt.QueryContext(ctx, s.db, &doc); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, ErrFailedQuery
		}
	}

	if doc.Creator != nil {
		s.enricher.EnrichJobInfoSafe(userInfo, doc.Creator)
	}

	return &doc, nil
}

func (s *Server) CreateDocument(ctx context.Context, req *CreateDocumentRequest) (*CreateDocumentResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "CreateDocument",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ErrFailedQuery
	}
	// Defer a rollback in case anything fails
	defer tx.Rollback()

	tDocs := table.FivenetDocuments
	stmt := tDocs.
		INSERT(
			tDocs.CategoryID,
			tDocs.Title,
			tDocs.Summary,
			tDocs.Content,
			tDocs.ContentType,
			tDocs.Data,
			tDocs.CreatorID,
			tDocs.CreatorJob,
			tDocs.State,
			tDocs.Closed,
			tDocs.Public,
		).
		VALUES(
			req.CategoryId,
			req.Title,
			utils.StringFirstN(htmlsanitizer.StripTags(req.Content), DocShortContentLength),
			req.Content,
			documents.DocContentType_DOC_CONTENT_TYPE_HTML,
			req.Data,
			userInfo.UserId,
			userInfo.Job,
			req.State,
			req.Closed,
			req.Public,
		)

	result, err := stmt.ExecContext(ctx, tx)
	if err != nil {
		return nil, ErrFailedQuery
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return nil, ErrFailedQuery
	}

	if err := s.handleDocumentAccessChanges(ctx, tx, AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_UPDATE, uint64(lastId), req.Access); err != nil {
		return nil, ErrFailedQuery
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, ErrFailedQuery
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_CREATED)

	return &CreateDocumentResponse{
		DocumentId: uint64(lastId),
	}, nil
}

func (s *Server) UpdateDocument(ctx context.Context, req *UpdateDocumentRequest) (*UpdateDocumentResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "UpdateDocument",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userInfo, documents.AccessLevel_ACCESS_LEVEL_EDIT)
	if err != nil {
		return nil, ErrNotFoundOrNoPerms
	}
	if !check && !userInfo.SuperUser {
		return nil, ErrPermissionDenied
	}

	doc, err := s.getDocument(ctx,
		tDocs.ID.EQ(jet.Uint64(req.DocumentId)),
		userInfo)
	if err != nil {
		return nil, ErrFailedQuery
	}

	// Either the document is closed and the update request isn't re-opening the document
	if doc.Closed && req.Closed && !userInfo.SuperUser {
		return nil, ErrClosedDoc
	}

	// Field Permission Check
	fieldsAttr, err := s.ps.Attr(userInfo, permsdocstore.DocStoreServicePerm, permsdocstore.DocStoreServiceUpdateDocumentPerm, permsdocstore.DocStoreServiceUpdateDocumentAccessPermField)
	if err != nil {
		return nil, ErrFailedQuery
	}
	var fields perms.StringList
	if fieldsAttr != nil {
		fields = fieldsAttr.([]string)
	}
	if !s.checkIfHasAccess(fields, userInfo, doc.Creator) {
		return nil, ErrDocUpdateDenied
	}

	// Begin transaction
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, ErrFailedQuery
	}

	tDocs := table.FivenetDocuments
	stmt := tDocs.
		UPDATE(
			tDocs.CategoryID,
			tDocs.Title,
			tDocs.Summary,
			tDocs.Content,
			tDocs.Data,
			tDocs.State,
			tDocs.Closed,
			tDocs.Public,
		).
		SET(
			req.CategoryId,
			req.Title,
			utils.StringFirstN(htmlsanitizer.StripTags(req.Content), DocShortContentLength),
			req.Content,
			tDocs.Data,
			req.State,
			req.Closed,
			req.Public,
		).
		WHERE(
			tDocs.ID.EQ(jet.Uint64(req.DocumentId)),
		)

	if _, err := stmt.ExecContext(ctx, tx); err != nil {
		return nil, ErrFailedQuery
	}

	if err := s.handleDocumentAccessChanges(ctx, tx, AccessLevelUpdateMode_ACCESS_LEVEL_UPDATE_MODE_UPDATE, req.DocumentId, req.Access); err != nil {
		return nil, ErrFailedQuery
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		return nil, ErrFailedQuery
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &UpdateDocumentResponse{
		DocumentId: req.DocumentId,
	}, nil
}

func (s *Server) DeleteDocument(ctx context.Context, req *DeleteDocumentRequest) (*DeleteDocumentResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "DeleteDocument",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userInfo, documents.AccessLevel_ACCESS_LEVEL_EDIT)
	if err != nil {
		return nil, ErrNotFoundOrNoPerms
	}
	if !check && !userInfo.SuperUser {
		if !userInfo.SuperUser {
			return nil, ErrDocDeleteDenied
		}
	}

	doc, err := s.getDocument(ctx, tDocs.ID.EQ(jet.Uint64(req.DocumentId)), userInfo)
	if err != nil {
		return nil, ErrFailedQuery
	}

	// Field Permission Check
	fieldsAttr, err := s.ps.Attr(userInfo, permsdocstore.DocStoreServicePerm, permsdocstore.DocStoreServiceDeleteDocumentPerm, permsdocstore.DocStoreServiceDeleteDocumentAccessPermField)
	if err != nil {
		return nil, ErrFailedQuery
	}
	var fields perms.StringList
	if fieldsAttr != nil {
		fields = fieldsAttr.([]string)
	}
	if !s.checkIfHasAccess(fields, userInfo, doc.Creator) {
		return nil, ErrDocDeleteDenied
	}

	stmt := tDocs.
		UPDATE(
			tDocs.DeletedAt,
		).
		SET(
			tDocs.DeletedAt.SET(jet.CURRENT_TIMESTAMP()),
		).
		WHERE(
			tDocs.ID.EQ(jet.Uint64(req.DocumentId)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, ErrFailedQuery
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_DELETED)

	return &DeleteDocumentResponse{}, nil
}

func (s *Server) ToggleDocument(ctx context.Context, req *ToggleDocumentRequest) (*ToggleDocumentResponse, error) {
	userInfo := auth.MustGetUserInfoFromContext(ctx)

	auditEntry := &model.FivenetAuditLog{
		Service: DocStoreService_ServiceDesc.ServiceName,
		Method:  "ToggleDocument",
		UserID:  userInfo.UserId,
		UserJob: userInfo.Job,
		State:   int16(rector.EventType_EVENT_TYPE_ERRORED),
	}
	defer s.auditer.Log(auditEntry, req)

	check, err := s.checkIfUserHasAccessToDoc(ctx, req.DocumentId, userInfo, documents.AccessLevel_ACCESS_LEVEL_STATUS)
	if err != nil {
		return nil, ErrNotFoundOrNoPerms
	}
	if !check && !userInfo.SuperUser {
		if !userInfo.SuperUser {
			return nil, ErrDocToggleDenied
		}
	}

	doc, err := s.getDocument(ctx, tDocs.ID.EQ(jet.Uint64(req.DocumentId)), userInfo)
	if err != nil {
		return nil, ErrFailedQuery
	}

	// Field Permission Check
	fieldsAttr, err := s.ps.Attr(userInfo, permsdocstore.DocStoreServicePerm, permsdocstore.DocStoreServiceToggleDocumentPerm, permsdocstore.DocStoreServiceToggleDocumentAccessPermField)
	if err != nil {
		return nil, ErrFailedQuery
	}
	var fields perms.StringList
	if fieldsAttr != nil {
		fields = fieldsAttr.([]string)
	}
	if !s.checkIfHasAccess(fields, userInfo, doc.Creator) {
		return nil, ErrDocToggleDenied
	}

	stmt := tDocs.
		UPDATE(
			tDocs.Closed,
		).
		SET(
			tDocs.Closed.SET(jet.Bool(req.Closed)),
		).
		WHERE(
			tDocs.ID.EQ(jet.Uint64(req.DocumentId)),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, ErrFailedQuery
	}

	auditEntry.State = int16(rector.EventType_EVENT_TYPE_UPDATED)

	return &ToggleDocumentResponse{}, nil
}