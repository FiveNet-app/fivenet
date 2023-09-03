// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/centrum/centrum.proto

package centrum

import (
	"github.com/galexrt/fivenet/pkg/perms"
)

const (
	CentrumServicePerm perms.Category = "CentrumService"

	CentrumServiceCreateDispatchPerm     perms.Name = "CreateDispatch"
	CentrumServiceCreateOrUpdateUnitPerm perms.Name = "CreateOrUpdateUnit"
	CentrumServiceDeleteDispatchPerm     perms.Name = "DeleteDispatch"
	CentrumServiceDeleteUnitPerm         perms.Name = "DeleteUnit"
	CentrumServiceStreamPerm             perms.Name = "Stream"
	CentrumServiceTakeControlPerm        perms.Name = "TakeControl"
	CentrumServiceTakeDispatchPerm       perms.Name = "TakeDispatch"
	CentrumServiceUpdateDispatchPerm     perms.Name = "UpdateDispatch"
	CentrumServiceUpdateSettingsPerm     perms.Name = "UpdateSettings"
)

var PermsRemap = map[string]string{

	// Service: CentrumService
	"CentrumService/AssignDispatch":       "CentrumService/TakeControl",
	"CentrumService/AssignUnit":           "CentrumService/TakeControl",
	"CentrumService/GetSettings":          "CentrumService/Stream",
	"CentrumService/JoinUnit":             "CentrumService/Stream",
	"CentrumService/ListDispatchActivity": "CentrumService/Stream",
	"CentrumService/ListDispatches":       "CentrumService/Stream",
	"CentrumService/ListUnitActivity":     "CentrumService/Stream",
	"CentrumService/ListUnits":            "CentrumService/Stream",
	"CentrumService/UpdateDispatchStatus": "CentrumService/TakeDispatch",
	"CentrumService/UpdateUnitStatus":     "CentrumService/TakeDispatch",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: CentrumService
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceCreateDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceCreateOrUpdateUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceDeleteDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceDeleteUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceStreamPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceTakeControlPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceTakeDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceUpdateDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: CentrumServicePerm,
			Name:     CentrumServiceUpdateSettingsPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
