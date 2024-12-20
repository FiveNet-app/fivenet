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

var FivenetCalendarSubs = newFivenetCalendarSubsTable("", "fivenet_calendar_subs", "")

type fivenetCalendarSubsTable struct {
	mysql.Table

	// Columns
	CalendarID mysql.ColumnInteger
	EntryID    mysql.ColumnInteger
	UserID     mysql.ColumnInteger
	CreatedAt  mysql.ColumnTimestamp
	Confirmed  mysql.ColumnBool
	Muted      mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetCalendarSubsTable struct {
	fivenetCalendarSubsTable

	NEW fivenetCalendarSubsTable
}

// AS creates new FivenetCalendarSubsTable with assigned alias
func (a FivenetCalendarSubsTable) AS(alias string) *FivenetCalendarSubsTable {
	return newFivenetCalendarSubsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetCalendarSubsTable with assigned schema name
func (a FivenetCalendarSubsTable) FromSchema(schemaName string) *FivenetCalendarSubsTable {
	return newFivenetCalendarSubsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetCalendarSubsTable with assigned table prefix
func (a FivenetCalendarSubsTable) WithPrefix(prefix string) *FivenetCalendarSubsTable {
	return newFivenetCalendarSubsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetCalendarSubsTable with assigned table suffix
func (a FivenetCalendarSubsTable) WithSuffix(suffix string) *FivenetCalendarSubsTable {
	return newFivenetCalendarSubsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetCalendarSubsTable(schemaName, tableName, alias string) *FivenetCalendarSubsTable {
	return &FivenetCalendarSubsTable{
		fivenetCalendarSubsTable: newFivenetCalendarSubsTableImpl(schemaName, tableName, alias),
		NEW:                      newFivenetCalendarSubsTableImpl("", "new", ""),
	}
}

func newFivenetCalendarSubsTableImpl(schemaName, tableName, alias string) fivenetCalendarSubsTable {
	var (
		CalendarIDColumn = mysql.IntegerColumn("calendar_id")
		EntryIDColumn    = mysql.IntegerColumn("entry_id")
		UserIDColumn     = mysql.IntegerColumn("user_id")
		CreatedAtColumn  = mysql.TimestampColumn("created_at")
		ConfirmedColumn  = mysql.BoolColumn("confirmed")
		MutedColumn      = mysql.BoolColumn("muted")
		allColumns       = mysql.ColumnList{CalendarIDColumn, EntryIDColumn, UserIDColumn, CreatedAtColumn, ConfirmedColumn, MutedColumn}
		mutableColumns   = mysql.ColumnList{CalendarIDColumn, EntryIDColumn, UserIDColumn, CreatedAtColumn, ConfirmedColumn, MutedColumn}
	)

	return fivenetCalendarSubsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		CalendarID: CalendarIDColumn,
		EntryID:    EntryIDColumn,
		UserID:     UserIDColumn,
		CreatedAt:  CreatedAtColumn,
		Confirmed:  ConfirmedColumn,
		Muted:      MutedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
