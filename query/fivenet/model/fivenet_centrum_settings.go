//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FivenetCentrumSettings struct {
	Job          string `sql:"primary_key" json:"job"`
	Enabled      bool   `json:"enabled"`
	Active       *bool  `json:"active"`
	Mode         *int32 `json:"mode"`
	FallbackMode *int32 `json:"fallback_mode"`
}
