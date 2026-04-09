package models

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model

	FirstName  string `gorm:"column:first_name;type:varchar(100);not null"`
	MiddleName string `gorm:"column:middle_name;type:varchar(100)"`
	LastName   string `gorm:"column:last_name;type:varchar(100);not null"`

	Email string `gorm:"column:email;type:varchar(150);uniqueIndex;not null"`

	GuardianName     string  `gorm:"column:guardian_name;type:varchar(150)"`
	GuardianRelation string  `gorm:"column:guardian_relation;type:varchar(50)"`
	GuardianContact  string  `gorm:"column:guardian_contact;type:varchar(20)"`
	Class            string  `gorm:"column:class;type:varchar(20)"`
	Address          Address `gorm:"column:address;type:jsonb"`
	IsActive         bool    `gorm:"column:is_active;type:bool;default:true"`
}

type RankHistory struct {
	gorm.Model
	StudentID uint `gorm:"column:student_id;not null;index"`

	Term string `gorm:"column:term;type:varchar(20);not null"`
	Year int    `gorm:"column:year;not null"`

	Rank          int    `gorm:"column:rank;not null"`
	MarksAttained int    `gorm:"column:marks_attained"`
	Grade         string `gorm:"column:grade;type:varchar(5)"`
}

type Address struct {
	Line1 string
	Line2 string
	City  string
	State string
	Pin   string
}

type Admin struct {
	gorm.Model
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"-" gorm:"column:password_hash;not null;size:128"`
}
