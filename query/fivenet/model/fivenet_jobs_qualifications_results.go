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

type FivenetJobsQualificationsResults struct {
	ID              uint64     `sql:"primary_key" json:"id"`
	CreatedAt       *time.Time `json:"created_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	QualificationID uint64     `json:"qualification_id"`
	UserID          int32      `json:"user_id"`
	Status          *int16     `json:"status"`
	Score           *int32     `json:"score"`
	Summary         *string    `json:"summary"`
	CreatorID       *int32     `json:"creator_id"`
	CreatorJob      string     `json:"creator_job"`
}
