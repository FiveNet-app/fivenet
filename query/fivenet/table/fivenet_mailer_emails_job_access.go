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

var FivenetMailerEmailsJobAccess = newFivenetMailerEmailsJobAccessTable("", "fivenet_mailer_emails_job_access", "")

type fivenetMailerEmailsJobAccessTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	EmailID      mysql.ColumnInteger
	Job          mysql.ColumnString
	MinimumGrade mysql.ColumnInteger
	Access       mysql.ColumnInteger

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetMailerEmailsJobAccessTable struct {
	fivenetMailerEmailsJobAccessTable

	NEW fivenetMailerEmailsJobAccessTable
}

// AS creates new FivenetMailerEmailsJobAccessTable with assigned alias
func (a FivenetMailerEmailsJobAccessTable) AS(alias string) *FivenetMailerEmailsJobAccessTable {
	return newFivenetMailerEmailsJobAccessTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetMailerEmailsJobAccessTable with assigned schema name
func (a FivenetMailerEmailsJobAccessTable) FromSchema(schemaName string) *FivenetMailerEmailsJobAccessTable {
	return newFivenetMailerEmailsJobAccessTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetMailerEmailsJobAccessTable with assigned table prefix
func (a FivenetMailerEmailsJobAccessTable) WithPrefix(prefix string) *FivenetMailerEmailsJobAccessTable {
	return newFivenetMailerEmailsJobAccessTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetMailerEmailsJobAccessTable with assigned table suffix
func (a FivenetMailerEmailsJobAccessTable) WithSuffix(suffix string) *FivenetMailerEmailsJobAccessTable {
	return newFivenetMailerEmailsJobAccessTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetMailerEmailsJobAccessTable(schemaName, tableName, alias string) *FivenetMailerEmailsJobAccessTable {
	return &FivenetMailerEmailsJobAccessTable{
		fivenetMailerEmailsJobAccessTable: newFivenetMailerEmailsJobAccessTableImpl(schemaName, tableName, alias),
		NEW:                               newFivenetMailerEmailsJobAccessTableImpl("", "new", ""),
	}
}

func newFivenetMailerEmailsJobAccessTableImpl(schemaName, tableName, alias string) fivenetMailerEmailsJobAccessTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		EmailIDColumn      = mysql.IntegerColumn("email_id")
		JobColumn          = mysql.StringColumn("job")
		MinimumGradeColumn = mysql.IntegerColumn("minimum_grade")
		AccessColumn       = mysql.IntegerColumn("access")
		allColumns         = mysql.ColumnList{IDColumn, CreatedAtColumn, EmailIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
		mutableColumns     = mysql.ColumnList{CreatedAtColumn, EmailIDColumn, JobColumn, MinimumGradeColumn, AccessColumn}
	)

	return fivenetMailerEmailsJobAccessTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		CreatedAt:    CreatedAtColumn,
		EmailID:      EmailIDColumn,
		Job:          JobColumn,
		MinimumGrade: MinimumGradeColumn,
		Access:       AccessColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}