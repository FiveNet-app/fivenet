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

type Users struct {
	ID          int32      `json:"id"`
	Identifier  string     `sql:"primary_key" json:"identifier"`
	Group       *string    `json:"group"`
	Job         *string    `json:"job"`
	JobGrade    *int32     `json:"job_grade"`
	Firstname   *string    `json:"firstname"`
	Lastname    *string    `json:"lastname"`
	Dateofbirth *string    `json:"dateofbirth"`
	Sex         *string    `json:"sex"`
	Height      *string    `json:"height"`
	PhoneNumber *string    `json:"phone_number"`
	Visum       *int32     `json:"visum"`
	Playtime    *int32     `json:"playtime"`
	LastSeen    *time.Time `json:"last_seen"`
}
