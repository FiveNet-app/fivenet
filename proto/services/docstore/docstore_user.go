package docstore

import (
	context "context"
	"errors"

	"github.com/galexrt/arpanet/pkg/auth"
	"github.com/galexrt/arpanet/proto/resources/documents"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

func (s *Server) GetUserDocuments(ctx context.Context, req *GetUserDocumentsRequest) (*GetUserDocumentsResponse, error) {
	userId, job, jobGrade := auth.GetUserInfoFromContext(ctx)

	resp := &GetUserDocumentsResponse{}
	// An user can never see their own activity on their own "profile"
	if userId == req.UserId {
		return resp, nil
	}

	condition := jet.OR(
		docRel.SourceUserID.EQ(jet.Int32(req.UserId)),
		docRel.TargetUserID.EQ(jet.Int32(req.UserId)),
	)
	// TODO use query to get documents which the user has access to before selecting
	var docIds []uint64
	idStmt := docRel.
		SELECT(
			docs.ID,
		).
		FROM(
			docRel.
				LEFT_JOIN(docs,
					docRel.DocumentID.EQ(docs.ID),
				),
		).
		WHERE(
			condition,
		)

	if err := idStmt.QueryContext(ctx, s.db, &docIds); err != nil {
		return nil, err
	}

	ids, err := s.checkIfUserHasAccessToDocIDs(ctx, userId, job, jobGrade, true, documents.DOC_ACCESS_VIEW, docIds...)
	if err != nil {
		return nil, err
	}

	dIds := make([]jet.Expression, len(ids))
	for i := 0; i < len(ids); i++ {
		dIds[i] = jet.Uint64(ids[i])
	}

	dCreator := user.AS("creator")
	uSource := user.AS("source_user")
	uTarget := user.AS("target_user")
	stmt := docRel.
		SELECT(
			docRel.ID,
			docRel.CreatedAt,
			docRel.DeletedAt,
			docRel.DocumentID,
			docRel.SourceUserID,
			docs.ID,
			docs.CreatedAt,
			docs.UpdatedAt,
			docs.CategoryID,
			docs.CreatorID,
			docs.State,
			docs.Closed,
			docs.Title,
			dCreator.ID,
			dCreator.Identifier,
			dCreator.Job,
			dCreator.JobGrade,
			dCreator.Firstname,
			dCreator.Lastname,
			docRel.SourceUserID,
			uSource.ID,
			uSource.Identifier,
			uSource.Job,
			uSource.JobGrade,
			uSource.Firstname,
			uSource.Lastname,
			docRel.Relation,
			docRel.TargetUserID,
			uTarget.ID,
			uTarget.Identifier,
			uTarget.Job,
			uTarget.JobGrade,
			uTarget.Firstname,
			uTarget.Lastname,
		).
		FROM(
			docRel.
				LEFT_JOIN(docs,
					docRel.DocumentID.EQ(docs.ID),
				).
				LEFT_JOIN(dCreator,
					docs.CreatorID.EQ(dCreator.ID),
				).
				LEFT_JOIN(uSource,
					uSource.ID.EQ(docRel.SourceUserID),
				).
				LEFT_JOIN(uTarget,
					uTarget.ID.EQ(docRel.TargetUserID),
				),
		).
		WHERE(
			docRel.DocumentID.IN(dIds...),
		)

	if err := stmt.QueryContext(ctx, s.db, &resp.Relations); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
	}

	for i := 0; i < len(resp.Relations); i++ {
		s.c.EnrichJobInfo(resp.Relations[i].SourceUser)
		s.c.EnrichJobInfo(resp.Relations[i].TargetUser)
	}

	return resp, nil
}
