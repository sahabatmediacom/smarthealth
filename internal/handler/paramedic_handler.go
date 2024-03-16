package handler

import (
	"fmt"
	"net/http"
	"pamer-api/helper"
	"pamer-api/internal/dto"
	"pamer-api/internal/errorhandler"
	"pamer-api/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type paramedicHandler struct {
	service service.ParamedicService
}

func NewParamedicHandler(service service.ParamedicService) *paramedicHandler {
	return &paramedicHandler{
		service: service,
	}
}

func (h *paramedicHandler) GetAll(c *gin.Context) {
	filter := helper.FilterParams(c)
	paramedics, paginate, err := h.service.FindAll(filter)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "List Paramedic",
		Paginate:   paginate,
		Data:       paramedics,
	})

	c.JSON(http.StatusOK, res)

}

func (h *paramedicHandler) Get(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	paramedic, err := h.service.Detail(idInt)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.NotFoundError{
			Message: err.Error(),
		})
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Data Paramedic with id %v", idInt),
		Data:       paramedic,
	})

	c.JSON(http.StatusOK, res)
}

func (h *paramedicHandler) Create(c *gin.Context) {
	var paramedic dto.ParamedicRequest
	if err := c.ShouldBindJSON(&paramedic); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Create(&paramedic); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "paramedic added",
	})

	c.JSON(http.StatusCreated, res)
}

func (h *paramedicHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	var paramedic dto.ParamedicRequest
	if err := c.ShouldBindJSON(&paramedic); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	if err := h.service.Update(idInt, &paramedic); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Paramedic with id %v updated", idInt),
	})

	c.JSON(http.StatusOK, res)
}

func (h *paramedicHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	idInt, _ := strconv.Atoi(idStr)

	if err := h.service.Delete(idInt); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	res := helper.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Paramedic data deleted",
	})

	c.JSON(http.StatusOK, res)
}
