package repository

import (
	"fmt"
	"pamer-api/internal/dto"
	"pamer-api/internal/entity"

	"gorm.io/gorm"
)

type ParamedicRepository interface {
	Create(request *entity.Paramedic) error
	Update(id int, request *entity.Paramedic) error
	Delete(id int) error
	FindById(id int) (*entity.Paramedic, error)
	FindAll(params *dto.FilterParams) (*[]entity.Paramedic, error)
	TotalData(params *dto.FilterParams) (int64, error)
}

type paramedicRepo struct {
	db *gorm.DB
}

func NewParamedicRepository(db *gorm.DB) *paramedicRepo {
	return &paramedicRepo{
		db: db,
	}
}

func (r *paramedicRepo) Create(paramedic *entity.Paramedic) error {
	newParamedic := &entity.Paramedic{
		Name:        paramedic.Name,
		IDSatusehat: paramedic.IDSatusehat,
		Hospitals:   []entity.Hospital{},
	}

	fmt.Println(paramedic.Hospitals)
	// Tambahkan relasi Hospitals
	for _, hospitalID := range paramedic.Hospitals {
		hospital := &entity.Hospital{}
		r.db.First(hospital, hospitalID) // Mengambil data hospital dari database berdasarkan ID
		newParamedic.Hospitals = append(newParamedic.Hospitals, hospitalID)
	}

	// Cek apakah Hospitals sudah terisi dengan benar
	fmt.Println("Hospitals:", newParamedic.Hospitals)

	// Simpan paramedic ke database
	if err := r.db.Create(newParamedic).Error; err != nil {
		return err
	}

	// Simpan relasi many-to-many ke tabel paramedic_hospitals
	if err := r.db.Model(&newParamedic).Association("Hospitals").Append(&newParamedic.Hospitals); err != nil {
		return err
	}

	return nil
}

func (r *paramedicRepo) Update(id int, paramedic *entity.Paramedic) error {
	// Cari paramedic berdasarkan ID
	existingParamedic := &entity.Paramedic{}
	if err := r.db.Preload("Hospitals").First(existingParamedic, id).Error; err != nil {
		return err
	}

	// Update nilai-nilai yang berubah
	existingParamedic.Name = paramedic.Name
	existingParamedic.IDSatusehat = paramedic.IDSatusehat

	// Hapus semua relasi hospitals yang ada sebelumnya
	if err := r.db.Model(existingParamedic).Association("Hospitals").Clear(); err != nil {
		return err
	}

	// Tambahkan kembali relasi Hospitals yang baru
	for _, hospitalID := range paramedic.Hospitals {
		hospital := &entity.Hospital{}
		if err := r.db.First(hospital, hospitalID).Error; err != nil {
			return err
		}
		existingParamedic.Hospitals = append(existingParamedic.Hospitals, *hospital)
	}

	// Simpan perubahan ke database
	if err := r.db.Save(existingParamedic).Error; err != nil {
		return err
	}

	return nil
}

func (r *paramedicRepo) Delete(id int) error {
	var paramedic entity.Paramedic
	err := r.db.Delete(&paramedic, id).Error

	// Hapus relasi antara paramedic dan rumah sakit (Hospitals) dari tabel paramedic_hospitals
	if err := r.db.Exec("DELETE FROM paramedic_hospitals WHERE paramedic_id = ?", id).Error; err != nil {
		return err
	}

	return err
}

func (r *paramedicRepo) FindById(id int) (*entity.Paramedic, error) {
	var paramedic *entity.Paramedic
	err := r.db.Model(&entity.Paramedic{}).Preload("Hospitals").First(&paramedic, "id = ?", id).Error

	return paramedic, err
}

func (r *paramedicRepo) FindAll(params *dto.FilterParams) (*[]entity.Paramedic, error) {
	var paramedics *[]entity.Paramedic
	query := r.db.Model(&entity.Paramedic{}).Preload("Hospitals").Select("id, name, id_satusehat, created_at")

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Offset(params.Offset).Limit(params.Limit).Find(&paramedics).Error

	return paramedics, err
}

func (r *paramedicRepo) TotalData(params *dto.FilterParams) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Paramedic{})

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Count(&count)
	if err != nil {
		return count, err.Error
	}

	return count, nil
}
