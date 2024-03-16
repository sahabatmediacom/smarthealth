package router

import (
	"pamer-api/config"
	"pamer-api/internal/handler"
	"pamer-api/internal/repository"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

func ParamedicRouter(api *gin.RouterGroup) {
	paramedicRepository := repository.NewParamedicRepository(config.DB)
	paramedicService := service.NewParamedicService(paramedicRepository)
	paramedicHandler := handler.NewParamedicHandler(paramedicService)

	r := api.Group("/paramedics")
	r.GET("/", paramedicHandler.GetAll)
	r.GET("/:id", paramedicHandler.Get)
	r.POST("/", paramedicHandler.Create)
	r.PUT("/:id", paramedicHandler.Update)
	r.DELETE("/:id", paramedicHandler.Delete)
}
