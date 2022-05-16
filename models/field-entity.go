package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityCities struct {
	gorm.Model
	Name string
	Fields []EntityFields `gorm:"foreignKey:CityID"`
	AdminID uint
}

type EntitySportTypes struct {
	gorm.Model
	Name string
	Fields []EntityFields `gorm:"foreignKey:SportTypeID"`
	OrganizationID uint
	AdminID uint
}

type EntityFields struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Contacts []string `gorm:"type:text"`
	Description string `gorm:"type:varchar(500)"`
	ImageLinks []string `gorm:"type:text"`
	Price string `gorm:"type:int"`
	Capacity string `gorm:"type:int"`
	FieldRates []EntityFieldRates `gorm:"foreignKey:FieldID"`
	Bookings []EntityBookings `gorm:"foreignKey:FieldID"`
	OrganizationID uint
	AdminID uint
	CityID uint
	SportTypeID uint
	ModeratorID uint
}

type EntityFieldRates struct {
	gorm.Model
	Price uint
	StartDate time.Time
	EndDate time.Time
	Name string
	Email string
	FieldID uint
}