// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/centrum/centrum.proto

package centrum

import (
	permkeys "github.com/fivenet-app/fivenet/gen/go/proto/services/centrum/perms"
	"github.com/fivenet-app/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{
	// Service: CentrumService
	"CentrumService/AssignDispatch":       "CentrumService/TakeControl",
	"CentrumService/AssignUnit":           "CentrumService/TakeControl",
	"CentrumService/GetDispatch":          "CentrumService/Stream",
	"CentrumService/GetSettings":          "CentrumService/Stream",
	"CentrumService/JoinUnit":             "CentrumService/Stream",
	"CentrumService/ListDispatchActivity": "CentrumService/Stream",
	"CentrumService/ListDispatches":       "CentrumService/Stream",
	"CentrumService/ListUnitActivity":     "CentrumService/Stream",
	"CentrumService/ListUnits":            "CentrumService/Stream",
	"CentrumService/UpdateDispatchStatus": "CentrumService/TakeDispatch",
	"CentrumService/UpdateUnitStatus":     "CentrumService/TakeDispatch",
}

func init() {
	perms.AddPermsToList([]*perms.Perm{

		// Service: CentrumService
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceCreateDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceCreateOrUpdateUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceDeleteDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceDeleteUnitPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceStreamPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceTakeControlPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceTakeDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceUpdateDispatchPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.CentrumServicePerm,
			Name:     permkeys.CentrumServiceUpdateSettingsPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
