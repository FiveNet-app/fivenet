// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/jobs.proto

package jobs

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	JobsServicePerm perms.Category = "JobsService"

	JobsServiceColleaguesListPerm                perms.Name = "ColleaguesList"
	JobsServiceConductCreateEntryPerm            perms.Name = "ConductCreateEntry"
	JobsServiceConductDeleteEntryPerm            perms.Name = "ConductDeleteEntry"
	JobsServiceConductListEntriesPerm            perms.Name = "ConductListEntries"
	JobsServiceConductListEntriesAccessPermField perms.Key  = "Access"
	JobsServiceConductUpdateEntryPerm            perms.Name = "ConductUpdateEntry"
)

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
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: JobsServicePerm,
			Name:     JobsServiceConductUpdateEntryPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
