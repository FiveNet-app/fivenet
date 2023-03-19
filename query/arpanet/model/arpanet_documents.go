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

type ArpanetDocuments struct {
	ID          uint64     `sql:"primary_key" json:"id"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	CategoryID  *uint64    `json:"category_id"`
	Title       string     `json:"title"`
	ContentType int16      `json:"content_type"`
	Content     string     `json:"content"`
	Data        *string    `json:"data"`
	CreatorID   int32      `json:"creator_id"`
	CreatorJob  string     `json:"creator_job"`
	State       string     `json:"state"`
	Closed      *bool      `json:"closed"`
	Public      bool       `json:"public"`
}
