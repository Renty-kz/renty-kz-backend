package db

import (
	"log"

	registerAuthController "github.com/KadirbekSharau/rentykz-backend/controllers/auth-controllers/register"
	"github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

func AdminDataMigrator(db *gorm.DB) *models.EntityUsers {
	registerAdminRepository := registerAuthController.NewRepositoryRegister(db)
	registerAdminService := registerAuthController.NewServiceRegister(registerAdminRepository)
	admin := registerAuthController.InputRegister{
		Fullname: "admin1",
		Email: "admin1@gmail.com",
		Password: "admin532",
	}
	newAdmin, errAdmin := registerAdminService.AdminRegisterService(&admin)
	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
		log.Fatal(errAdmin)
	}
	return newAdmin;
}

func CitiesDataMigrator(db *gorm.DB) {
	
}