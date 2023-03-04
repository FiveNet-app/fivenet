// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/galexrt/arpanet/pkg/permify/models"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID             int32               `gorm:"column:id;type:int(11);not null;uniqueIndex:id,priority:1" json:"id"`
	Identifier     string              `gorm:"column:identifier;type:varchar(64);primaryKey" json:"identifier"`
	Job            string              `gorm:"column:job;type:varchar(20);index:idx_users_job,priority:1;default:unemployed" json:"job"`
	JobGrade       int                 `gorm:"column:job_grade;type:int(11);default:1" json:"job_grade"`
	Firstname      string              `gorm:"column:firstname;type:varchar(50);index:idx_users_firstname_lastname,priority:1" json:"firstname"`
	Lastname       string              `gorm:"column:lastname;type:varchar(50);index:idx_users_firstname_lastname,priority:2" json:"lastname"`
	Dateofbirth    string              `gorm:"column:dateofbirth;type:varchar(25)" json:"dateofbirth"`
	Sex            Sex                 `gorm:"column:sex;type:varchar(10)" json:"sex"`
	Height         string              `gorm:"column:height;type:varchar(5)" json:"height"`
	Jail           int32               `gorm:"column:jail;type:int(11);not null" json:"jail"`
	PhoneNumber    string              `gorm:"column:phone_number;type:varchar(20)" json:"phone_number"`
	Accounts       MoneyAccounts       `gorm:"serializer:json" json:"-"`
	Disabled       bool                `gorm:"column:disabled;type:tinyint(1)" json:"disabled"`
	Visum          int32               `gorm:"column:visum;type:int(11)" json:"visum"`
	Playtime       int32               `gorm:"column:playtime;type:int(11)" json:"playtime"`
	CreatedAt      time.Time           `gorm:"column:created_at;type:timestamp;default:current_timestamp()" json:"created_at"`
	UpdatedAt      time.Time           `gorm:"column:last_seen;type:timestamp" json:"updated_at"`
	UserProps      UserProps           `gorm:"foreignkey:Identifier" json:"user_props"`
	UserLicenses   []UserLicense       `gorm:"foreignkey:Owner" json:"user_licenses"`
	Documents      []Document          `gorm:"foreignkey:Creator" json:"documents"`
	UserActivity   []UserActivity      `gorm:"foreignkey:Identifier" json:"user_activity"`
	CausedActivity []UserActivity      `gorm:"foreignkey:CauseIdentifier" json:"caused_activity"`
	Roles          []models.Role       `gorm:"many2many:arpanet_user_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"roles"`
	Permissions    []models.Permission `gorm:"many2many:arpanet_user_permissions;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"permissions"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
