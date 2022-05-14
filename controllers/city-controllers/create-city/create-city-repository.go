package createCityController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateCityRepository(input *models.EntityCities) (*models.EntityCities, string)
}

type repository struct {
	db *gorm.DB
}

func NewCreateFieldRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateCityRepository(input *models.EntityCities) (*models.EntityCities, string) {
	var city models.EntityCities
	db := r.db.Model(&city)
	errorCode := make(chan string, 1)

	city.Name = input.Name

	addNewCity := db.Debug().Create(&city)
	db.Commit()

	if addNewCity.Error != nil {
		errorCode <- "CREATE_CITY_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &city, <-errorCode
}