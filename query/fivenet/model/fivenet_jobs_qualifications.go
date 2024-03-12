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

type FivenetJobsQualifications struct {
	ID              uint64     `sql:"primary_key" json:"id"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at"`
	Job             string     `json:"job"`
	Weight          *uint32    `json:"weight"`
	Closed          *bool      `json:"closed"`
	Abbreviation    string     `json:"abbreviation"`
	Title           *string    `json:"title"`
	Summary         *string    `json:"summary"`
	Description     *string    `json:"description"`
	CreatorID       *int32     `json:"creator_id"`
	CreatorJob      string     `json:"creator_job"`
	DiscordSettings *string    `json:"discord_settings"`
}
