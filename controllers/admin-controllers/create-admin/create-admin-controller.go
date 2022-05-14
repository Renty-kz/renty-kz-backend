package createAdminController

import model "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	CreateAdminService(input *InputCreateAdmin) (*model.EntityFields, string)
}

type service struct {
	repository Repository
}

func NewServiceCreate(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateAdminService(input *InputCreateAdmin) (*model.EntityUsers, string) {

	admins := model.EntityUsers{
		Fullname: input.Fullname,
		Email: input.Email,
		Password: input.Password,
		Active: input.Active,
		IsAdmin: input.IsAdmin,
	}

	resultCreateAdmin, errCreateAdmin := s.repository.CreateAdminRepository(&admins)

	return resultCreateAdmin, errCreateAdmin
}