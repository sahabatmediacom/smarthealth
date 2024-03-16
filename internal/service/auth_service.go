package service

import (
	"pamer-api/helper"
	"pamer-api/internal/dto"
	"pamer-api/internal/entity"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/repository"

	"github.com/go-playground/validator/v10"
)

type AuthService interface {
	Register(req *dto.RegisterRequest) error
	Login(req *dto.LoginRequest) (*dto.LoginResponse, error)
	ChangePassword(req *dto.ChangePasswordRequest) error
}

type authService struct {
	repository repository.AuthRepository
	validator  *validator.Validate
}

func NewAuthService(r repository.AuthRepository) *authService {
	return &authService{
		repository: r,
		validator:  validator.New(),
	}
}

func (s *authService) Register(req *dto.RegisterRequest) error {

	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	if userExists := s.repository.UserExists(req.Username); userExists {
		return &errorhandler.BadRequestError{Message: "User already registered"}
	}

	passwordHash, err := helper.HashPassword(req.Password)
	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	user := entity.User{
		Username:    req.Username,
		Password:    passwordHash,
		ParamedicID: req.ParamedicID,
	}

	if err := s.repository.Register(&user); err != nil {
		return &errorhandler.InternalServerError{
			Message: err.Error(),
		}
	}

	return nil
}

func (s *authService) Login(req *dto.LoginRequest) (*dto.LoginResponse, error) {
	var data dto.LoginResponse
	if err := s.validator.Struct(req); err != nil {
		return nil, &errorhandler.BadRequestError{Message: err.Error()}
	}

	var user *entity.User

	user, err := s.repository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "user tidak ditemukan"}
	}

	if passwordValid := helper.ComparePassword(user.Password, req.Password); !passwordValid {
		return nil, &errorhandler.BadRequestError{Message: "Password salah"}
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	data = dto.LoginResponse{
		ID:          int(user.ID),
		Username:    user.Username,
		Token:       token,
		ParamedicID: user.ParamedicID,
	}

	return &data, nil
}

func (s *authService) ChangePassword(req *dto.ChangePasswordRequest) error {

	if err := s.validator.Struct(req); err != nil {
		return &errorhandler.BadRequestError{Message: err.Error()}
	}

	var user *entity.User
	user, err := s.repository.GetUserByUsername(req.Username)
	if err != nil {
		return &errorhandler.NotFoundError{Message: "User tidak ditemukan"}
	}

	if passwordValid := helper.ComparePassword(user.Password, req.OldPassword); !passwordValid {
		return &errorhandler.BadRequestError{Message: "Invalid Password"}
	}

	if err := s.repository.ChangePassword(user.Username, req.NewPassword); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil

}
