// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/rector.proto

package rector

import "github.com/galexrt/fivenet/pkg/perms"

var PermsRemap = map[string]string{
	// Service: RectorService
	"RectorService/GetRole": "RectorService/GetRoles",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

const (
	RectorServicePerm perms.Category = "RectorService"

	RectorServiceCreateRolePerm              perms.Name = "CreateRole"
	RectorServiceDeleteRolePerm              perms.Name = "DeleteRole"
	RectorServiceGetJobPropsPerm             perms.Name = "GetJobProps"
	RectorServiceGetPermissionsPerm          perms.Name = "GetPermissions"
	RectorServiceGetPermissionsJobsPermField perms.Key  = "Jobs"
	RectorServiceGetRolesPerm                perms.Name = "GetRoles"
	RectorServiceSetJobPropsPerm             perms.Name = "SetJobProps"
	RectorServiceUpdateRolePermsPerm         perms.Name = "UpdateRolePerms"
	RectorServiceViewAuditLogPerm            perms.Name = "ViewAuditLog"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: RectorService
		{
			Category: RectorServicePerm,
			Name:     RectorServiceCreateRolePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceDeleteRolePerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceGetJobPropsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceGetPermissionsPerm,
			Attrs: []perms.Attr{
				{
					Key:         RectorServiceGetPermissionsJobsPermField,
					Type:        perms.JobListAttributeType,
					ValidValues: "config.C.Game.Livemap.Jobs",
				},
			},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceGetRolesPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceSetJobPropsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceUpdateRolePermsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: RectorServicePerm,
			Name:     RectorServiceViewAuditLogPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
