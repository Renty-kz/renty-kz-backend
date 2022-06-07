package db

import (
	"log"

	registerAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth/register"
	createCityController "github.com/KadirbekSharau/rentykz-backend/controllers/city/create-city"
	createSportTypeController "github.com/KadirbekSharau/rentykz-backend/controllers/sport-type/create-sport-type"
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

func AccountsDataMigrator(db *gorm.DB) (*models.EntityUsers, *models.EntityOrganizations, *models.EntityModerators) {
	registerRepository := registerAuthController.NewRepositoryRegister(db)
	registerService := registerAuthController.NewServiceRegister(registerRepository)
	admin := registerAuthController.InputUserRegister{
		Fullname: "admin1",
		Email: "admin1@gmail.com",
		Password: "admin532",
	}
	organization := registerAuthController.InputOrganizationRegister{
		Fullname: "Organization1",
		Email: "organization1@gmail.com",
		Password: "organization532",
	}
	moderator := registerAuthController.InputModeratorRegister{
		Fullname: "Moderator2",
		Email: "moderator2@gmail.com",
		Password: "organization532",
		OrganizationID: 5,
	}
	newAdmin, errAdmin := registerService.AdminRegisterService(&admin)
	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Printf(errAdmin)
	}
	newOrganization, errOrgaization := registerService.InactiveOrganizationRegisterService(&organization)
	if errOrgaization == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Printf(errAdmin)
	}
	newModerator, errModerator := registerService.ModeratorRegisterService(&moderator)
	if errModerator == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Printf(errAdmin)
	}
	return newAdmin, newOrganization, newModerator;
}

func CityDataMigrator(city string, db *gorm.DB) *models.EntityCities {
	cityRepository := createCityController.NewCreateCityRepository(db)
	cityCreateService := createCityController.NewCreateCityService(cityRepository)

	city1 := createCityController.InputCreateCity{
		Name: city,
	}
	newCity, errCity := cityCreateService.CreateCityService(&city1)
	if errCity == "REGISTER_CONFLICT_409" || errCity == "REGISTER_FAILED_403" {
		log.Printf(errCity)
	}
	
	return newCity
}

func SportTypeDataMigrator(sportType string, db *gorm.DB) *models.EntitySportTypes {
	sportTypeRepository := createSportTypeController.NewCreateSportTypeRepository(db)
	sportTypeCreateService := createSportTypeController.NewCreateSportTypeService(sportTypeRepository)

	addSportType := createSportTypeController.InputCreateSportType{
		Name: sportType,
	}
	newSportType, errSportType := sportTypeCreateService.CreateSportTypeService(&addSportType)
	if errSportType == "REGISTER_CONFLICT_409" || errSportType == "REGISTER_FAILED_403" {
		log.Printf(errSportType)
	}

	return newSportType
}