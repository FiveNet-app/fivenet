//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type JobGrades struct {
	JobName string `sql:"primary_key" json:"job_name"`
	Grade   int32  `sql:"primary_key" json:"grade"`
	Name    string `json:"name"`
	Label   string `json:"label"`
	Salary  int32  `json:"salary"`
}
