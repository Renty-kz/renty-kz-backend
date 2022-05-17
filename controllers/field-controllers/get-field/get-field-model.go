package getFieldController

type InputField struct {
	ID string `validate:"required,uuid"`
}