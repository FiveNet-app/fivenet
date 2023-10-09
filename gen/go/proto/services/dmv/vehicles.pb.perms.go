// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/dmv/vehicles.proto

package dmv

import (
	permkeys "github.com/galexrt/fivenet/gen/go/proto/services/dmv/perms"
	"github.com/galexrt/fivenet/pkg/perms"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: DMVService
		{
			Category: permkeys.DMVServicePerm,
			Name:     permkeys.DMVServiceListVehiclesPerm,
			Attrs:    []perms.Attr{},
		},
	})
}
