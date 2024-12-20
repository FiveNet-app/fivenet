//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var JobGrades = newJobGradesTable("", "job_grades", "")

type jobGradesTable struct {
	mysql.Table

	// Columns
	JobName mysql.ColumnString
	Grade   mysql.ColumnInteger
	Name    mysql.ColumnString
	Label   mysql.ColumnString
	Salary  mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type JobGradesTable struct {
	jobGradesTable

	NEW jobGradesTable
}

// AS creates new JobGradesTable with assigned alias
func (a JobGradesTable) AS(alias string) *JobGradesTable {
	return newJobGradesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new JobGradesTable with assigned schema name
func (a JobGradesTable) FromSchema(schemaName string) *JobGradesTable {
	return newJobGradesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new JobGradesTable with assigned table prefix
func (a JobGradesTable) WithPrefix(prefix string) *JobGradesTable {
	return newJobGradesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new JobGradesTable with assigned table suffix
func (a JobGradesTable) WithSuffix(suffix string) *JobGradesTable {
	return newJobGradesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newJobGradesTable(schemaName, tableName, alias string) *JobGradesTable {
	return &JobGradesTable{
		jobGradesTable: newJobGradesTableImpl(schemaName, tableName, alias),
		NEW:            newJobGradesTableImpl("", "new", ""),
	}
}

func newJobGradesTableImpl(schemaName, tableName, alias string) jobGradesTable {
	var (
		JobNameColumn  = mysql.StringColumn("job_name")
		GradeColumn    = mysql.IntegerColumn("grade")
		NameColumn     = mysql.StringColumn("name")
		LabelColumn    = mysql.StringColumn("label")
		SalaryColumn   = mysql.IntegerColumn("salary")
		allColumns     = mysql.ColumnList{JobNameColumn, GradeColumn, NameColumn, LabelColumn, SalaryColumn}
		mutableColumns = mysql.ColumnList{NameColumn, LabelColumn, SalaryColumn}
	)

	return jobGradesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		JobName: JobNameColumn,
		Grade:   GradeColumn,
		Name:    NameColumn,
		Label:   LabelColumn,
		Salary:  SalaryColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
