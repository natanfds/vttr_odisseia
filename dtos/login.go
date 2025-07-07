package dtos

type Login struct {
	Username string `json:"user" validate:"required,min=3,max=50"`
	Password string `json:"pass" validate:"required,min=3,max=50"`
}
