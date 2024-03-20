package repository

type PatientRepository interface {
	GetAll()
}

type patientRepository struct {
}

func NewPatientRepository() *patientRepository {
	return &patientRepository{}
}

func (r *patientRepository) GetAll() {

}
