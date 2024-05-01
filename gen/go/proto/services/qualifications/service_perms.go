// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/qualifications/qualifications.proto

package qualifications

import (
	"github.com/fivenet-app/fivenet/gen/go/proto/resources/permissions"
	permkeys "github.com/fivenet-app/fivenet/gen/go/proto/services/qualifications/perms"
	"github.com/fivenet-app/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: QualificationsService
	"QualificationsService/CreateOrUpdateQualificationRequest": "QualificationsService/GetQualification",
	"QualificationsService/ListQualificationRequests":          "QualificationsService/GetQualification",
	"QualificationsService/ListQualificationsResults":          "QualificationsService/GetQualification",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{

		// Service: QualificationsService
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceCreateOrUpdateQualificationResultPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceCreateQualificationPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceDeleteQualificationPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.QualificationsServiceDeleteQualificationAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceDeleteQualificationReqPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceDeleteQualificationResultPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceGetQualificationPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceListQualificationsPerm,
			Attrs:    []perms.Attr{},
		},
		{
			Category: permkeys.QualificationsServicePerm,
			Name:     permkeys.QualificationsServiceUpdateQualificationPerm,
			Attrs: []perms.Attr{
				{
					Key:         permkeys.QualificationsServiceUpdateQualificationAccessPermField,
					Type:        permissions.StringListAttributeType,
					ValidValues: []string{"Own", "Lower_Rank", "Same_Rank", "Any"},
				},
			},
		},
	})
}
