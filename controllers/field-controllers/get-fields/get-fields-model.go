package getFieldsController

type InputFieldByCityId struct {
	CityID string `validate:"required,uuid"`
}

type InputFieldsBySportTypeId struct {
	SportTypeID string `validate:"required,uuid"`
}

type InputFieldsByOrganizationId struct {
	OrganizationID string `validate:"required,uuid"`
}

type InputFieldsByModeratorId struct {
	ModeratorID string `validate:"required,uuid"`
}