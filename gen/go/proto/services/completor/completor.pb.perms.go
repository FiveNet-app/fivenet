// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/completor/completor.proto

package completor

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	permkeys "github.com/galexrt/fivenet/gen/go/proto/services/completor/perms"
	"github.com/galexrt/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: CompletorService
	"CompletorService/ListLawBooks": "Any",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: CompletorService
		{
			Category: permkeys.CompletorServicePerm,
			Name:     permkeys.CompletorServiceCompleteCitizensPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CompletorServicePerm,
			Name:     permkeys.CompletorServiceCompleteDocumentCategoriesPerm,
			Attrs: []perms.Attr{
				{
					Key:  permkeys.CompletorServiceCompleteDocumentCategoriesJobsPermField,
					Type: permissions.JobListAttributeType,
				},
			},
		},
		{
			Category: permkeys.CompletorServicePerm,
			Name:     permkeys.CompletorServiceCompleteJobsPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
