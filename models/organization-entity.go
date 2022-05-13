package models

import "gorm.io/gorm"

type EntityOrganizations struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Field EntityFields `json:"field"`
	Bookings []EntityBookings `gorm:"foreignKey:OrganizationID"`
	Moderators []EntityModerators `gorm:"foreignKey:OrganizationID"`
	Fields []EntityFields `gorm:"foreignKey:OrganizationID"`
	SportTypes []EntitySportType `gorm:"foreignKey:OrganizationID"`
	UserID int
}