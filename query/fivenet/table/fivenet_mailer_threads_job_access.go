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

var FivenetMailerThreadsJobAccess = newFivenetMailerThreadsJobAccessTable("", "fivenet_mailer_threads_job_access", "")

type fivenetMailerThreadsJobAccessTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	AddressID    mysql.ColumnInteger
	Job          mysql.ColumnString
	MinimumGrade mysql.ColumnInteger
	Access       mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetMailerThreadsJobAccessTable struct {
	fivenetMailerThreadsJobAccessTable

	NEW fivenetMailerThreadsJobAccessTable
}

// AS creates new FivenetMailerThreadsJobAccessTable with assigned alias
func (a FivenetMailerThreadsJobAccessTable) AS(alias string) *FivenetMailerThreadsJobAccessTable {
	return newFivenetMailerThreadsJobAccessTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetMailerThreadsJobAccessTable with assigned schema name
func (a FivenetMailerThreadsJobAccessTable) FromSchema(schemaName string) *FivenetMailerThreadsJobAccessTable {
	return newFivenetMailerThreadsJobAccessTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetMailerThreadsJobAccessTable with assigned table prefix
func (a FivenetMailerThreadsJobAccessTable) WithPrefix(prefix string) *FivenetMailerThreadsJobAccessTable {
	return newFivenetMailerThreadsJobAccessTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetMailerThreadsJobAccessTable with assigned table suffix
func (a FivenetMailerThreadsJobAccessTable) WithSuffix(suffix string) *FivenetMailerThreadsJobAccessTable {
	return newFivenetMailerThreadsJobAccessTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetMailerThreadsJobAccessTable(schemaName, tableName, alias string) *FivenetMailerThreadsJobAccessTable {
	return &FivenetMailerThreadsJobAccessTable{
		fivenetMailerThreadsJobAccessTable: newFivenetMailerThreadsJobAccessTableImpl(schemaName, tableName, alias),
		NEW:                                newFivenetMailerThreadsJobAccessTableImpl("", "new", ""),
	}
}

func newFivenetMailerThreadsJobAccessTableImpl(schemaName, tableName, alias string) fivenetMailerThreadsJobAccessTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		AddressIDColumn    = mysql.IntegerColumn("address_id")
		JobColumn          = mysql.StringColumn("job")
		MinimumGradeColumn = mysql.IntegerColumn("minimum_grade")
		AccessColumn       = mysql.IntegerColumn("access")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, AddressIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, AddressIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
	)

	return fivenetMailerThreadsJobAccessTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		AddressID:    AddressIDColumn,
		Job:          JobColumn,
		MinimumGrade: MinimumGradeColumn,
		Access:       AccessColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
