package updateOrganizationController

type InputActivateOrganization struct {
	ID string `json:"id" validate:"required,uuid"`
}