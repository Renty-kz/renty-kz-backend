package createFieldController

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateFieldRepository(input *model.EntityFields) (*model.EntityFields, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateFieldRepository(input *model.EntityFields) (*model.EntityFields, string) {

	var fields model.EntityFields
	db := r.db.Model(&fields)
	errorCode := make(chan string, 1)

	checkFieldExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&fields)

	if checkFieldExist.RowsAffected > 0 {
		errorCode <- "CREATE_STUDENT_CONFLICT_409"
		return &fields, <-errorCode
	}

	fields.Name = input.Name
	fields.Address = input.Address
	fields.Price = input.Price
	fields.Capacity = input.Capacity
	fields.Description = input.Description
	fields.Contacts = input.Contacts
	fields.ImageLinks = input.ImageLinks
	fields.AdminID = 1
	fields.CityID = input.CityID
	fields.SportTypeID = input.SportTypeID
	fields.OrganizationID = input.OrganizationID
	fields.ModeratorID = input.ModeratorID

	addNewField := db.Debug().Create(&fields)
	db.Commit()

	if addNewField.Error != nil {
		errorCode <- "CREATE_FIELD_FAILED_403"
		return &fields, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &fields, <-errorCode
}