package createCityController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	CreateCityService(input *InputCreateCity) (*models.EntityCities, string)
}

type service struct {
	repository Repository
}

func NewCreateCityService(repository Repository) Service {
	return &service{repository: repository}
}

func (s *service) CreateCityService(input *InputCreateCity) (*models.EntityCities, string) {
	city := models.EntityCities{
		Name: input.Name,
	}

	// newCity, errCity 
	return s.repository.CreateCityRepository(&city)
	// return newCity, errCity
}