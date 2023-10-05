// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/jobs.proto

package jobs

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	JobsServicePerm perms.Category = "JobsService"

	JobsServiceColleaguesListPerm                  perms.Name = "ColleaguesList"
	JobsServiceConductCreateEntryPerm              perms.Name = "ConductCreateEntry"
	JobsServiceConductDeleteEntryPerm              perms.Name = "ConductDeleteEntry"
	JobsServiceConductListEntriesPerm              perms.Name = "ConductListEntries"
	JobsServiceConductListEntriesAccessPermField   perms.Key  = "Access"
	JobsServiceConductUpdateEntryPerm              perms.Name = "ConductUpdateEntry"
	JobsServiceRequestsCreateEntryPerm             perms.Name = "RequestsCreateEntry"
	JobsServiceRequestsCreateOrUpdateTypePerm      perms.Name = "RequestsCreateOrUpdateType"
	JobsServiceRequestsDeleteCommentPerm           perms.Name = "RequestsDeleteComment"
	JobsServiceRequestsDeleteEntryPerm             perms.Name = "RequestsDeleteEntry"
	JobsServiceRequestsDeleteTypePerm              perms.Name = "RequestsDeleteType"
	JobsServiceRequestsListEntriesPerm             perms.Name = "RequestsListEntries"
	JobsServiceRequestsListEntriesAccessPermField  perms.Key  = "Access"
	JobsServiceRequestsPostCommentPerm             perms.Name = "RequestsPostComment"
	JobsServiceRequestsUpdateEntryPerm             perms.Name = "RequestsUpdateEntry"
	JobsServiceTimeclockListEntriesPerm            perms.Name = "TimeclockListEntries"
	JobsServiceTimeclockListEntriesAccessPermField perms.Key  = "Access"
)

var PermsRemap = map[string]string{

	// Service: JobsService
	"JobsService/RequestsListTypes": "JobsService/RequestsListEntries",
	"JobsService/TimeclockStats":    "JobsService/ConductListEntries",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: JobsService
		{
			Category: JobsServicePerm,
			Name:     JobsServiceColleaguesListPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceConductCreateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceConductDeleteEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceConductListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           JobsServiceConductListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceConductUpdateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsCreateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsCreateOrUpdateTypePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsDeleteCommentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsDeleteEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsDeleteTypePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           JobsServiceRequestsListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsPostCommentPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceRequestsUpdateEntryPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceTimeclockListEntriesPerm,
			Attrs: []perms.Attr{
				{
					Key:           JobsServiceTimeclockListEntriesAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "All"},
					DefaultValues: []string{"Own"},
				},
			},
		},
	})
}
