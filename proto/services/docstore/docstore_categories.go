package docstore

import (
	"context"
	"errors"

	"github.com/galexrt/fivenet/pkg/auth"
	"github.com/galexrt/fivenet/proto/resources/documents"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"github.com/go-jet/jet/v2/qrm"
)

var (
	dCategory = table.FivenetDocumentsCategories.AS("category")
)

func (s *Server) ListDocumentCategories(ctx context.Context, req *ListDocumentCategoriesRequest) (*ListDocumentCategoriesResponse, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	dCategory := table.FivenetDocumentsCategories.AS("documentcategory")
	stmt := dCategory.
		SELECT(
			dCategory.AllColumns,
		).
		FROM(
			dCategory,
		).
		WHERE(
			dCategory.Job.EQ(jet.String(job)),
		)

	resp := &ListDocumentCategoriesResponse{}
	if err := stmt.QueryContext(ctx, s.db, &resp.Category); err != nil {
		if !errors.Is(qrm.ErrNoRows, err) {
			return nil, err
		}
	}

	return resp, nil
}

func (s *Server) getDocumentCategory(ctx context.Context, id uint64) (*documents.DocumentCategory, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	dCategory := table.FivenetDocumentsCategories.AS("documentcategory")
	stmt := dCategory.
		SELECT(
			dCategory.AllColumns,
		).
		FROM(
			dCategory,
		).
		WHERE(
			jet.AND(
				dCategory.Job.EQ(jet.String(job)),
				dCategory.ID.EQ(jet.Uint64(id)),
			),
		)

	var dest documents.DocumentCategory
	if err := stmt.QueryContext(ctx, s.db, &dest); err != nil {
		return nil, err
	}

	return &dest, nil
}

func (s *Server) CreateDocumentCategory(ctx context.Context, req *CreateDocumentCategoryRequest) (*CreateDocumentCategoryResponse, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	dCategory := table.FivenetDocumentsCategories
	stmt := dCategory.
		INSERT(
			dCategory.Name,
			dCategory.Description,
			dCategory.Job,
		).
		VALUES(
			req.Category.Name,
			req.Category.Description,
			job,
		)

	res, err := stmt.ExecContext(ctx, s.db)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &CreateDocumentCategoryResponse{
		Id: lastId,
	}, nil
}

func (s *Server) UpdateDocumentCategory(ctx context.Context, req *UpdateDocumentCategoryRequest) (*UpdateDocumentCategoryResponse, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	dCategory := table.FivenetDocumentsCategories
	stmt := dCategory.
		UPDATE(
			dCategory.Name,
			dCategory.Description,
			dCategory.Job,
		).
		SET(
			req.Category.Name,
			req.Category.Description,
			job,
		).
		WHERE(
			jet.AND(
				dCategory.ID.EQ(jet.Uint64(req.Category.Id)),
				dCategory.Job.EQ(jet.String(job)),
			),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	return &UpdateDocumentCategoryResponse{}, nil
}

func (s *Server) DeleteDocumentCategory(ctx context.Context, req *DeleteDocumentCategoryRequest) (*DeleteDocumentCategoryResponse, error) {
	_, job, _ := auth.GetUserInfoFromContext(ctx)

	ids := make([]jet.Expression, len(req.Ids))
	for i := 0; i < len(req.Ids); i++ {
		ids[i] = jet.Uint64(req.Ids[i])
	}

	dCategory := table.FivenetDocumentsCategories
	stmt := dCategory.
		DELETE().
		WHERE(
			jet.AND(
				dCategory.Job.EQ(jet.String(job)),
				dCategory.ID.IN(ids...),
			),
		)

	if _, err := stmt.ExecContext(ctx, s.db); err != nil {
		return nil, err
	}

	return &DeleteDocumentCategoryResponse{}, nil
}
