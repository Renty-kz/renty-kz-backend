package createFieldController

type InputCreateField struct {
	Name string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Contacts []string `json:"contacts" validate:"required"`
	ImageLinks []string `json:"imageLinks" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price uint `json:"price" validate:"required"`
	Capacity uint `json:"capacity" validate:"required"`
	OrganizationID uint `json:"organization_id" validate:"required"`
	CityID uint `json:"city_id" validate:"required"`
	SportTypeID uint `json:"sport_type_id" validate:"required"`
	ModeratorID uint `json:"moderator_id" validate:"required"`
}