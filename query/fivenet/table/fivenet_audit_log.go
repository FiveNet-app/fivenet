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

var FivenetAuditLog = newFivenetAuditLogTable("", "fivenet_audit_log", "")

type fivenetAuditLogTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	UserID       mysql.ColumnInteger
	UserJob      mysql.ColumnString
	TargetUserID mysql.ColumnInteger
	Service      mysql.ColumnString
	Method       mysql.ColumnString
	State        mysql.ColumnInteger
	Data         mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetAuditLogTable struct {
	fivenetAuditLogTable

	NEW fivenetAuditLogTable
}

// AS creates new FivenetAuditLogTable with assigned alias
func (a FivenetAuditLogTable) AS(alias string) *FivenetAuditLogTable {
	return newFivenetAuditLogTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetAuditLogTable with assigned schema name
func (a FivenetAuditLogTable) FromSchema(schemaName string) *FivenetAuditLogTable {
	return newFivenetAuditLogTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetAuditLogTable with assigned table prefix
func (a FivenetAuditLogTable) WithPrefix(prefix string) *FivenetAuditLogTable {
	return newFivenetAuditLogTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetAuditLogTable with assigned table suffix
func (a FivenetAuditLogTable) WithSuffix(suffix string) *FivenetAuditLogTable {
	return newFivenetAuditLogTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetAuditLogTable(schemaName, tableName, alias string) *FivenetAuditLogTable {
	return &FivenetAuditLogTable{
		fivenetAuditLogTable: newFivenetAuditLogTableImpl(schemaName, tableName, alias),
		NEW:                  newFivenetAuditLogTableImpl("", "new", ""),
	}
}

func newFivenetAuditLogTableImpl(schemaName, tableName, alias string) fivenetAuditLogTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		UserIDColumn       = mysql.IntegerColumn("user_id")
		UserJobColumn      = mysql.StringColumn("user_job")
		TargetUserIDColumn = mysql.IntegerColumn("target_user_id")
		ServiceColumn      = mysql.StringColumn("service")
		MethodColumn       = mysql.StringColumn("method")
		StateColumn        = mysql.IntegerColumn("state")
		DataColumn         = mysql.StringColumn("data")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, UserIDColumn, UserJobColumn, TargetUserIDColumn, ServiceColumn, MethodColumn, StateColumn, DataColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, UserIDColumn, UserJobColumn, TargetUserIDColumn, ServiceColumn, MethodColumn, StateColumn, DataColumn}
	)

	return fivenetAuditLogTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		UserID:       UserIDColumn,
		UserJob:      UserJobColumn,
		TargetUserID: TargetUserIDColumn,
		Service:      ServiceColumn,
		Method:       MethodColumn,
		State:        StateColumn,
		Data:         DataColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}