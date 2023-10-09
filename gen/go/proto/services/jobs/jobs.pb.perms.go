// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/jobs.proto

package jobs

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	permkeys "github.com/galexrt/fivenet/gen/go/proto/services/jobs/perms"
	"github.com/galexrt/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: JobsService
	"JobsService/RequestsListComments": "JobsService/RequestsListEntries",
	"JobsService/RequestsListTypes":    "JobsService/RequestsListEntries",
	"JobsService/RequestsPostComment":  "JobsService/RequestsCreateEntry",
	"JobsService/TimeclockStats":       "JobsService/TimeclockListEntries",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: JobsService
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceColleaguesListPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceConductCreateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceConductDeleteEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceConductListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           permkeys.JobsServiceConductListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceConductUpdateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsCreateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsCreateOrUpdateTypePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsDeleteCommentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsDeleteEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsDeleteTypePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           permkeys.JobsServiceRequestsListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceRequestsUpdateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.JobsServicePerm,
			Name:     permkeys.JobsServiceTimeclockListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           permkeys.JobsServiceTimeclockListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
	})
}
