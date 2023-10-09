// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/rector.proto

package rector

import (
	permkeys "github.com/galexrt/fivenet/gen/go/proto/services/rector/perms"
	"github.com/galexrt/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: RectorService
	"RectorService/CreateOrUpdateLaw":     "SuperUser",
	"RectorService/CreateOrUpdateLawBook": "SuperUser",
	"RectorService/DeleteLaw":             "SuperUser",
	"RectorService/DeleteLawBook":         "SuperUser",
	"RectorService/GetPermissions":        "RectorService/GetRoles",
	"RectorService/GetRole":               "RectorService/GetRoles",
	"RectorService/UpdateRoleLimits":      "SuperUser",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: RectorService
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceCreateRolePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceDeleteRolePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceGetJobPropsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceGetRolesPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceSetJobPropsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceUpdateRolePermsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.RectorServicePerm,
			Name:     permkeys.RectorServiceViewAuditLogPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
