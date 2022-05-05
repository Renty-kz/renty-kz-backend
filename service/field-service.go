package service

import "github.com/KadirbekSharau/rentykz-backend/entity"

type FieldService interface {
	Save(entity.Field) entity.Field
	FindAll() []entity.Field
}

type fieldService struct {
	fields []entity.Field
}

func New() FieldService {
	return &fieldService{}
}

func (service *fieldService) Save(field entity.Field) entity.Field {
	service.fields = append(service.fields, field)
	return field
} 
 
func (service *fieldService) FindAll() []entity.Field {
	return service.fields
}