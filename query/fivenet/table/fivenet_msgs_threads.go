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

var FivenetMsgsThreads = newFivenetMsgsThreadsTable("", "fivenet_msgs_threads", "")

type fivenetMsgsThreadsTable struct {
	mysql.Table

	// Columns
	ID         mysql.ColumnInteger
	CreatedAt  mysql.ColumnTimestamp
	UpdatedAt  mysql.ColumnTimestamp
	DeletedAt  mysql.ColumnTimestamp
	Title      mysql.ColumnString
	CreatorJob mysql.ColumnString
	CreatorID  mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetMsgsThreadsTable struct {
	fivenetMsgsThreadsTable

	NEW fivenetMsgsThreadsTable
}

// AS creates new FivenetMsgsThreadsTable with assigned alias
func (a FivenetMsgsThreadsTable) AS(alias string) *FivenetMsgsThreadsTable {
	return newFivenetMsgsThreadsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetMsgsThreadsTable with assigned schema name
func (a FivenetMsgsThreadsTable) FromSchema(schemaName string) *FivenetMsgsThreadsTable {
	return newFivenetMsgsThreadsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetMsgsThreadsTable with assigned table prefix
func (a FivenetMsgsThreadsTable) WithPrefix(prefix string) *FivenetMsgsThreadsTable {
	return newFivenetMsgsThreadsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetMsgsThreadsTable with assigned table suffix
func (a FivenetMsgsThreadsTable) WithSuffix(suffix string) *FivenetMsgsThreadsTable {
	return newFivenetMsgsThreadsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetMsgsThreadsTable(schemaName, tableName, alias string) *FivenetMsgsThreadsTable {
	return &FivenetMsgsThreadsTable{
		fivenetMsgsThreadsTable: newFivenetMsgsThreadsTableImpl(schemaName, tableName, alias),
		NEW:                     newFivenetMsgsThreadsTableImpl("", "new", ""),
	}
}

func newFivenetMsgsThreadsTableImpl(schemaName, tableName, alias string) fivenetMsgsThreadsTable {
	var (
		IDColumn         = mysql.IntegerColumn("id")
		CreatedAtColumn  = mysql.TimestampColumn("created_at")
		UpdatedAtColumn  = mysql.TimestampColumn("updated_at")
		DeletedAtColumn  = mysql.TimestampColumn("deleted_at")
		TitleColumn      = mysql.StringColumn("title")
		CreatorJobColumn = mysql.StringColumn("creator_job")
		CreatorIDColumn  = mysql.IntegerColumn("creator_id")
		allColumns       = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, TitleColumn, CreatorJobColumn, CreatorIDColumn}
		mutableColumns   = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, TitleColumn, CreatorJobColumn, CreatorIDColumn}
	)

	return fivenetMsgsThreadsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,
		DeletedAt:  DeletedAtColumn,
		Title:      TitleColumn,
		CreatorJob: CreatorJobColumn,
		CreatorID:  CreatorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
