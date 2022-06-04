package loginAuthController

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
)

type Service interface {
	UserLoginService(input *InputLogin) (*model.EntityUsers, string)
	OrganizationLoginService(input *InputLogin) (*model.EntityOrganizations, string)
	ModeratorLoginService(input *InputLogin) (*model.EntityModerators, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

/* User Login Service */
func (s *service) UserLoginService(input *InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.UserLoginRepository(&user)

	return resultLogin, errLogin
}

/* Organization Login Service */
func (s *service) OrganizationLoginService(input *InputLogin) (*model.EntityOrganizations, string) {

	organization := model.EntityOrganizations{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.OrganizationLoginRepository(&organization)

	return resultLogin, errLogin
}

/* Moderator Login Service */
func (s *service) ModeratorLoginService(input *InputLogin) (*model.EntityModerators, string) {

	moderator := model.EntityModerators{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.ModeratorLoginRepository(&moderator)

	return resultLogin, errLogin
}