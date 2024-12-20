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

var Users = newUsersTable("", "users", "")

type usersTable struct {
	mysql.Table

	// Columns
	ID           mysql.ColumnInteger
	Identifier   mysql.ColumnString
	Group        mysql.ColumnString
	Skin         mysql.ColumnString
	Job          mysql.ColumnString
	JobGrade     mysql.ColumnInteger
	Loadout      mysql.ColumnString
	Position     mysql.ColumnString
	Firstname    mysql.ColumnString
	Lastname     mysql.ColumnString
	Dateofbirth  mysql.ColumnString
	Sex          mysql.ColumnString
	Height       mysql.ColumnString
	IsDead       mysql.ColumnBool
	LastProperty mysql.ColumnString
	Jail         mysql.ColumnInteger
	Inventory    mysql.ColumnString
	PhoneNumber  mysql.ColumnString
	Accounts     mysql.ColumnString
	Tattoos      mysql.ColumnString
	Disabled     mysql.ColumnBool
	Visum        mysql.ColumnInteger
	Playtime     mysql.ColumnInteger
	LevelData    mysql.ColumnString
	OnDuty       mysql.ColumnInteger
	Health       mysql.ColumnInteger
	Armor        mysql.ColumnInteger
	CreatedAt    mysql.ColumnTimestamp
	LastSeen     mysql.ColumnTimestamp
	Metadata     mysql.ColumnString
	Crew         mysql.ColumnString
	CrewLeader   mysql.ColumnBool

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type UsersTable struct {
	usersTable

	NEW usersTable
}

// AS creates new UsersTable with assigned alias
func (a UsersTable) AS(alias string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersTable with assigned schema name
func (a UsersTable) FromSchema(schemaName string) *UsersTable {
	return newUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersTable with assigned table prefix
func (a UsersTable) WithPrefix(prefix string) *UsersTable {
	return newUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersTable with assigned table suffix
func (a UsersTable) WithSuffix(suffix string) *UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersTable(schemaName, tableName, alias string) *UsersTable {
	return &UsersTable{
		usersTable: newUsersTableImpl(schemaName, tableName, alias),
		NEW:        newUsersTableImpl("", "new", ""),
	}
}

func newUsersTableImpl(schemaName, tableName, alias string) usersTable {
	var (
		IDColumn           = mysql.IntegerColumn("id")
		IdentifierColumn   = mysql.StringColumn("identifier")
		GroupColumn        = mysql.StringColumn("group")
		SkinColumn         = mysql.StringColumn("skin")
		JobColumn          = mysql.StringColumn("job")
		JobGradeColumn     = mysql.IntegerColumn("job_grade")
		LoadoutColumn      = mysql.StringColumn("loadout")
		PositionColumn     = mysql.StringColumn("position")
		FirstnameColumn    = mysql.StringColumn("firstname")
		LastnameColumn     = mysql.StringColumn("lastname")
		DateofbirthColumn  = mysql.StringColumn("dateofbirth")
		SexColumn          = mysql.StringColumn("sex")
		HeightColumn       = mysql.StringColumn("height")
		IsDeadColumn       = mysql.BoolColumn("is_dead")
		LastPropertyColumn = mysql.StringColumn("last_property")
		JailColumn         = mysql.IntegerColumn("jail")
		InventoryColumn    = mysql.StringColumn("inventory")
		PhoneNumberColumn  = mysql.StringColumn("phone_number")
		AccountsColumn     = mysql.StringColumn("accounts")
		TattoosColumn      = mysql.StringColumn("tattoos")
		DisabledColumn     = mysql.BoolColumn("disabled")
		VisumColumn        = mysql.IntegerColumn("visum")
		PlaytimeColumn     = mysql.IntegerColumn("playtime")
		LevelDataColumn    = mysql.StringColumn("levelData")
		OnDutyColumn       = mysql.IntegerColumn("onDuty")
		HealthColumn       = mysql.IntegerColumn("health")
		ArmorColumn        = mysql.IntegerColumn("armor")
		CreatedAtColumn    = mysql.TimestampColumn("created_at")
		LastSeenColumn     = mysql.TimestampColumn("last_seen")
		MetadataColumn     = mysql.StringColumn("metadata")
		CrewColumn         = mysql.StringColumn("crew")
		CrewLeaderColumn   = mysql.BoolColumn("crewLeader")
		allColumns         = mysql.ColumnList{IDColumn, IdentifierColumn, GroupColumn, SkinColumn, JobColumn, JobGradeColumn, LoadoutColumn, PositionColumn, FirstnameColumn, LastnameColumn, DateofbirthColumn, SexColumn, HeightColumn, IsDeadColumn, LastPropertyColumn, JailColumn, InventoryColumn, PhoneNumberColumn, AccountsColumn, TattoosColumn, DisabledColumn, VisumColumn, PlaytimeColumn, LevelDataColumn, OnDutyColumn, HealthColumn, ArmorColumn, CreatedAtColumn, LastSeenColumn, MetadataColumn, CrewColumn, CrewLeaderColumn}
		mutableColumns     = mysql.ColumnList{IDColumn, GroupColumn, SkinColumn, JobColumn, JobGradeColumn, LoadoutColumn, PositionColumn, FirstnameColumn, LastnameColumn, DateofbirthColumn, SexColumn, HeightColumn, IsDeadColumn, LastPropertyColumn, JailColumn, InventoryColumn, PhoneNumberColumn, AccountsColumn, TattoosColumn, DisabledColumn, VisumColumn, PlaytimeColumn, LevelDataColumn, OnDutyColumn, HealthColumn, ArmorColumn, CreatedAtColumn, LastSeenColumn, MetadataColumn, CrewColumn, CrewLeaderColumn}
	)

	return usersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:           IDColumn,
		Identifier:   IdentifierColumn,
		Group:        GroupColumn,
		Skin:         SkinColumn,
		Job:          JobColumn,
		JobGrade:     JobGradeColumn,
		Loadout:      LoadoutColumn,
		Position:     PositionColumn,
		Firstname:    FirstnameColumn,
		Lastname:     LastnameColumn,
		Dateofbirth:  DateofbirthColumn,
		Sex:          SexColumn,
		Height:       HeightColumn,
		IsDead:       IsDeadColumn,
		LastProperty: LastPropertyColumn,
		Jail:         JailColumn,
		Inventory:    InventoryColumn,
		PhoneNumber:  PhoneNumberColumn,
		Accounts:     AccountsColumn,
		Tattoos:      TattoosColumn,
		Disabled:     DisabledColumn,
		Visum:        VisumColumn,
		Playtime:     PlaytimeColumn,
		LevelData:    LevelDataColumn,
		OnDuty:       OnDutyColumn,
		Health:       HealthColumn,
		Armor:        ArmorColumn,
		CreatedAt:    CreatedAtColumn,
		LastSeen:     LastSeenColumn,
		Metadata:     MetadataColumn,
		Crew:         CrewColumn,
		CrewLeader:   CrewLeaderColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
