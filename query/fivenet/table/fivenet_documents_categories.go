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

var FivenetDocumentsCategories = newFivenetDocumentsCategoriesTable("", "fivenet_documents_categories", "")

type fivenetDocumentsCategoriesTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	Name        mysql.ColumnString
	Description mysql.ColumnString
	Job         mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetDocumentsCategoriesTable struct {
	fivenetDocumentsCategoriesTable

	NEW fivenetDocumentsCategoriesTable
}

// AS creates new FivenetDocumentsCategoriesTable with assigned alias
func (a FivenetDocumentsCategoriesTable) AS(alias string) *FivenetDocumentsCategoriesTable {
	return newFivenetDocumentsCategoriesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetDocumentsCategoriesTable with assigned schema name
func (a FivenetDocumentsCategoriesTable) FromSchema(schemaName string) *FivenetDocumentsCategoriesTable {
	return newFivenetDocumentsCategoriesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetDocumentsCategoriesTable with assigned table prefix
func (a FivenetDocumentsCategoriesTable) WithPrefix(prefix string) *FivenetDocumentsCategoriesTable {
	return newFivenetDocumentsCategoriesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetDocumentsCategoriesTable with assigned table suffix
func (a FivenetDocumentsCategoriesTable) WithSuffix(suffix string) *FivenetDocumentsCategoriesTable {
	return newFivenetDocumentsCategoriesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetDocumentsCategoriesTable(schemaName, tableName, alias string) *FivenetDocumentsCategoriesTable {
	return &FivenetDocumentsCategoriesTable{
		fivenetDocumentsCategoriesTable: newFivenetDocumentsCategoriesTableImpl(schemaName, tableName, alias),
		NEW:                             newFivenetDocumentsCategoriesTableImpl("", "new", ""),
	}
}

func newFivenetDocumentsCategoriesTableImpl(schemaName, tableName, alias string) fivenetDocumentsCategoriesTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		NameColumn        = mysql.StringColumn("name")
		DescriptionColumn = mysql.StringColumn("description")
		JobColumn         = mysql.StringColumn("job")
		allColumns        = mysql.ColumnList{IDColumn, NameColumn, DescriptionColumn, JobColumn}
		mutableColumns    = mysql.ColumnList{NameColumn, DescriptionColumn, JobColumn}
	)

	return fivenetDocumentsCategoriesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		Job:         JobColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
