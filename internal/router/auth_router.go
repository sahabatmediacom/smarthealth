package router

import (
	"pamer-api/config"
	"pamer-api/internal/handler"
	"pamer-api/internal/repository"
	"pamer-api/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	authRepository := repository.NewAuthRepository(config.DB)
	authService := service.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	r := api.Group("/auth")
	r.POST("/login", authHandler.Login)
	r.POST("/register", authHandler.Register)
	r.PUT("/changepwd", authHandler.ChangePassword)
}
