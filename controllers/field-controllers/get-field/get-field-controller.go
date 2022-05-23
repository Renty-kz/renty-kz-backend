package getFieldController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	GetFieldByIdService(input *InputField) (*models.EntityFields, string)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetFieldByIdService(input *InputField) (*models.EntityFields, string) {

	field := InputField{
		ID: input.ID,
	}
	return s.repository.GetFieldByIdRepository(&field)
}