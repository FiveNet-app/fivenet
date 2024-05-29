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

var FivenetQualificationsExamQuestions = newFivenetQualificationsExamQuestionsTable("", "fivenet_qualifications_exam_questions", "")

type fivenetQualificationsExamQuestionsTable struct {
	mysql.Table

	// Columns
	ID              mysql.ColumnInteger
	QualificationID mysql.ColumnInteger
	CreatedAt       mysql.ColumnTimestamp
	UpdatedAt       mysql.ColumnTimestamp
	Title           mysql.ColumnString
	Description     mysql.ColumnString
	Data            mysql.ColumnString
	Answer          mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type FivenetQualificationsExamQuestionsTable struct {
	fivenetQualificationsExamQuestionsTable

	NEW fivenetQualificationsExamQuestionsTable
}

// AS creates new FivenetQualificationsExamQuestionsTable with assigned alias
func (a FivenetQualificationsExamQuestionsTable) AS(alias string) *FivenetQualificationsExamQuestionsTable {
	return newFivenetQualificationsExamQuestionsTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FivenetQualificationsExamQuestionsTable with assigned schema name
func (a FivenetQualificationsExamQuestionsTable) FromSchema(schemaName string) *FivenetQualificationsExamQuestionsTable {
	return newFivenetQualificationsExamQuestionsTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FivenetQualificationsExamQuestionsTable with assigned table prefix
func (a FivenetQualificationsExamQuestionsTable) WithPrefix(prefix string) *FivenetQualificationsExamQuestionsTable {
	return newFivenetQualificationsExamQuestionsTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FivenetQualificationsExamQuestionsTable with assigned table suffix
func (a FivenetQualificationsExamQuestionsTable) WithSuffix(suffix string) *FivenetQualificationsExamQuestionsTable {
	return newFivenetQualificationsExamQuestionsTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFivenetQualificationsExamQuestionsTable(schemaName, tableName, alias string) *FivenetQualificationsExamQuestionsTable {
	return &FivenetQualificationsExamQuestionsTable{
		fivenetQualificationsExamQuestionsTable: newFivenetQualificationsExamQuestionsTableImpl(schemaName, tableName, alias),
		NEW:                                     newFivenetQualificationsExamQuestionsTableImpl("", "new", ""),
	}
}

func newFivenetQualificationsExamQuestionsTableImpl(schemaName, tableName, alias string) fivenetQualificationsExamQuestionsTable {
	var (
		IDColumn              = mysql.IntegerColumn("id")
		QualificationIDColumn = mysql.IntegerColumn("qualification_id")
		CreatedAtColumn       = mysql.TimestampColumn("created_at")
		UpdatedAtColumn       = mysql.TimestampColumn("updated_at")
		TitleColumn           = mysql.StringColumn("title")
		DescriptionColumn     = mysql.StringColumn("description")
		DataColumn            = mysql.StringColumn("data")
		AnswerColumn          = mysql.StringColumn("answer")
		allColumns            = mysql.ColumnList{IDColumn, QualificationIDColumn, CreatedAtColumn, UpdatedAtColumn, TitleColumn, DescriptionColumn, DataColumn, AnswerColumn}
		mutableColumns        = mysql.ColumnList{QualificationIDColumn, CreatedAtColumn, UpdatedAtColumn, TitleColumn, DescriptionColumn, DataColumn, AnswerColumn}
	)

	return fivenetQualificationsExamQuestionsTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:              IDColumn,
		QualificationID: QualificationIDColumn,
		CreatedAt:       CreatedAtColumn,
		UpdatedAt:       UpdatedAtColumn,
		Title:           TitleColumn,
		Description:     DescriptionColumn,
		Data:            DataColumn,
		Answer:          AnswerColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}