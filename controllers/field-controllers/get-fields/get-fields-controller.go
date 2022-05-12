package getFieldsController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	GetFieldsService() (*[]models.Field, string)
}

type service struct {
	repository Repository
}

func NewGetFieldsService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetFieldsService() (*[]models.Field, string) {

	resultFields, errFields := s.repository.GetFieldsRepository()

	return resultFields, errFields
}