package getFieldsController

import "github.com/KadirbekSharau/rentykz-backend/models"

type Service interface {
	GetFieldsService() (*[]models.EntityFields, string)
	GetFieldsByCityIdService(input *InputFieldByCityId) (*[]models.EntityFields, string)
	GetFieldsBySportTypeIdService(input *InputFieldsBySportTypeId) (*[]models.EntityFields, string)
	GetFieldsByOrganizationIdService(input *InputFieldsByOrganizationId) (*[]models.EntityFields, string)
	GetFieldsByModeratorIdService(input *InputFieldsByModeratorId) (*[]models.EntityFields, string)
}

type service struct {
	repository Repository
}

func NewGetFieldsService(repository Repository) *service {
	return &service{repository: repository}
}

/* Get All Fields Service */
func (s *service) GetFieldsService() (*[]models.EntityFields, string) {

	resultFields, errFields := s.repository.GetFieldsRepository()

	return resultFields, errFields
}

/* Get Fields By City Id Service */
func (s *service) GetFieldsByCityIdService(input *InputFieldByCityId) (*[]models.EntityFields, string) {

	city := InputFieldByCityId{
		CityID: input.CityID,
	}

	return s.repository.GetFieldsByCityIdRepository(&city)
}

/* Get Fields By Sport Type Id Service */
func (s *service) GetFieldsBySportTypeIdService(input *InputFieldsBySportTypeId) (*[]models.EntityFields, string) {

	sportType := InputFieldsBySportTypeId{
		SportTypeID: input.SportTypeID,
	}

	return s.repository.GetFieldsBySportTypeIdRepository(&sportType)
}

/* Get Fields By Organization Id Service */
func (s *service) GetFieldsByOrganizationIdService(input *InputFieldsByOrganizationId) (*[]models.EntityFields, string) {

	organization := InputFieldsByOrganizationId{
		OrganizationID: input.OrganizationID,
	}

	return s.repository.GetFieldsByOrganizationIdRepository(&organization)
}

/* Get Fields By Moderator Id Service */
func (s *service) GetFieldsByModeratorIdService(input *InputFieldsByModeratorId) (*[]models.EntityFields, string) {

	moderator := InputFieldsByModeratorId{
		ModeratorID: input.ModeratorID,
	}

	return s.repository.GetFieldsByModeratorIdRepository(&moderator)
}