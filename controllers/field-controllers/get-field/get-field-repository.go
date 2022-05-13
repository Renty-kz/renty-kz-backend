package getFieldController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetFieldRepository(model *models.EntityFields) (*models.EntityFields)
}

type database struct {
	db *gorm.DB
}

func NewGetFieldRepository(db *gorm.DB) *database {
	return &database{db: db}
}

