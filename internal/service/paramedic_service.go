package service

import (
	"math"
	"pamer-api/internal/dto"
	"pamer-api/internal/entity"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/repository"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ParamedicService interface {
	Create(req *dto.ParamedicRequest) error
	Update(id int, req *dto.ParamedicRequest) error
	Delete(id int) error
	FindAll(params *dto.FilterParams) (*[]entity.Paramedic, *dto.Paginate, error)
	Detail(id int) (*entity.Paramedic, error)
}

type paramedicService struct {
	repository repository.ParamedicRepository
	validator  *validator.Validate
}

func NewParamedicService(r repository.ParamedicRepository) *paramedicService {
	return &paramedicService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *paramedicService) Create(req *dto.ParamedicRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{
			Message: err.Error(),
		}
	}

	// Membuat slice Hospitals berdasarkan HospitalIDs
	var hospitals []entity.Hospital
	for _, hospitalID := range req.Hospitals {
		hospitals = append(hospitals, entity.Hospital{Model: gorm.Model{ID: uint(hospitalID)}})
	}

	paramedic := entity.Paramedic{
		Name:        req.Name,
		IDSatusehat: req.IDSatusehat,
		Hospitals:   hospitals,
	}

	if err := s.repository.Create(&paramedic); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *paramedicService) Update(id int, req *dto.ParamedicRequest) error {
	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{
			Message: err.Error(),
		}
	}

	var hospitals []entity.Hospital
	for _, hospitalID := range req.Hospitals {
		hospitals = append(hospitals, entity.Hospital{Model: gorm.Model{ID: uint(hospitalID)}})
	}

	paramedic := entity.Paramedic{
		Name:        req.Name,
		IDSatusehat: req.IDSatusehat,
		Hospitals:   hospitals,
	}

	if err := s.repository.Update(id, &paramedic); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil

}

func (s *paramedicService) Delete(id int) error {
	return s.repository.Delete(id)
}

func (s *paramedicService) FindAll(params *dto.FilterParams) (*[]entity.Paramedic, *dto.Paginate, error) {
	total, err := s.repository.TotalData(params)
	if err != nil {
		return nil, nil, &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	paramedics, err := s.repository.FindAll(params)
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

	return paramedics, paginate, nil
}

func (s *paramedicService) Detail(id int) (*entity.Paramedic, error) {
	paramedic, err := s.repository.FindById(id)
	if err != nil {
		return nil, &errorhandler.NotFoundError{
			Message: err.Error(),
		}
	}

	return paramedic, nil
}
