package createSportTypeController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	CreateSportTypeService(input *InputCreateSportType) (*models.EntitySportTypes, string)
}

type service struct {
	repository Repository
}

func NewCreateSportTypeService(repository Repository) Service {
	return &service{repository: repository}
}

/* Create sport type controller */
func (s *service) CreateSportTypeService(input *InputCreateSportType) (*models.EntitySportTypes, string) {
	sportType := models.EntitySportTypes{
		Name: input.Name,
	}

	return s.repository.CreateSportTypeRepository(&sportType)
}