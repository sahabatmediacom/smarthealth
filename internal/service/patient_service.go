package service

import (
	"pamer-api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type PatientService interface {
	FindAll()
}

type patientService struct {
	repository repository.PatientRepository
	validator  *validator.Validate
}

func NewPatientService(r repository.PatientRepository) *patientService {
	return &patientService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *patientService) FindAll() {}
