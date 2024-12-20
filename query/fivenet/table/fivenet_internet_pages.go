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

var FivenetInternetPages = newFivenetInternetPagesTable("", "fivenet_internet_pages", "")

type fivenetInternetPagesTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	CreatedAt   mysql.ColumnTimestamp
	UpdatedAt   mysql.ColumnTimestamp
	DeletedAt   mysql.ColumnTimestamp
	DomainID    mysql.ColumnInteger
	Path        mysql.ColumnString
	Title       mysql.ColumnString
	Description mysql.ColumnString
	Data        mysql.ColumnString
	CreatorJob  mysql.ColumnString
	CreatorID   mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetInternetPagesTable struct {
	fivenetInternetPagesTable

	NEW fivenetInternetPagesTable
}

// AS creates new FivenetInternetPagesTable with assigned alias
func (a FivenetInternetPagesTable) AS(alias string) *FivenetInternetPagesTable {
	return newFivenetInternetPagesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetInternetPagesTable with assigned schema name
func (a FivenetInternetPagesTable) FromSchema(schemaName string) *FivenetInternetPagesTable {
	return newFivenetInternetPagesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetInternetPagesTable with assigned table prefix
func (a FivenetInternetPagesTable) WithPrefix(prefix string) *FivenetInternetPagesTable {
	return newFivenetInternetPagesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetInternetPagesTable with assigned table suffix
func (a FivenetInternetPagesTable) WithSuffix(suffix string) *FivenetInternetPagesTable {
	return newFivenetInternetPagesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetInternetPagesTable(schemaName, tableName, alias string) *FivenetInternetPagesTable {
	return &FivenetInternetPagesTable{
		fivenetInternetPagesTable: newFivenetInternetPagesTableImpl(schemaName, tableName, alias),
		NEW:                       newFivenetInternetPagesTableImpl("", "new", ""),
	}
}

func newFivenetInternetPagesTableImpl(schemaName, tableName, alias string) fivenetInternetPagesTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		CreatedAtColumn   = mysql.TimestampColumn("created_at")
		UpdatedAtColumn   = mysql.TimestampColumn("updated_at")
		DeletedAtColumn   = mysql.TimestampColumn("deleted_at")
		DomainIDColumn    = mysql.IntegerColumn("domain_id")
		PathColumn        = mysql.StringColumn("path")
		TitleColumn       = mysql.StringColumn("title")
		DescriptionColumn = mysql.StringColumn("description")
		DataColumn        = mysql.StringColumn("data")
		CreatorJobColumn  = mysql.StringColumn("creator_job")
		CreatorIDColumn   = mysql.IntegerColumn("creator_id")
		allColumns        = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, DomainIDColumn, PathColumn, TitleColumn, DescriptionColumn, DataColumn, CreatorJobColumn, CreatorIDColumn}
		mutableColumns    = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, DomainIDColumn, PathColumn, TitleColumn, DescriptionColumn, DataColumn, CreatorJobColumn, CreatorIDColumn}
	)

	return fivenetInternetPagesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		DeletedAt:   DeletedAtColumn,
		DomainID:    DomainIDColumn,
		Path:        PathColumn,
		Title:       TitleColumn,
		Description: DescriptionColumn,
		Data:        DataColumn,
		CreatorJob:  CreatorJobColumn,
		CreatorID:   CreatorIDColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}