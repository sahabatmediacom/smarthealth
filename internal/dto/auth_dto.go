package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=15"`
}

type RegisterRequest struct {
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,min=6"`
	PasswordConfirmation string `json:"password_confirm" validate:"required,eqfield=Password"`
	ParamedicID          string `json:"paramedic_id"`
}

type LoginResponse struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	ParamedicID string `json:"paramedic_id"`
	Token       string `json:"token"`
}

type ChangePasswordRequest struct {
	Username    string `json:"username" validate:"required"`
	OldPassword string `json:"current_password" validate:"required, min=6"`
	NewPassword string `json:"new_password" validate:"required, min=6"`
}
