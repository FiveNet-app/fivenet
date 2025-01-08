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

var FivenetJobs = newFivenetJobsTable("", "fivenet_jobs", "")

type fivenetJobsTable struct {
	mysql.Table

	// Columns
	Name  mysql.ColumnString
	Label mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetJobsTable struct {
	fivenetJobsTable

	NEW fivenetJobsTable
}

// AS creates new FivenetJobsTable with assigned alias
func (a FivenetJobsTable) AS(alias string) *FivenetJobsTable {
	return newFivenetJobsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetJobsTable with assigned schema name
func (a FivenetJobsTable) FromSchema(schemaName string) *FivenetJobsTable {
	return newFivenetJobsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetJobsTable with assigned table prefix
func (a FivenetJobsTable) WithPrefix(prefix string) *FivenetJobsTable {
	return newFivenetJobsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetJobsTable with assigned table suffix
func (a FivenetJobsTable) WithSuffix(suffix string) *FivenetJobsTable {
	return newFivenetJobsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetJobsTable(schemaName, tableName, alias string) *FivenetJobsTable {
	return &FivenetJobsTable{
		fivenetJobsTable: newFivenetJobsTableImpl(schemaName, tableName, alias),
		NEW:              newFivenetJobsTableImpl("", "new", ""),
	}
}

func newFivenetJobsTableImpl(schemaName, tableName, alias string) fivenetJobsTable {
	var (
		NameColumn     = mysql.StringColumn("name")
		LabelColumn    = mysql.StringColumn("label")
		allColumns     = mysql.ColumnList{NameColumn, LabelColumn}
		mutableColumns = mysql.ColumnList{LabelColumn}
	)

	return fivenetJobsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		Name:  NameColumn,
		Label: LabelColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
