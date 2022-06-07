package updateOrganizationController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	ActivateOrganizationByIdService(input *InputActivateOrganization) (*models.EntityOrganizations, string)
	InactivateOrganizationByIdService(input *InputActivateOrganization) (*models.EntityOrganizations, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

/* Activate Organization By Id Service */
func (s *service) ActivateOrganizationByIdService(input *InputActivateOrganization) (*models.EntityOrganizations, string) {

	result, err := s.repository.ActivateOrganizationByIdRepository(input)

	return result, err
}

/* Inactivate Organization By Id Service */
func (s *service) InactivateOrganizationByIdService(input *InputActivateOrganization) (*models.EntityOrganizations, string) {

	result, err := s.repository.InactivateOrganizationByIdRepository(input)

	return result, err
}