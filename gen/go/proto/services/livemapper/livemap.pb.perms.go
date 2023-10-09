// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/livemapper/livemap.proto

package livemapper

import (
	"github.com/galexrt/fivenet/gen/go/proto/resources/permissions"
	permkeys "github.com/galexrt/fivenet/gen/go/proto/services/livemapper/perms"
	"github.com/galexrt/fivenet/pkg/perms"
)

func init() {
	perms.AddPermsToList([]*perms.Perm{
		// Service: LivemapperService
		{
			Category: permkeys.LivemapperServicePerm,
			Name:     permkeys.LivemapperServiceCreateOrUpdateMarkerPerm,
			Attrs: []perms.Attr{
				{
					Key:           permkeys.LivemapperServiceCreateOrUpdateMarkerAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: permkeys.LivemapperServicePerm,
			Name:     permkeys.LivemapperServiceDeleteMarkerPerm,
			Attrs: []perms.Attr{
				{
					Key:           permkeys.LivemapperServiceDeleteMarkerAccessPermField,
					Type:          permissions.StringListAttributeType,
					ValidValues:   []string{"Own", "Lower_Rank", "Same_Rank"},
					DefaultValues: []string{"Own"},
				},
			},
		},
		{
			Category: permkeys.LivemapperServicePerm,
			Name:     permkeys.LivemapperServiceStreamPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.LivemapperServiceStreamMarkersPermField,
					Type:        permissions.JobListAttributeType,
					ValidValues: "config.Game.Livemap.Jobs",
				},
				{
					Key:  permkeys.LivemapperServiceStreamPlayersPermField,
					Type: permissions.JobGradeListAttributeType,
				},
			},
		},
	})
}
