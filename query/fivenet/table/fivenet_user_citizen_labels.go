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

var FivenetUserCitizenLabels = newFivenetUserCitizenLabelsTable("", "fivenet_user_citizen_labels", "")

type fivenetUserCitizenLabelsTable struct {
	mysql.Table

	// Columns
	UserID      mysql.ColumnInteger
	AttributeID mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetUserCitizenLabelsTable struct {
	fivenetUserCitizenLabelsTable

	NEW fivenetUserCitizenLabelsTable
}

// AS creates new FivenetUserCitizenLabelsTable with assigned alias
func (a FivenetUserCitizenLabelsTable) AS(alias string) *FivenetUserCitizenLabelsTable {
	return newFivenetUserCitizenLabelsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetUserCitizenLabelsTable with assigned schema name
func (a FivenetUserCitizenLabelsTable) FromSchema(schemaName string) *FivenetUserCitizenLabelsTable {
	return newFivenetUserCitizenLabelsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetUserCitizenLabelsTable with assigned table prefix
func (a FivenetUserCitizenLabelsTable) WithPrefix(prefix string) *FivenetUserCitizenLabelsTable {
	return newFivenetUserCitizenLabelsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetUserCitizenLabelsTable with assigned table suffix
func (a FivenetUserCitizenLabelsTable) WithSuffix(suffix string) *FivenetUserCitizenLabelsTable {
	return newFivenetUserCitizenLabelsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetUserCitizenLabelsTable(schemaName, tableName, alias string) *FivenetUserCitizenLabelsTable {
	return &FivenetUserCitizenLabelsTable{
		fivenetUserCitizenLabelsTable: newFivenetUserCitizenLabelsTableImpl(schemaName, tableName, alias),
		NEW:                           newFivenetUserCitizenLabelsTableImpl("", "new", ""),
	}
}

func newFivenetUserCitizenLabelsTableImpl(schemaName, tableName, alias string) fivenetUserCitizenLabelsTable {
	var (
		UserIDColumn      = mysql.IntegerColumn("user_id")
		AttributeIDColumn = mysql.IntegerColumn("attribute_id")
		allColumns        = mysql.ColumnList{UserIDColumn, AttributeIDColumn}
		mutableColumns    = mysql.ColumnList{UserIDColumn, AttributeIDColumn}
	)

	return fivenetUserCitizenLabelsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		UserID:      UserIDColumn,
		AttributeID: AttributeIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
