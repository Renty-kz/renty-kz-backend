package getFieldController

type InputField struct {
	ID string `json:"id" validate:"required,uuid"`
}