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

var FivenetMailerMessages = newFivenetMailerMessagesTable("", "fivenet_mailer_messages", "")

type fivenetMailerMessagesTable struct {
	mysql.Table

	// Columns
	ID        mysql.ColumnInteger
	ThreadID  mysql.ColumnInteger
	SenderID  mysql.ColumnInteger
	CreatedAt mysql.ColumnTimestamp
	UpdatedAt mysql.ColumnTimestamp
	DeletedAt mysql.ColumnTimestamp
	Title     mysql.ColumnString
	Content   mysql.ColumnString
	Data      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetMailerMessagesTable struct {
	fivenetMailerMessagesTable

	NEW fivenetMailerMessagesTable
}

// AS creates new FivenetMailerMessagesTable with assigned alias
func (a FivenetMailerMessagesTable) AS(alias string) *FivenetMailerMessagesTable {
	return newFivenetMailerMessagesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetMailerMessagesTable with assigned schema name
func (a FivenetMailerMessagesTable) FromSchema(schemaName string) *FivenetMailerMessagesTable {
	return newFivenetMailerMessagesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetMailerMessagesTable with assigned table prefix
func (a FivenetMailerMessagesTable) WithPrefix(prefix string) *FivenetMailerMessagesTable {
	return newFivenetMailerMessagesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetMailerMessagesTable with assigned table suffix
func (a FivenetMailerMessagesTable) WithSuffix(suffix string) *FivenetMailerMessagesTable {
	return newFivenetMailerMessagesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetMailerMessagesTable(schemaName, tableName, alias string) *FivenetMailerMessagesTable {
	return &FivenetMailerMessagesTable{
		fivenetMailerMessagesTable: newFivenetMailerMessagesTableImpl(schemaName, tableName, alias),
		NEW:                        newFivenetMailerMessagesTableImpl("", "new", ""),
	}
}

func newFivenetMailerMessagesTableImpl(schemaName, tableName, alias string) fivenetMailerMessagesTable {
	var (
		IDColumn        = mysql.IntegerColumn("id")
		ThreadIDColumn  = mysql.IntegerColumn("thread_id")
		SenderIDColumn  = mysql.IntegerColumn("sender_id")
		CreatedAtColumn = mysql.TimestampColumn("created_at")
		UpdatedAtColumn = mysql.TimestampColumn("updated_at")
		DeletedAtColumn = mysql.TimestampColumn("deleted_at")
		TitleColumn     = mysql.StringColumn("title")
		ContentColumn   = mysql.StringColumn("content")
		DataColumn      = mysql.StringColumn("data")
		allColumns      = mysql.ColumnList{IDColumn, ThreadIDColumn, SenderIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, TitleColumn, ContentColumn, DataColumn}
		mutableColumns  = mysql.ColumnList{ThreadIDColumn, SenderIDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, TitleColumn, ContentColumn, DataColumn}
	)

	return fivenetMailerMessagesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:        IDColumn,
		ThreadID:  ThreadIDColumn,
		SenderID:  SenderIDColumn,
		CreatedAt: CreatedAtColumn,
		UpdatedAt: UpdatedAtColumn,
		DeletedAt: DeletedAtColumn,
		Title:     TitleColumn,
		Content:   ContentColumn,
		Data:      DataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
