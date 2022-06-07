package updateOrganizationController

import (
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	ActivateOrganizationByIdRepository(input *InputActivateOrganization) (*models.EntityOrganizations, string)
	InactivateOrganizationByIdRepository(input *InputActivateOrganization) (*models.EntityOrganizations, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Activate Organization Repository */
func (r *repository) ActivateOrganizationByIdRepository(input *InputActivateOrganization) (*models.EntityOrganizations, string) {
	var organizations models.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)


	result := db.Debug().Where("id = ?", input.ID).Update("active", true).Find(&organizations)


	if result.RowsAffected < 1 {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	}else if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}

/* Inactivate Organization Repository */
func (r *repository) InactivateOrganizationByIdRepository(input *InputActivateOrganization) (*models.EntityOrganizations, string) {
	var organizations models.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)


	result := db.Debug().Where("id = ?", input.ID).Update("active", false).Find(&organizations)


	if result.RowsAffected < 1 {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	}else if result.Error != nil {
		errorCode <- "RESULTS_ORGANIZATION_NOT_FOUND_404"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}