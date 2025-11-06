package controllers

import (
	"api/internal/container"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.RouterGroup, container *container.Container) {
	router.GET("", container.AuthHandler.GetTokenData)
	router.POST("", container.AuthHandler.PostToken)
	router.POST("/refresh", container.AuthHandler.PostRefreshToken)
}
