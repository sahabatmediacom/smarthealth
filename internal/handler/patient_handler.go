package handler

import (
	"net/http"
	"pamer-api/helper"
	"pamer-api/internal/dto"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	service service.PatientService
}

func NewPatientHandler(service service.PatientService) *patientHandler {
	return &patientHandler{
		service: service,
	}
}

func (h *patientHandler) GetAll(c *gin.Context) {
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusNotImplemented,
		Message:    "Not implemented yet",
	})

	c.JSON(http.StatusNotImplemented, res)
}

func (h *patientHandler) GetPatientRecord(c *gin.Context) {
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusNotImplemented,
		Message:    "Not implemented yet",
	})

	c.JSON(http.StatusNotImplemented, res)
}

func (h *patientHandler) GetPatientData(c *gin.Context) {
	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusNotImplemented,
		Message:    "Not implemented yet",
	})

	c.JSON(http.StatusNotImplemented, res)
}
