package dtos

type CreateUser struct {
	Username    string `json:"user" validate:"required,min=3,max=50"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=50"`
	Password    string `json:"pass" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
}
