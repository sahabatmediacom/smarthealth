package dto

type CreateUserRequest struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirm" validate:"required,eqfield=Password"`
	ParamedicID          string `json:"paramedic_id"`
}

type UpdateUserRequest struct {
	Username             string `json:"username" validate:"required"`
	PasswordConfirmation string `json:"password_confirm" validate:"required,eqfield=Password"`
	ParamedicID          string `json:"paramedic_id"`
}

type UserResponse struct {
	Username string `json:"username"`
}
