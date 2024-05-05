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

type FivenetCalendarEntries struct {
	ID         uint64     `sql:"primary_key" json:"id"`
	CreatedAt  *time.Time `json:"created_at"`
	UpdatedAt  *time.Time `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	CalendarID uint64     `json:"calendar_id"`
	Job        *string    `json:"job"`
	StartTime  time.Time  `json:"start_time"`
	EndTime    *time.Time `json:"end_time"`
	Title      string     `json:"title"`
	Content    *string    `json:"content"`
	RsvpOpen   *bool      `json:"rsvp_open"`
	CreatorID  *int32     `json:"creator_id"`
	CreatorJob string     `json:"creator_job"`
}
