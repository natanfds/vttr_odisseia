package dtos

type CreateUser struct {
	Username    string `json:"user" validate:"required,min=3,max=50"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=50"`
	Password    string `json:"pass" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
}

type UpdateUser struct {
	DisplayName string `json:"display_name" validate:"omitempty,min=3,max=50"`
	Password    string `json:"pass" validate:"omitempty,min=3,max=50"`
	Email       string `json:"email" validate:"omitempty,email"`
}

type GetUser struct {
	Username    string `json:"user" validate:"omitempty,min=3,max=50"`
	DisplayName string `json:"display_name" validate:"omitempty,min=3,max=50"`
	Email       string `json:"email" validate:"omitempty,email"`
}
