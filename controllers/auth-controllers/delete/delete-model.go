package delete

type InputDelete struct {
	Id uint `json:"id" validate:"required,id"`
}
