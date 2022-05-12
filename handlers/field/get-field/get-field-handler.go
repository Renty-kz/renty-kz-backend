package getFieldService

import "github.com/KadirbekSharau/rentykz-backend/models"

type FieldService interface {
	Save(models.Field) models.Field
	FindAll() []models.Field
}

type fieldService struct {
	fields []models.Field
}

func New() FieldService {
	return &fieldService{}
}

func (service *fieldService) Save(field models.Field) models.Field {
	service.fields = append(service.fields, field)
	return field
} 
 
func (service *fieldService) FindAll() []models.Field {
	return service.fields
}