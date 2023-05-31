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

var FivenetUserProps = newFivenetUserPropsTable("", "fivenet_user_props", "")

type fivenetUserPropsTable struct {
	mysql.Table

	// Columns
	UserID mysql.ColumnInteger
	Wanted mysql.ColumnBool
	Job    mysql.ColumnString
	Grade  mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetUserPropsTable struct {
	fivenetUserPropsTable

	NEW fivenetUserPropsTable
}

// AS creates new FivenetUserPropsTable with assigned alias
func (a FivenetUserPropsTable) AS(alias string) *FivenetUserPropsTable {
	return newFivenetUserPropsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetUserPropsTable with assigned schema name
func (a FivenetUserPropsTable) FromSchema(schemaName string) *FivenetUserPropsTable {
	return newFivenetUserPropsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetUserPropsTable with assigned table prefix
func (a FivenetUserPropsTable) WithPrefix(prefix string) *FivenetUserPropsTable {
	return newFivenetUserPropsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetUserPropsTable with assigned table suffix
func (a FivenetUserPropsTable) WithSuffix(suffix string) *FivenetUserPropsTable {
	return newFivenetUserPropsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetUserPropsTable(schemaName, tableName, alias string) *FivenetUserPropsTable {
	return &FivenetUserPropsTable{
		fivenetUserPropsTable: newFivenetUserPropsTableImpl(schemaName, tableName, alias),
		NEW:                   newFivenetUserPropsTableImpl("", "new", ""),
	}
}

func newFivenetUserPropsTableImpl(schemaName, tableName, alias string) fivenetUserPropsTable {
	var (
		UserIDColumn   = mysql.IntegerColumn("user_id")
		WantedColumn   = mysql.BoolColumn("wanted")
		JobColumn      = mysql.StringColumn("job")
		GradeColumn    = mysql.IntegerColumn("grade")
		allColumns     = mysql.ColumnList{UserIDColumn, WantedColumn, JobColumn, GradeColumn}
		mutableColumns = mysql.ColumnList{UserIDColumn, WantedColumn, JobColumn, GradeColumn}
	)

	return fivenetUserPropsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID: UserIDColumn,
		Wanted: WantedColumn,
		Job:    JobColumn,
		Grade:  GradeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
