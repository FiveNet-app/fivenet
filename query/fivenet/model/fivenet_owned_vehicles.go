//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FivenetOwnedVehicles struct {
	Owner *string `json:"owner"`
	Plate string  `sql:"primary_key" json:"plate"`
	Model *string `json:"model"`
	Type  string  `json:"type"`
}