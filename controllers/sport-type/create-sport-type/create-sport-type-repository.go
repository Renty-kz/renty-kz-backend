package createSportTypeController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateSportTypeRepository(input *models.EntitySportTypes) (*models.EntitySportTypes, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateSportTypeRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateSportTypeRepository(input *models.EntitySportTypes) (*models.EntitySportTypes, string) {
	var sportType models.EntitySportTypes
	db := r.db.Model(&sportType)
	errorCode := make(chan string, 1)

	checkSportTypeExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&sportType)

	if checkSportTypeExist.RowsAffected > 0 {
		errorCode <- "CREATE_SPORT_TYPE_CONFLICT_409"
		return &sportType, <-errorCode
	}

	sportType.Name = input.Name
	sportType.AdminID = 1

	addNewSportType := db.Debug().Create(&sportType)
	db.Commit()

	if addNewSportType.Error != nil {
		errorCode <- "CREATE_SPORT_TYPE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &sportType, <-errorCode
}