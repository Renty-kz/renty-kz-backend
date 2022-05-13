package models

import (
	"time"

	"github.com/KadirbekSharau/rentykz-backend/util"
	"gorm.io/gorm"
)

type EntityUsers struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	IsAdmin    bool   `gorm:"type:bool;default:false"`
	SportTypes []EntitySportType `gorm:"foreignKey:UserID"`
	Cities []EntityCities `gorm:"foreignKey:UserID"`
	Bookings []EntityBookings `gorm:"foreignKey:UserID"`
	Organizations []EntityOrganizations `gorm:"foreignKey:UserID"`
	Fields []EntityFields `gorm:"foreignKey:UserID"`
}

func (entity *EntityUsers) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityUsers) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}