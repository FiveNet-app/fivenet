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

type FivenetCalendar struct {
	ID          uint64     `sql:"primary_key" json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	Job         *string    `json:"job"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Public      *bool      `json:"public"`
	Closed      *bool      `json:"closed"`
	CreatorID   *int32     `json:"creator_id"`
	CreatorJob  string     `json:"creator_job"`
}
