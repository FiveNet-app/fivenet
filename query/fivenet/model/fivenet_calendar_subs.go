//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"time"
)

type FivenetCalendarSubs struct {
	CalendarID uint64     `json:"calendar_id"`
	UserID     int32      `json:"user_id"`
	CreatedAt  *time.Time `json:"created_at"`
	Confirmed  bool       `json:"confirmed"`
	Muted      bool       `json:"muted"`
}
