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
	UserID int
	OrganizationID int
	ModeratorID int
	CourtID int
}