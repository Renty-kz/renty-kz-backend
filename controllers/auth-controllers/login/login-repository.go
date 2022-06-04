package loginAuthController

import (
	util "github.com/KadirbekSharau/rentykz-backend/util"
	model "github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	UserLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string)
	OrganizationLoginRepository(input *model.EntityOrganizations) (*model.EntityOrganizations, string)
	ModeratorLoginRepository(input *model.EntityModerators) (*model.EntityModerators, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryLogin(db *gorm.DB) *repository {
	return &repository{db: db}
}

//func CheckEmailExists() {}

/* User Login Repository Service */
func (r *repository) UserLoginRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	users.Email = input.Email
	users.Password = input.Password

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &users, <-errorCode
	}

	if !users.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &users, <-errorCode
	}

	comparePassword := util.ComparePassword(users.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}

/* Organization Login Repository Service */
func (r *repository) OrganizationLoginRepository(input *model.EntityOrganizations) (*model.EntityOrganizations, string) {

	var organizations model.EntityOrganizations
	db := r.db.Model(&organizations)
	errorCode := make(chan string, 1)

	organizations.Email = input.Email
	organizations.Password = input.Password

	checkOrganizationAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&organizations)

	if checkOrganizationAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &organizations, <-errorCode
	}

	if !organizations.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &organizations, <-errorCode
	}

	comparePassword := util.ComparePassword(organizations.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &organizations, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &organizations, <-errorCode
}

/* Moderator Login Repository Service */
func (r *repository) ModeratorLoginRepository(input *model.EntityModerators) (*model.EntityModerators, string) {

	var moderators model.EntityModerators
	db := r.db.Model(&moderators)
	errorCode := make(chan string, 1)

	moderators.Email = input.Email
	moderators.Password = input.Password

	checkModeratorAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&moderators)

	if checkModeratorAccount.RowsAffected < 1 {
		errorCode <- "LOGIN_NOT_FOUND_404"
		return &moderators, <-errorCode
	}

	if !moderators.Active {
		errorCode <- "LOGIN_NOT_ACTIVE_403"
		return &moderators, <-errorCode
	}

	comparePassword := util.ComparePassword(moderators.Password, input.Password)

	if comparePassword != nil {
		errorCode <- "LOGIN_WRONG_PASSWORD_403"
		return &moderators, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &moderators, <-errorCode
}
