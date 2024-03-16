package service

import (
	"math"
	"time"

	"pamer-api/internal/dto"
	"pamer-api/internal/entity"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type HospitalService interface {
	Create(req *dto.HospitalRequest) error
	Update(id int, req *dto.HospitalRequest) error
	Delete(id int) error
	FindAll(params *dto.FilterParams) (*[]dto.HospitalResponse, *dto.Paginate, error)
	Detail(id int) (*dto.HospitalResponse, error)
}

type hospitalService struct {
	repository repository.HospitalRepository
	validator  *validator.Validate
}

func NewHospitalService(r repository.HospitalRepository) *hospitalService {
	return &hospitalService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *hospitalService) Create(req *dto.HospitalRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{
			Message: err.Error(),
		}
	}

	hospital := entity.Hospital{
		Name: req.Name,
		IP:   req.IP,
	}

	if err := s.repository.Create(&hospital); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *hospitalService) FindAll(params *dto.FilterParams) (*[]dto.HospitalResponse, *dto.Paginate, error) {
	total, err := s.repository.TotalData(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	hospitals, err := s.repository.FindAll(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	paginate := &dto.Paginate{
		Total:     int(total),
		PerPage:   params.Limit,
		Page:      params.Page,
		TotalPage: int(math.Ceil(float64(total) / float64(params.Limit))),
	}

	return hospitals, paginate, nil
}

func (s *hospitalService) Detail(id int) (*dto.HospitalResponse, error) {
	hospital, err := s.repository.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}

	return hospital, nil
}

func (s *hospitalService) Update(id int, req *dto.HospitalRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	hospital := entity.Hospital{
		Name:      req.Name,
		IP:        req.IP,
		UpdatedAt: time.Now(),
	}

	if err := s.repository.Update(id, &hospital); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (s *hospitalService) Delete(id int) error {
	if err := s.repository.Delete(id); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}
