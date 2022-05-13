package getFieldsController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetFieldsRepository() (*[]models.EntityFields, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetFieldsRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetFieldsRepository() (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultFields := db.Debug().Select("*").Find(&fields)

	if resultFields.Error != nil {
		errorCode <- "RESULTS_STUDENT_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}

