package createAdminController

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	CreateAdminRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateAdminRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var admins model.EntityUsers
	db := r.db.Model(&admins)
	errorCode := make(chan string, 1)

	//checkFieldExist := db.Debug().Select("*").Where("id = ?", input.ID).Find(&fields)

	// if checkFieldExist.RowsAffected > 0 {
	// 	errorCode <- "CREATE_STUDENT_CONFLICT_409"
	// 	return &fields, <-errorCode
	// }
	checkAdminAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&admins)

	if checkAdminAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &admins, <-errorCode
	}

	admins.Fullname = input.Fullname
	admins.Email = input.Email
	admins.Password = input.Password
	admins.Active = input.Active
	admins.IsAdmin = input.IsAdmin

	addNewAdmin := db.Debug().Create(&admins)
	db.Commit()

	if addNewAdmin.Error != nil {
		errorCode <- "CREATE_STUDENT_FAILED_403"
		return &admins, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &admins, <-errorCode
}