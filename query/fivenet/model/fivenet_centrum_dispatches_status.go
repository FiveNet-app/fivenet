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

type FivenetCentrumDispatchesStatus struct {
	ID         uint64     `sql:"primary_key" json:"id"`
	CreatedAt  *time.Time `json:"created_at"`
	DispatchID uint64     `json:"dispatch_id"`
	UnitID     *uint64    `json:"unit_id"`
	Status     int16      `json:"status"`
	Reason     *string    `json:"reason"`
	Code       *string    `json:"code"`
	X          *float64   `json:"x"`
	Y          *float64   `json:"y"`
	UserID     *int32     `json:"user_id"`
}
