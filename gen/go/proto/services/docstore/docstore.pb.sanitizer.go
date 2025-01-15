// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/docstore/docstore.proto

package docstore

import (
	"github.com/fivenet-app/fivenet/pkg/html/htmlsanitizer"
)

func (m *AddDocumentReferenceRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Reference
	if m.Reference != nil {
		if v, ok := interface{}(m.GetReference()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *AddDocumentReferenceResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *AddDocumentRelationRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Relation
	if m.Relation != nil {
		if v, ok := interface{}(m.GetRelation()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *AddDocumentRelationResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ChangeDocumentOwnerRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ChangeDocumentOwnerResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *CreateCategoryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Category
	if m.Category != nil {
		if v, ok := interface{}(m.GetCategory()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CreateCategoryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *CreateDocumentReqRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Data
	if m.Data != nil {
		if v, ok := interface{}(m.GetData()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Reason

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *CreateDocumentReqResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Request
	if m.Request != nil {
		if v, ok := interface{}(m.GetRequest()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CreateDocumentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Access
	if m.Access != nil {
		if v, ok := interface{}(m.GetAccess()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Content
	if m.Content != nil {
		if v, ok := interface{}(m.GetContent()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: State
	m.State = htmlsanitizer.Sanitize(m.State)

	// Field: Title
	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}

func (m *CreateDocumentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *CreateTemplateRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Template
	if m.Template != nil {
		if v, ok := interface{}(m.GetTemplate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *CreateTemplateResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCategoryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCategoryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCommentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteCommentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteDocumentReqRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteDocumentReqResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteDocumentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteDocumentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteTemplateRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteTemplateResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *EditCommentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Comment
	if m.Comment != nil {
		if v, ok := interface{}(m.GetComment()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *EditCommentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Comment
	if m.Comment != nil {
		if v, ok := interface{}(m.GetComment()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetCommentsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetCommentsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Comments
	for idx, item := range m.Comments {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetDocumentAccessRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetDocumentAccessResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Access
	if m.Access != nil {
		if v, ok := interface{}(m.GetAccess()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetDocumentReferencesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetDocumentReferencesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: References
	for idx, item := range m.References {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GetDocumentRelationsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetDocumentRelationsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Relations
	for idx, item := range m.Relations {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *GetDocumentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *GetDocumentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Access
	if m.Access != nil {
		if v, ok := interface{}(m.GetAccess()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Document
	if m.Document != nil {
		if v, ok := interface{}(m.GetDocument()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetTemplateRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Data
	if m.Data != nil {
		if v, ok := interface{}(m.GetData()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *GetTemplateResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Template
	if m.Template != nil {
		if v, ok := interface{}(m.GetTemplate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListCategoriesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ListCategoriesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Categories
	for idx, item := range m.Categories {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListDocumentActivityRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: ActivityTypes
	for idx, item := range m.ActivityTypes {
		_, _ = idx, item

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentActivityResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Activity
	for idx, item := range m.Activity {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentPinsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentPinsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Documents
	for idx, item := range m.Documents {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentReqsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentReqsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Requests
	for idx, item := range m.Requests {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListDocumentsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: From
	if m.From != nil {
		if v, ok := interface{}(m.GetFrom()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Sort
	if m.Sort != nil {
		if v, ok := interface{}(m.GetSort()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: To
	if m.To != nil {
		if v, ok := interface{}(m.GetTo()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListDocumentsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Documents
	for idx, item := range m.Documents {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *ListTemplatesRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ListTemplatesResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Templates
	for idx, item := range m.Templates {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ListUserDocumentsRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Relations
	for idx, item := range m.Relations {
		_, _ = idx, item

	}

	return nil
}

func (m *ListUserDocumentsResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Pagination
	if m.Pagination != nil {
		if v, ok := interface{}(m.GetPagination()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Relations
	for idx, item := range m.Relations {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *PostCommentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Comment
	if m.Comment != nil {
		if v, ok := interface{}(m.GetComment()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *PostCommentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Comment
	if m.Comment != nil {
		if v, ok := interface{}(m.GetComment()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *RemoveDocumentReferenceRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *RemoveDocumentReferenceResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *RemoveDocumentRelationRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *RemoveDocumentRelationResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *SetDocumentAccessRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Access
	if m.Access != nil {
		if v, ok := interface{}(m.GetAccess()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *SetDocumentAccessResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *SetDocumentReminderRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Message

	if m.Message != nil {
		*m.Message = htmlsanitizer.StripTags(*m.Message)
	}

	// Field: ReminderTime
	if m.ReminderTime != nil {
		if v, ok := interface{}(m.GetReminderTime()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *SetDocumentReminderResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ToggleDocumentPinRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ToggleDocumentPinResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ToggleDocumentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *ToggleDocumentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *UpdateCategoryRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Category
	if m.Category != nil {
		if v, ok := interface{}(m.GetCategory()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateCategoryResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *UpdateDocumentReqRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Data
	if m.Data != nil {
		if v, ok := interface{}(m.GetData()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Reason

	if m.Reason != nil {
		*m.Reason = htmlsanitizer.Sanitize(*m.Reason)
	}

	return nil
}

func (m *UpdateDocumentReqResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Request
	if m.Request != nil {
		if v, ok := interface{}(m.GetRequest()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateDocumentRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Access
	if m.Access != nil {
		if v, ok := interface{}(m.GetAccess()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: Content
	if m.Content != nil {
		if v, ok := interface{}(m.GetContent()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	// Field: State
	m.State = htmlsanitizer.Sanitize(m.State)

	// Field: Title
	m.Title = htmlsanitizer.StripTags(m.Title)

	return nil
}

func (m *UpdateDocumentResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	return nil
}

func (m *UpdateTemplateRequest) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Template
	if m.Template != nil {
		if v, ok := interface{}(m.GetTemplate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (m *UpdateTemplateResponse) Sanitize() error {
	if m == nil {
		return nil
	}

	// Field: Template
	if m.Template != nil {
		if v, ok := interface{}(m.GetTemplate()).(interface{ Sanitize() error }); ok {
			if err := v.Sanitize(); err != nil {
				return err
			}
		}
	}

	return nil
}
