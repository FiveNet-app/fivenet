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

var FivenetJobsQualificationsResults = newFivenetJobsQualificationsResultsTable("", "fivenet_jobs_qualifications_results", "")

type fivenetJobsQualificationsResultsTable struct {
	mysql.Table

	// Columns
	ID              mysql.ColumnInteger
	CreatedAt       mysql.ColumnTimestamp
	DeletedAt       mysql.ColumnTimestamp
	QualificationID mysql.ColumnInteger
	UserID          mysql.ColumnInteger
	Status          mysql.ColumnInteger
	Score           mysql.ColumnInteger
	Summary         mysql.ColumnString
	CreatorID       mysql.ColumnInteger
	CreatorJob      mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetJobsQualificationsResultsTable struct {
	fivenetJobsQualificationsResultsTable

	NEW fivenetJobsQualificationsResultsTable
}

// AS creates new FivenetJobsQualificationsResultsTable with assigned alias
func (a FivenetJobsQualificationsResultsTable) AS(alias string) *FivenetJobsQualificationsResultsTable {
	return newFivenetJobsQualificationsResultsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetJobsQualificationsResultsTable with assigned schema name
func (a FivenetJobsQualificationsResultsTable) FromSchema(schemaName string) *FivenetJobsQualificationsResultsTable {
	return newFivenetJobsQualificationsResultsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetJobsQualificationsResultsTable with assigned table prefix
func (a FivenetJobsQualificationsResultsTable) WithPrefix(prefix string) *FivenetJobsQualificationsResultsTable {
	return newFivenetJobsQualificationsResultsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetJobsQualificationsResultsTable with assigned table suffix
func (a FivenetJobsQualificationsResultsTable) WithSuffix(suffix string) *FivenetJobsQualificationsResultsTable {
	return newFivenetJobsQualificationsResultsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetJobsQualificationsResultsTable(schemaName, tableName, alias string) *FivenetJobsQualificationsResultsTable {
	return &FivenetJobsQualificationsResultsTable{
		fivenetJobsQualificationsResultsTable: newFivenetJobsQualificationsResultsTableImpl(schemaName, tableName, alias),
		NEW:                                   newFivenetJobsQualificationsResultsTableImpl("", "new", ""),
	}
}

func newFivenetJobsQualificationsResultsTableImpl(schemaName, tableName, alias string) fivenetJobsQualificationsResultsTable {
	var (
		IDColumn              = mysql.IntegerColumn("id")
		CreatedAtColumn       = mysql.TimestampColumn("created_at")
		DeletedAtColumn       = mysql.TimestampColumn("deleted_at")
		QualificationIDColumn = mysql.IntegerColumn("qualification_id")
		UserIDColumn          = mysql.IntegerColumn("user_id")
		StatusColumn          = mysql.IntegerColumn("status")
		ScoreColumn           = mysql.IntegerColumn("score")
		SummaryColumn         = mysql.StringColumn("summary")
		CreatorIDColumn       = mysql.IntegerColumn("creator_id")
		CreatorJobColumn      = mysql.StringColumn("creator_job")
		allColumns            = mysql.ColumnList{IDColumn, CreatedAtColumn, DeletedAtColumn, QualificationIDColumn, UserIDColumn, StatusColumn, ScoreColumn, SummaryColumn, CreatorIDColumn, CreatorJobColumn}
		mutableColumns        = mysql.ColumnList{CreatedAtColumn, DeletedAtColumn, QualificationIDColumn, UserIDColumn, StatusColumn, ScoreColumn, SummaryColumn, CreatorIDColumn, CreatorJobColumn}
	)

	return fivenetJobsQualificationsResultsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		CreatedAt:       CreatedAtColumn,
		DeletedAt:       DeletedAtColumn,
		QualificationID: QualificationIDColumn,
		UserID:          UserIDColumn,
		Status:          StatusColumn,
		Score:           ScoreColumn,
		Summary:         SummaryColumn,
		CreatorID:       CreatorIDColumn,
		CreatorJob:      CreatorJobColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
