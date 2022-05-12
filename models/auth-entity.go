package models

import "time"

type EntityUsers struct {
	ID string `gorm:"primaryKey;"`
	Fullname string `gorm:"type:varchar(255)"`
	Email string `gorm:"type:varchar(255);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Active    bool   `gorm:"type:bool;default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time
}