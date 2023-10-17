// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/jobs/jobs.proto

package permsjobs

import (
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
	JobsServiceRequestsUpdateEntryPerm             perms.Name = "RequestsUpdateEntry"
	JobsServiceTimeclockListEntriesPerm            perms.Name = "TimeclockListEntries"
	JobsServiceTimeclockListEntriesAccessPermField perms.Key  = "Access"
)