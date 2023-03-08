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

type ArpanetDocumentsTemplates struct {
	ID             uint64     `sql:"primary_key" json:"id"`
	CreatedAt      *time.Time `json:"created_at"`
	UpdatedAt      *time.Time `json:"updated_at"`
	Job            string     `json:"job"`
	JobGrade       int32      `json:"job_grade"`
	Title          string     `json:"title"`
	Description    string     `json:"description"`
	ContentTitle   string     `json:"content_title"`
	Content        string     `json:"content"`
	AdditionalData *string    `json:"additional_data"`
	CreatorID      int32      `json:"creator_id"`
}
