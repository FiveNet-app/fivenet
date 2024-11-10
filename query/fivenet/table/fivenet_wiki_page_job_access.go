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

var FivenetWikiPageJobAccess = newFivenetWikiPageJobAccessTable("", "fivenet_wiki_page_job_access", "")

type fivenetWikiPageJobAccessTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	PageID       mysql.ColumnInteger
	Job          mysql.ColumnString
	MinimumGrade mysql.ColumnInteger
	Access       mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetWikiPageJobAccessTable struct {
	fivenetWikiPageJobAccessTable

	NEW fivenetWikiPageJobAccessTable
}

// AS creates new FivenetWikiPageJobAccessTable with assigned alias
func (a FivenetWikiPageJobAccessTable) AS(alias string) *FivenetWikiPageJobAccessTable {
	return newFivenetWikiPageJobAccessTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetWikiPageJobAccessTable with assigned schema name
func (a FivenetWikiPageJobAccessTable) FromSchema(schemaName string) *FivenetWikiPageJobAccessTable {
	return newFivenetWikiPageJobAccessTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetWikiPageJobAccessTable with assigned table prefix
func (a FivenetWikiPageJobAccessTable) WithPrefix(prefix string) *FivenetWikiPageJobAccessTable {
	return newFivenetWikiPageJobAccessTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetWikiPageJobAccessTable with assigned table suffix
func (a FivenetWikiPageJobAccessTable) WithSuffix(suffix string) *FivenetWikiPageJobAccessTable {
	return newFivenetWikiPageJobAccessTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetWikiPageJobAccessTable(schemaName, tableName, alias string) *FivenetWikiPageJobAccessTable {
	return &FivenetWikiPageJobAccessTable{
		fivenetWikiPageJobAccessTable: newFivenetWikiPageJobAccessTableImpl(schemaName, tableName, alias),
		NEW:                           newFivenetWikiPageJobAccessTableImpl("", "new", ""),
	}
}

func newFivenetWikiPageJobAccessTableImpl(schemaName, tableName, alias string) fivenetWikiPageJobAccessTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		PageIDColumn       = mysql.IntegerColumn("page_id")
		JobColumn          = mysql.StringColumn("job")
		MinimumGradeColumn = mysql.IntegerColumn("minimum_grade")
		AccessColumn       = mysql.IntegerColumn("access")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, PageIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, PageIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
	)

	return fivenetWikiPageJobAccessTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		PageID:       PageIDColumn,
		Job:          JobColumn,
		MinimumGrade: MinimumGradeColumn,
		Access:       AccessColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}