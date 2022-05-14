package createAdminController

type InputCreateAdmin struct {
	Fullname string `json:"FullName"`
	Email string `json:"Email"`
	Password  string `json:"Password"`
	Active    bool   `json:"Active"`
	IsAdmin    bool   `json:"IsAdmin"`
}