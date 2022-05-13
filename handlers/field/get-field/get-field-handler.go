package getFieldService

import "github.com/KadirbekSharau/rentykz-backend/models"

type FieldService interface {
	Save(models.EntityFields) models.EntityFields
	FindAll() []models.EntityFields
}

type fieldService struct {
	fields []models.EntityFields
}

func New() FieldService {
	return &fieldService{}
}

func (service *fieldService) Save(field models.EntityFields) models.EntityFields {
	service.fields = append(service.fields, field)
	return field
} 
 
func (service *fieldService) FindAll() []models.EntityFields {
	return service.fields
}