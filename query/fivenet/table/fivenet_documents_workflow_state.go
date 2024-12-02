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

var FivenetDocumentsWorkflowState = newFivenetDocumentsWorkflowStateTable("", "fivenet_documents_workflow_state", "")

type fivenetDocumentsWorkflowStateTable struct {
	mysql.Table

	// Columns
	DocumentID        mysql.ColumnInteger
	NextReminderTime  mysql.ColumnTimestamp
	NextReminderCount mysql.ColumnInteger
	AutoCloseTime     mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetDocumentsWorkflowStateTable struct {
	fivenetDocumentsWorkflowStateTable

	NEW fivenetDocumentsWorkflowStateTable
}

// AS creates new FivenetDocumentsWorkflowStateTable with assigned alias
func (a FivenetDocumentsWorkflowStateTable) AS(alias string) *FivenetDocumentsWorkflowStateTable {
	return newFivenetDocumentsWorkflowStateTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetDocumentsWorkflowStateTable with assigned schema name
func (a FivenetDocumentsWorkflowStateTable) FromSchema(schemaName string) *FivenetDocumentsWorkflowStateTable {
	return newFivenetDocumentsWorkflowStateTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetDocumentsWorkflowStateTable with assigned table prefix
func (a FivenetDocumentsWorkflowStateTable) WithPrefix(prefix string) *FivenetDocumentsWorkflowStateTable {
	return newFivenetDocumentsWorkflowStateTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetDocumentsWorkflowStateTable with assigned table suffix
func (a FivenetDocumentsWorkflowStateTable) WithSuffix(suffix string) *FivenetDocumentsWorkflowStateTable {
	return newFivenetDocumentsWorkflowStateTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetDocumentsWorkflowStateTable(schemaName, tableName, alias string) *FivenetDocumentsWorkflowStateTable {
	return &FivenetDocumentsWorkflowStateTable{
		fivenetDocumentsWorkflowStateTable: newFivenetDocumentsWorkflowStateTableImpl(schemaName, tableName, alias),
		NEW:                                newFivenetDocumentsWorkflowStateTableImpl("", "new", ""),
	}
}

func newFivenetDocumentsWorkflowStateTableImpl(schemaName, tableName, alias string) fivenetDocumentsWorkflowStateTable {
	var (
		DocumentIDColumn        = mysql.IntegerColumn("document_id")
		NextReminderTimeColumn  = mysql.TimestampColumn("next_reminder_time")
		NextReminderCountColumn = mysql.IntegerColumn("next_reminder_count")
		AutoCloseTimeColumn     = mysql.TimestampColumn("auto_close_time")
		allColumns              = mysql.ColumnList{DocumentIDColumn, NextReminderTimeColumn, NextReminderCountColumn, AutoCloseTimeColumn}
		mutableColumns          = mysql.ColumnList{DocumentIDColumn, NextReminderTimeColumn, NextReminderCountColumn, AutoCloseTimeColumn}
	)

	return fivenetDocumentsWorkflowStateTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		DocumentID:        DocumentIDColumn,
		NextReminderTime:  NextReminderTimeColumn,
		NextReminderCount: NextReminderCountColumn,
		AutoCloseTime:     AutoCloseTimeColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
