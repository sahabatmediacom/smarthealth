package dto

type ParamedicRequest struct {
	Name        string `json:"name" validate:"required"`
	Hospitals   []int  `json:"hospitals"`
	IDSatusehat string `json:"id_satusehat" validate:"required"`
}

type ParamedicResponse struct {
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	Hospitals   []HospitalResponse `json:"hospitals"`
	IDSatusehat string             `json:"id_satusehat"`
}
