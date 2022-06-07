package getFieldController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetFieldByIdRepository(input *InputField) (*models.EntityFields, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetFieldByIdRepository(input *InputField) (*models.EntityFields, string) {
	var field models.EntityFields
	db := r.db.Model(&field)
	errorCode := make(chan string, 1)

	resultField := db.Debug().Select("*").Where("id = ?", input.ID).Find(&field)

	if resultField.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &field, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &field, <-errorCode

}

