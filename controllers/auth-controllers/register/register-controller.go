package registerAuthController

import (
	"fmt"

	model "github.com/KadirbekSharau/rentykz-backend/models"
)

type Service interface {
	ActiveUserRegisterService(input *InputRegister) (*model.EntityUsers, string)
	InactiveUserRegisterService(input *InputRegister) (*model.EntityUsers, string)
	AdminRegisterService(input *InputRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

/* Active User Registration Service */
func (s *service) ActiveUserRegisterService(input *InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}
	fmt.Printf(input.Password + "\n")
	resultRegister, errRegister := s.repository.ActiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}

/* Inactive User Registration Service */
func (s *service) InactiveUserRegisterService(input *InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.InactiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}

/* Admin Registration Service */
func (s *service) AdminRegisterService(input *InputRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}

	resultRegister, errRegister := s.repository.AdminRegisterRepository(&users)

	return resultRegister, errRegister
}