package createFieldController

import (
	model "github.com/KadirbekSharau/rentykz-backend/models"
	//"github.com/lib/pq"
)

type Service interface {
	CreateFieldService(input *InputCreateField) (*model.EntityFields, string)
}

type service struct {
	repository Repository
}

func NewCreateFieldService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) CreateFieldService(input *InputCreateField) (*model.EntityFields, string) {
	fields := model.EntityFields{
		Name: input.Name,
		Address: input.Address,
		Price: input.Price,
		Capacity: input.Capacity,
		Description: input.Description,
		Contacts: input.Contacts,
		ImageLinks: input.ImageLinks,
		OrganizationID: input.OrganizationID,
		CityID: input.CityID,
		SportTypeID: input.SportTypeID,
		ModeratorID: input.ModeratorID,
	}

	resultCreateField, errCreateField := s.repository.CreateFieldRepository(&fields)

	return resultCreateField, errCreateField
}