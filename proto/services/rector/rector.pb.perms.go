// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/rector/rector.proto

package rector

import "github.com/galexrt/fivenet/pkg/perms"

var PermsRemap = map[string]string{
	// Service: RectorService
	"RectorService/GetRole":            "RectorService/GetRoles",
	"RectorService/RemovePermFromRole": "RectorService/AddPermToRole",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

const (
	RectorServicePermKey = "RectorService"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: RectorService
		{
			Key:  RectorServicePermKey,
			Name: "AddPermToRole",
		},
		{
			Key:  RectorServicePermKey,
			Name: "DeleteRole",
		},
		{
			Key:    RectorServicePermKey,
			Name:   "GetPermissions",
			PerJob: true,
		},
		{
			Key:  RectorServicePermKey,
			Name: "GetRoles",
		},
	})
}
