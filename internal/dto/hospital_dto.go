package dto

type HospitalRequest struct {
	Name string `json:"name" validate:"required"`
	IP   string `json:"ip" validate:"required"`
}

type HospitalResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	IP        string `json:"ip"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
