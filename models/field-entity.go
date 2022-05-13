package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityCities struct {
	gorm.Model
	Name string `json:"name"`
	UserID int
	SportTypes []EntitySportType `gorm:"foreignKey:CityID"`
}

type EntitySportType struct {
	gorm.Model
	Name string `json:"name"`
	Fields []EntityFields `gorm:"foreignKey:SportTypeID"`
	OrganizationID int
	UserID int
	CityID int
}

type EntityFields struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Contacts string `gorm:"type:varchar(255)"`
	Description string `gorm:"type:varchar(500)"`
	Price string `gorm:"type:int"`
	Capacity string `gorm:"type:int"`
	FieldRates []EntityFieldRates `gorm:"foreignKey:FieldID"`
	Bookings []EntityBookings `gorm:"foreignKey:FieldID"`
	OrganizationID int
	UserID int
	CityID int
	SportTypeID int
	ModeratorID int
}

type EntityFieldRates struct {
	gorm.Model
	Price uint
	StartDate time.Time
	EndDate time.Time
	Name string `json:"name"`
	Email string `json:"email"`
	FieldID int
}