// Code generated by protoc-gen-customizer. DO NOT EDIT.
// source: services/notificator/notificator.proto

package notificator

import (
	"github.com/fivenet-app/fivenet/pkg/perms"
)

var PermsRemap = map[string]string{

	// Service: NotificatorService
	"NotificatorService/GetNotifications":  "Any",
	"NotificatorService/MarkNotifications": "Any",
	"NotificatorService/Stream":            "Any",
}

func (s *Server) GetPermsRemap() map[string]string {
	return PermsRemap
}

func init() {
	perms.AddPermsToList([]*perms.Perm{})
}
