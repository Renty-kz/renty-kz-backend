package getOrganizationsController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	GetActiveOrganizationsRepository() (*[]models.EntityOrganizations, string)
	GetInactiveOrganizationsRepository() (*[]models.EntityOrganizations, string)
	GetAllOrganizationsRepository() (*[]models.EntityOrganizations, string)

}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Get All Organizations Repository */
func (r *repository) GetAllOrganizationsRepository() (*[]models.EntityOrganizations, string) {
	var organizations []models.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Find(&organizations)

	if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}

/* Get All Active Organizations Repository */
func (r *repository) GetActiveOrganizationsRepository() (*[]models.EntityOrganizations, string) {
	var organizations []models.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Where("active = ?", true).Find(&organizations)

	if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}

/* Get All Inactive Repository (Not Implemented) */
func (r *repository) GetInactiveOrganizationsRepository() (*[]models.EntityOrganizations, string) {
	var organizations []models.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Where("active = ?", false).Find(&organizations)

	if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}