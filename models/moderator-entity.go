package models

import (
	"github.com/KadirbekSharau/rentykz-backend/util"
	"gorm.io/gorm"
)

type EntityModerators struct {
	gorm.Model
	Fullname string `gorm:"type:varchar(255)"`
	PhoneNumber string `gorm:"type:varchar(12)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	Fields []EntityFields `gorm:"foreignKey:ModeratorID"`
	AdminID uint
	OrganizationID uint
}

func (entity *EntityModerators) BeforeCreate(db *gorm.DB) error {
	entity.Password = util.HashPassword(entity.Password)
	return nil
}