package models

import (
	"github.com/KadirbekSharau/rentykz-backend/util"
	"gorm.io/gorm"
)

type EntityOrganizations struct {
	gorm.Model
	Fullname string
	Email string
	Password string `gorm:"type:varchar(255);not null"`
	Bookings []EntityBookings `gorm:"foreignKey:OrganizationID"`
	Moderators []EntityModerators `gorm:"foreignKey:OrganizationID"`
	Fields []EntityFields `gorm:"foreignKey:OrganizationID"`
	SportTypes []EntitySportTypes `gorm:"foreignKey:OrganizationID"`
	AdminID uint
}

func (entity *EntityOrganizations) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}