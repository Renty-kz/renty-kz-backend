package models

import "gorm.io/gorm"

type EntityModerators struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Fields []EntityFields `gorm:"foreignKey:ModeratorID"`
	UserID int
	OrganizationID int
}