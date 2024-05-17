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

var FivenetMsgsThreadsUserState = newFivenetMsgsThreadsUserStateTable("", "fivenet_msgs_threads_user_state", "")

type fivenetMsgsThreadsUserStateTable struct {
	mysql.Table

	// Columns
	ThreadID  mysql.ColumnInteger
	UserID    mysql.ColumnInteger
	LastRead  mysql.ColumnInteger
	Important mysql.ColumnBool
	Favorite  mysql.ColumnBool
	Muted     mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetMsgsThreadsUserStateTable struct {
	fivenetMsgsThreadsUserStateTable

	NEW fivenetMsgsThreadsUserStateTable
}

// AS creates new FivenetMsgsThreadsUserStateTable with assigned alias
func (a FivenetMsgsThreadsUserStateTable) AS(alias string) *FivenetMsgsThreadsUserStateTable {
	return newFivenetMsgsThreadsUserStateTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetMsgsThreadsUserStateTable with assigned schema name
func (a FivenetMsgsThreadsUserStateTable) FromSchema(schemaName string) *FivenetMsgsThreadsUserStateTable {
	return newFivenetMsgsThreadsUserStateTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetMsgsThreadsUserStateTable with assigned table prefix
func (a FivenetMsgsThreadsUserStateTable) WithPrefix(prefix string) *FivenetMsgsThreadsUserStateTable {
	return newFivenetMsgsThreadsUserStateTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetMsgsThreadsUserStateTable with assigned table suffix
func (a FivenetMsgsThreadsUserStateTable) WithSuffix(suffix string) *FivenetMsgsThreadsUserStateTable {
	return newFivenetMsgsThreadsUserStateTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetMsgsThreadsUserStateTable(schemaName, tableName, alias string) *FivenetMsgsThreadsUserStateTable {
	return &FivenetMsgsThreadsUserStateTable{
		fivenetMsgsThreadsUserStateTable: newFivenetMsgsThreadsUserStateTableImpl(schemaName, tableName, alias),
		NEW:                              newFivenetMsgsThreadsUserStateTableImpl("", "new", ""),
	}
}

func newFivenetMsgsThreadsUserStateTableImpl(schemaName, tableName, alias string) fivenetMsgsThreadsUserStateTable {
	var (
		ThreadIDColumn  = mysql.IntegerColumn("thread_id")
		UserIDColumn    = mysql.IntegerColumn("user_id")
		LastReadColumn  = mysql.IntegerColumn("last_read")
		ImportantColumn = mysql.BoolColumn("important")
		FavoriteColumn  = mysql.BoolColumn("favorite")
		MutedColumn     = mysql.BoolColumn("muted")
		allColumns      = mysql.ColumnList{ThreadIDColumn, UserIDColumn, LastReadColumn, ImportantColumn, FavoriteColumn, MutedColumn}
		mutableColumns  = mysql.ColumnList{LastReadColumn, ImportantColumn, FavoriteColumn, MutedColumn}
	)

	return fivenetMsgsThreadsUserStateTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ThreadID:  ThreadIDColumn,
		UserID:    UserIDColumn,
		LastRead:  LastReadColumn,
		Important: ImportantColumn,
		Favorite:  FavoriteColumn,
		Muted:     MutedColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
