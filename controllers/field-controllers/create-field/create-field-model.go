package createFieldController

type InputCreateField struct {
	Name string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Contacts []string 
	ImageLinks []string
	Description string `json:"description" validate:"required"`
	Price uint `json:"price" validate:"required"`
	Capacity uint `json:"capacity" validate:"required"`
	OrganizationID uint
	CityID uint
	SportTypeID uint
	ModeratorID uint
}