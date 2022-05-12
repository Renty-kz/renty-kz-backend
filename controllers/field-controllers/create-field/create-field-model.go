package createFieldController

type InputCreateField struct {
	Name string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	City string `json:"city" validate:"required"`
	Contacts string `json:"contacts" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price string `json:"price" validate:"required"`
	Capacity string `json:"capacity" validate:"required"`
}