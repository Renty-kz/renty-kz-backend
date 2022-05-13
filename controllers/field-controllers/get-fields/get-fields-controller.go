package getFieldsController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	GetFieldsService() (*[]models.EntityFields, string)
}

type service struct {
	repository Repository
}

func NewGetFieldsService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetFieldsService() (*[]models.EntityFields, string) {

	resultFields, errFields := s.repository.GetFieldsRepository()

	return resultFields, errFields
}