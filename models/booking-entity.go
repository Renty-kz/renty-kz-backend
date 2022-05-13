package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityBookings struct {
	gorm.Model
	Price uint
	StartDate time.Time
	EndDate  time.Time
	UserID uint
	OrganizationID uint
	ModeratorID uint
	CourtID uint
}