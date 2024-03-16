package handler

import (
	"fmt"
	"net/http"
	"pamer-api/helper"
	"pamer-api/internal/dto"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *authHandler {
	return &authHandler{
		service: s,
	}
}

func (h *authHandler) Register(c *gin.Context) {
	var register dto.RegisterRequest
	if err := c.ShouldBindJSON(&register); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Register(&register); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "register success, please login",
	})

	c.JSON(http.StatusCreated, res)

}

func (h *authHandler) Login(c *gin.Context) {
	var login dto.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	result, err := h.service.Login(&login)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully login",
		Data:       result,
	})

	c.JSON(http.StatusOK, res)
}

func (h *authHandler) ChangePassword(c *gin.Context) {
	var changePwd dto.ChangePasswordRequest
	if err := c.ShouldBindJSON(&changePwd); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	err := h.service.ChangePassword(&changePwd)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("User with username %v, successfully changed password", changePwd.Username),
	})

	c.JSON(http.StatusOK, res)
}
