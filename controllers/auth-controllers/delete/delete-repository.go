package delete

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	DeleteRepository(inputDelete *InputDelete) (uint, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeleteRepository(inputDelete *InputDelete) (uint, string) {

	var users model.EntityUsers
	id := inputDelete.Id
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)
	checkUserAccount := db.Delete(&users, []uint{id})
	if checkUserAccount.RowsAffected < 1 {
		errorCode <- "ACCOUNT_NOT_FOUND_404"
		return id, <-errorCode
	}
	errorCode <- "nil"
	return id, <-errorCode
}
