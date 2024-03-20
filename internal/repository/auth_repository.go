package repository

import (
	"pamer-api/internal/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(*entity.User) error
	GetUserByUsername(username string) (*entity.User, error)
	UserExists(username string) bool
	ChangePassword(username, newPassword string) error
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(req *entity.User) error {
	// Cek apakah paramedicID sudah terisi
	// if req.ParamedicID != "" {
	// 	// Jika sudah terisi, cek apakah paramedis dengan ID tersebut ada
	// 	var paramedic entity.Paramedic
	// 	if err := r.db.First(&paramedic, "id = ?", req.ParamedicID).Error; err != nil {
	// 		// Jika paramedis tidak ditemukan, kembalikan error
	// 		return err
	// 	}
	// }

	// Simpan user ke database
	if err := r.db.Create(req).Error; err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	err := r.db.Model(&entity.User{}).First(&user, "username = ?", username).Error

	return &user, err
}

func (r *authRepository) UserExists(username string) bool {
	var user entity.User
	err := r.db.First(&user, "username = ?", username).Error

	return err == nil
}

func (r *authRepository) ChangePassword(username, newPassword string) error {
	var user entity.User
	if err := r.db.First(&user, "username = ?", username).Error; err != nil {
		return err
	}

	user.Password = newPassword
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}

	return nil
}
