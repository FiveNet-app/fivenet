// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/completor/completor.proto

package completor

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	CompletorServicePerm perms.Category = "CompletorService"

	CompletorServiceCompleteCitizensPerm                    perms.Name = "CompleteCitizens"
	CompletorServiceCompleteDocumentCategoriesPerm          perms.Name = "CompleteDocumentCategories"
	CompletorServiceCompleteDocumentCategoriesJobsPermField perms.Key  = "Jobs"
	CompletorServiceCompleteJobsPerm                        perms.Name = "CompleteJobs"
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
			Category: CompletorServicePerm,
			Name:     CompletorServiceCompleteCitizensPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CompletorServicePerm,
			Name:     CompletorServiceCompleteDocumentCategoriesPerm,
			Attrs: []perms.Attr{
				{
					Key:  CompletorServiceCompleteDocumentCategoriesJobsPermField,
					Type: permissions.JobListAttributeType,
				},
			},
		},
		{
			Category: CompletorServicePerm,
			Name:     CompletorServiceCompleteJobsPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
