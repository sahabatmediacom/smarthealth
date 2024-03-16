package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"pamer-api/helper"
	"pamer-api/internal/dto"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

type hospitalHandler struct {
	service service.HospitalService
}

func NewHospitalHandler(service service.HospitalService) *hospitalHandler {
	return &hospitalHandler{
		service: service,
	}
}

func (h *hospitalHandler) GetAll(c *gin.Context) {
	filter := helper.FilterParams(c)
	hospitals, paginate, err := h.service.FindAll(filter)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "List Hospital's",
		Paginate:   paginate,
		Data:       hospitals,
	})

	c.JSON(http.StatusOK, res)
}

func (h *hospitalHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	hospital, err := h.service.Detail(idInt)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.NotFoundError{
			Message: err.Error(),
		})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Data Paramedic with id %v", idInt),
		Data:       hospital,
	})

	c.JSON(http.StatusOK, res)

}

func (h *hospitalHandler) Create(c *gin.Context) {
	var hospital dto.HospitalRequest
	if err := c.ShouldBindJSON(&hospital); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Create(&hospital); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "hospital created",
	})

	c.JSON(http.StatusOK, res)
}

func (h *hospitalHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	var hospital dto.HospitalRequest
	if err := c.ShouldBindJSON(&hospital); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Update(idInt, &hospital); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "hospital successfully updated",
	})

	c.JSON(http.StatusOK, res)
}

func (h *hospitalHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	if err := h.service.Delete(idInt); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "hospital deleted",
	})

	c.JSON(http.StatusOK, res)
}
