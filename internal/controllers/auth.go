package controllers

import (
	"api/internal/container"

	"github.com/gin-gonic/gin"
)

/* PATH: /auth */
func RegisterAuthRoutes(authGroup *gin.RouterGroup, container *container.Container) {
	authGroup.GET("", container.AuthHandler.GetTokenData)
	authGroup.POST("", container.AuthHandler.PostToken)
	authGroup.POST("/refresh", container.AuthHandler.PostRefreshToken)
}
