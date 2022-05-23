package models

import (
	"github.com/KadirbekSharau/rentykz-backend/util"
	"gorm.io/gorm"
)

type EntityOrganizations struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(12)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	Bookings []EntityBookings `gorm:"foreignKey:OrganizationID"`
	Moderators []EntityModerators `gorm:"foreignKey:OrganizationID"`
	Fields []EntityFields `gorm:"foreignKey:OrganizationID"`
	AdminID uint
}

func (entity *EntityOrganizations) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}