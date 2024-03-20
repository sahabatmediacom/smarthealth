package middleware

import (
	"pamer-api/helper"
	"pamer-api/internal/errorhandler"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(c *gin.Context) {
	tokenHeader := c.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: "not authorized"})
		c.Abort()
		return
	}

	tokenString := strings.Split(tokenHeader, " ")

	userId, err := helper.ValidateToken(tokenString[1])
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.UnauthorizedError{Message: err.Error()})
		c.Abort()
		return
	}

	c.Set("userID", userId)
	c.Next()
}
