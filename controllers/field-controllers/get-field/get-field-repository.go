package getFieldController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	GetFieldRepository(model *models.Field) (*models.Field)
}

type database struct {
	db *gorm.DB
}

func NewGetFieldRepository(db *gorm.DB) *database {
	return &database{db: db}
}

