// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/docstore/docstore.proto

package docstore

import (
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/permissions"
	permkeys "github.com/fivenet-app/fivenet/gen/go/proto/services/docstore/perms"
	"github.com/fivenet-app/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: DocStoreService
	"DocStoreService/EditComment":             "DocStoreService/PostComment",
	"DocStoreService/GetComments":             "DocStoreService/GetDocument",
	"DocStoreService/GetDocumentAccess":       "DocStoreService/GetDocument",
	"DocStoreService/GetDocumentReferences":   "DocStoreService/GetDocument",
	"DocStoreService/GetDocumentRelations":    "DocStoreService/GetDocument",
	"DocStoreService/GetTemplate":             "DocStoreService/ListTemplates",
	"DocStoreService/RemoveDocumentReference": "DocStoreService/AddDocumentReference",
	"DocStoreService/RemoveDocumentRelation":  "DocStoreService/AddDocumentRelation",
	"DocStoreService/SetDocumentAccess":       "DocStoreService/CreateDocument",
	"DocStoreService/UpdateCategory":          "DocStoreService/CreateCategory",
	"DocStoreService/UpdateDocumentReq":       "DocStoreService/CreateDocumentReq",
	"DocStoreService/UpdateTemplate":          "DocStoreService/CreateTemplate",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{

		// Service: DocStoreService
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceAddDocumentReferencePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceAddDocumentRelationPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceChangeDocumentOwnerPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceChangeDocumentOwnerAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceCreateCategoryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceCreateDocumentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceCreateDocumentReqPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceCreateDocumentReqTypesPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Access", "Closure", "Update", "Deletion", "OwnerChange"},
				},
			},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceCreateTemplatePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceDeleteCategoryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceDeleteCommentPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceDeleteCommentAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceDeleteDocumentPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceDeleteDocumentAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceDeleteDocumentReqPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceDeleteTemplatePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceGetDocumentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListCategoriesPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListDocumentActivityPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListDocumentReqsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListDocumentsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListTemplatesPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceListUserDocumentsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServicePostCommentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceToggleDocumentPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceToggleDocumentAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
		{
			Category: permkeys.DocStoreServicePerm,
			Name:     permkeys.DocStoreServiceUpdateDocumentPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.DocStoreServiceUpdateDocumentAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
	})
}
