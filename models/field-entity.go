package models

import (
	"time"

	"github.com/lib/pq"
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
	AdminID uint
}

type EntityFields struct {
	gorm.Model
	Name string `gorm:"type:varchar(255)"`
	Address string `gorm:"type:varchar(255)"`
	Contacts pq.StringArray `gorm:"type:text[];not null"`
	Description string `gorm:"type:varchar(500)"`
	ImageLinks pq.StringArray `gorm:"type:text[];not null"`
	Price uint
	Capacity uint
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