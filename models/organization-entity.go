package models

import "gorm.io/gorm"

type EntityOrganizations struct {
	gorm.Model
	Name string
	Email string
	Password string `gorm:"type:varchar(255);not null"`
	Bookings []EntityBookings `gorm:"foreignKey:OrganizationID"`
	Moderators []EntityModerators `gorm:"foreignKey:OrganizationID"`
	Fields []EntityFields `gorm:"foreignKey:OrganizationID"`
	SportTypes []EntitySportType `gorm:"foreignKey:OrganizationID"`
	UserID uint
}