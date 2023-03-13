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

var ArpanetUserActivity = newArpanetUserActivityTable("arpanet", "arpanet_user_activity", "")

type arpanetUserActivityTable struct {
	mysql.Table

	//Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	UpdatedAt    mysql.ColumnTimestamp
	SourceUserID mysql.ColumnInteger
	TargetUserID mysql.ColumnInteger
	Type         mysql.ColumnInteger
	Key          mysql.ColumnString
	OldValue     mysql.ColumnString
	NewValue     mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type ArpanetUserActivityTable struct {
	arpanetUserActivityTable

	NEW arpanetUserActivityTable
}

// AS creates new ArpanetUserActivityTable with assigned alias
func (a ArpanetUserActivityTable) AS(alias string) *ArpanetUserActivityTable {
	return newArpanetUserActivityTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new ArpanetUserActivityTable with assigned schema name
func (a ArpanetUserActivityTable) FromSchema(schemaName string) *ArpanetUserActivityTable {
	return newArpanetUserActivityTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new ArpanetUserActivityTable with assigned table prefix
func (a ArpanetUserActivityTable) WithPrefix(prefix string) *ArpanetUserActivityTable {
	return newArpanetUserActivityTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new ArpanetUserActivityTable with assigned table suffix
func (a ArpanetUserActivityTable) WithSuffix(suffix string) *ArpanetUserActivityTable {
	return newArpanetUserActivityTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newArpanetUserActivityTable(schemaName, tableName, alias string) *ArpanetUserActivityTable {
	return &ArpanetUserActivityTable{
		arpanetUserActivityTable: newArpanetUserActivityTableImpl(schemaName, tableName, alias),
		NEW:                      newArpanetUserActivityTableImpl("", "new", ""),
	}
}

func newArpanetUserActivityTableImpl(schemaName, tableName, alias string) arpanetUserActivityTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		UpdatedAtColumn    = mysql.TimestampColumn("updated_at")
		SourceUserIDColumn = mysql.IntegerColumn("source_user_id")
		TargetUserIDColumn = mysql.IntegerColumn("target_user_id")
		TypeColumn         = mysql.IntegerColumn("type")
		KeyColumn          = mysql.StringColumn("key")
		OldValueColumn     = mysql.StringColumn("old_value")
		NewValueColumn     = mysql.StringColumn("new_value")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, SourceUserIDColumn, TargetUserIDColumn, TypeColumn, KeyColumn, OldValueColumn, NewValueColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, SourceUserIDColumn, TargetUserIDColumn, TypeColumn, KeyColumn, OldValueColumn, NewValueColumn}
	)

	return arpanetUserActivityTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		UpdatedAt:    UpdatedAtColumn,
		SourceUserID: SourceUserIDColumn,
		TargetUserID: TargetUserIDColumn,
		Type:         TypeColumn,
		Key:          KeyColumn,
		OldValue:     OldValueColumn,
		NewValue:     NewValueColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
