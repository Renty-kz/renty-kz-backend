package getOrganizationsController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	GetActiveOrganizationsService() (*[]models.EntityOrganizations, string)
	GetInactiveOrganizationsService() (*[]models.EntityOrganizations, string)
	GetAllOrganizationsService() (*[]models.EntityOrganizations, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

/* Get All Organizations Service */
func (s *service) GetAllOrganizationsService() (*[]models.EntityOrganizations, string) {

	result, err := s.repository.GetAllOrganizationsRepository()

	return result, err
}

/* Get All Active Organizations Service */
func (s *service) GetActiveOrganizationsService() (*[]models.EntityOrganizations, string) {

	result, err := s.repository.GetActiveOrganizationsRepository()

	return result, err
}

/* Get All Inactive Organizations Service */
func (s *service) GetInactiveOrganizationsService() (*[]models.EntityOrganizations, string) {

	result, err := s.repository.GetInactiveOrganizationsRepository()

	return result, err
}