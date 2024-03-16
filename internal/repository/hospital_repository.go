package repository

import (
	"fmt"

	"pamer-api/internal/dto"
	"pamer-api/internal/entity"

	"gorm.io/gorm"
)

type HospitalRepository interface {
	TotalData(params *dto.FilterParams) (int64, error)
	Create(hospital *entity.Hospital) error
	Update(id int, hospital *entity.Hospital) error
	Delete(id int) error
	FindAll(params *dto.FilterParams) (*[]dto.HospitalResponse, error)
	FindById(id int) (*dto.HospitalResponse, error)
}

type hospitalRepo struct {
	db *gorm.DB
}

func NewHospitalRepo(db *gorm.DB) *hospitalRepo {
	return &hospitalRepo{
		db: db,
	}
}

func (r *hospitalRepo) TotalData(params *dto.FilterParams) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Hospital{})

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Count(&count)
	if err != nil {
		return count, err.Error
	}

	return count, nil
}

func (r *hospitalRepo) Create(hospital *entity.Hospital) error {
	err := r.db.Create(&hospital).Error
	return err
}

func (r *hospitalRepo) Update(id int, hospital *entity.Hospital) error {
	err := r.db.Model(&hospital).Where("id = ?", id).Updates(&hospital).Error

	return err
}

func (r *hospitalRepo) Delete(id int) error {
	var hospital entity.Hospital
	err := r.db.Delete(&hospital, id).Error

	return err
}

func (r *hospitalRepo) FindAll(params *dto.FilterParams) (*[]dto.HospitalResponse, error) {
	var hospitalsResponse []dto.HospitalResponse
	query := r.db.Model(&entity.Hospital{}).Select("id, name,ip, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at")

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Offset(params.Offset).Limit(params.Limit).Find(&hospitalsResponse).Error

	return &hospitalsResponse, err
}

func (r *hospitalRepo) FindById(id int) (*dto.HospitalResponse, error) {
	var hospital dto.HospitalResponse
	query := r.db.Model(&entity.Hospital{}).Select("id, name,ip, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at")
	err := query.First(&hospital, "id = ?", id).Error

	return &hospital, err
}
