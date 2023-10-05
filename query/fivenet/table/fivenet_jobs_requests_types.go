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

var FivenetJobsRequestsTypes = newFivenetJobsRequestsTypesTable("", "fivenet_jobs_requests_types", "")

type fivenetJobsRequestsTypesTable struct {
	mysql.Table

	// Columns
	ID          mysql.ColumnInteger
	CreatedAt   mysql.ColumnTimestamp
	UpdatedAt   mysql.ColumnTimestamp
	DeletedAt   mysql.ColumnTimestamp
	Job         mysql.ColumnString
	Name        mysql.ColumnString
	Description mysql.ColumnString
	Weight      mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetJobsRequestsTypesTable struct {
	fivenetJobsRequestsTypesTable

	NEW fivenetJobsRequestsTypesTable
}

// AS creates new FivenetJobsRequestsTypesTable with assigned alias
func (a FivenetJobsRequestsTypesTable) AS(alias string) *FivenetJobsRequestsTypesTable {
	return newFivenetJobsRequestsTypesTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetJobsRequestsTypesTable with assigned schema name
func (a FivenetJobsRequestsTypesTable) FromSchema(schemaName string) *FivenetJobsRequestsTypesTable {
	return newFivenetJobsRequestsTypesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetJobsRequestsTypesTable with assigned table prefix
func (a FivenetJobsRequestsTypesTable) WithPrefix(prefix string) *FivenetJobsRequestsTypesTable {
	return newFivenetJobsRequestsTypesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetJobsRequestsTypesTable with assigned table suffix
func (a FivenetJobsRequestsTypesTable) WithSuffix(suffix string) *FivenetJobsRequestsTypesTable {
	return newFivenetJobsRequestsTypesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetJobsRequestsTypesTable(schemaName, tableName, alias string) *FivenetJobsRequestsTypesTable {
	return &FivenetJobsRequestsTypesTable{
		fivenetJobsRequestsTypesTable: newFivenetJobsRequestsTypesTableImpl(schemaName, tableName, alias),
		NEW:                           newFivenetJobsRequestsTypesTableImpl("", "new", ""),
	}
}

func newFivenetJobsRequestsTypesTableImpl(schemaName, tableName, alias string) fivenetJobsRequestsTypesTable {
	var (
		IDColumn          = mysql.IntegerColumn("id")
		CreatedAtColumn   = mysql.TimestampColumn("created_at")
		UpdatedAtColumn   = mysql.TimestampColumn("updated_at")
		DeletedAtColumn   = mysql.TimestampColumn("deleted_at")
		JobColumn         = mysql.StringColumn("job")
		NameColumn        = mysql.StringColumn("name")
		DescriptionColumn = mysql.StringColumn("description")
		WeightColumn      = mysql.IntegerColumn("weight")
		allColumns        = mysql.ColumnList{IDColumn, CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, JobColumn, NameColumn, DescriptionColumn, WeightColumn}
		mutableColumns    = mysql.ColumnList{CreatedAtColumn, UpdatedAtColumn, DeletedAtColumn, JobColumn, NameColumn, DescriptionColumn, WeightColumn}
	)

	return fivenetJobsRequestsTypesTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:          IDColumn,
		CreatedAt:   CreatedAtColumn,
		UpdatedAt:   UpdatedAtColumn,
		DeletedAt:   DeletedAtColumn,
		Job:         JobColumn,
		Name:        NameColumn,
		Description: DescriptionColumn,
		Weight:      WeightColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
