package models

type Organization struct {
	Name string `json:"name"`
	Email string `json:"email"`
	Field Field `json:"field"`
}

type Field struct {
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	Contacts string `json:"contacts"`
	Description string `json:"description"`
	Price string `json:"price"`
	Capacity string `json:"capacity"`
}	