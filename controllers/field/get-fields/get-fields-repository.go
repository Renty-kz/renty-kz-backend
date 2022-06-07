package getFieldsController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetFieldsRepository() (*[]models.EntityFields, string)
	GetFieldsByCityIdRepository(input *InputFieldByCityId) (*[]models.EntityFields, string)
	GetFieldsBySportTypeIdRepository(input *InputFieldsBySportTypeId) (*[]models.EntityFields, string)
	GetFieldsByOrganizationIdRepository(input *InputFieldsByOrganizationId) (*[]models.EntityFields, string)
	GetFieldsByModeratorIdRepository(input *InputFieldsByModeratorId) (*[]models.EntityFields, string)
}

type repository struct {
	db *gorm.DB
}

func NewGetFieldsRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Get All Fields Repository */
func (r *repository) GetFieldsRepository() (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultFields := db.Debug().Select("*").Find(&fields)

	if resultFields.Error != nil {
		errorCode <- "RESULTS_FIELD_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}

/* Get Fields By City Id Repository (Not Implemented) */
func (r *repository) GetFieldsByCityIdRepository(input *InputFieldByCityId) (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultField := db.Debug().Select("*").Where("city_id = ?", input.CityID).Find(&fields)

	if resultField.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}

/* Get Fields By Sport Type Id Repository (Not Implemented) */
func (r *repository) GetFieldsBySportTypeIdRepository(input *InputFieldsBySportTypeId) (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultFields := db.Debug().Select("*").Where("sport_type_id = ?", input.SportTypeID).Find(&fields)

	if resultFields.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}

/* Get Fields By Organization Id Repository (Not Implemented) */
func (r *repository) GetFieldsByOrganizationIdRepository(input *InputFieldsByOrganizationId) (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultFields := db.Debug().Select("*").Where("organization_id = ?", input.OrganizationID).Find(&fields)

	if resultFields.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}


/* Get Fields By Moderator Id Repository (Not Implemented) */
func (r *repository) GetFieldsByModeratorIdRepository(input *InputFieldsByModeratorId) (*[]models.EntityFields, string) {
	var fields []models.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	resultFields := db.Debug().Select("*").Where("moderator_id = ?", input.ModeratorID).Find(&fields)

	if resultFields.RowsAffected < 1 {
		errorCode <- "FIELD_NOT_FOUND_404"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}
