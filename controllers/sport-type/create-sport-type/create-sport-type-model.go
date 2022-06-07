package createSportTypeController

type InputCreateSportType struct {
	Name string `json:"sport_type" validate:"required"`
}