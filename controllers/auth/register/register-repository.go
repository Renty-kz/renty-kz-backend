package registerAuthController

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	ActiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	InactiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	AdminRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	InactiveOrganizationRegisterRepository(input *model.EntityOrganizations) (*model.EntityOrganizations, string)
	ModeratorRegisterRepository(input *model.EntityModerators) (*model.EntityModerators, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Active User Registration Repository */
func (r *repository) ActiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Fullname = input.Fullname
	users.Email = input.Email
	users.Password = input.Password
	users.Active = true

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Inactive User Registration Repository */
func (r *repository) InactiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Fullname = input.Fullname
	users.Email = input.Email
	users.Password = input.Password

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Admin Registration Repository */
func (r *repository) AdminRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Fullname = input.Fullname
	users.Email = input.Email
	users.Password = input.Password
	users.Active = true
	users.IsAdmin = true

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Inactive Organization Registration Repository */
func (r *repository) InactiveOrganizationRegisterRepository(input *model.EntityOrganizations) (*model.EntityOrganizations, string) {
	var organization model.EntityOrganizations
	//var admin model.EntityUsers
	db := r.db.Model(&organization)
	errorCode := make(chan string, 1)

	checkOrganizationAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&organization)

	if checkOrganizationAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &organization, <-errorCode
	}

	//adminID, _ := r.db.Debug().Select("entity_users").Where("is_admin = true").First(&admin, 1).Get("id")
	organization.Fullname = input.Fullname
	organization.Email = input.Email
	organization.Password = input.Password
	organization.Active = true
	organization.AdminID = 1

	addNewOrganization := db.Debug().Create(&organization)
	db.Commit()

	if addNewOrganization.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
	} else {
		errorCode <- "nil"
	}
	
	return &organization, <-errorCode
}


/* Moderator Registration Repository */
func (r *repository) ModeratorRegisterRepository(input *model.EntityModerators) (*model.EntityModerators, string) {
	var moderators model.EntityModerators
	db := r.db.Model(&moderators)
	errorCode := make(chan string, 1)

	checkModeratorAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&moderators)

	if checkModeratorAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &moderators, <-errorCode
	}

	moderators.Fullname = input.Fullname
	moderators.Email = input.Email
	moderators.Password = input.Password
	moderators.AdminID = 1
	moderators.OrganizationID = input.OrganizationID

	addNewModerator := db.Debug().Create(&moderators)
	db.Commit()

	if addNewModerator.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &moderators, <-errorCode
}