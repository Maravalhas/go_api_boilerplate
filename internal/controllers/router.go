package controllers

import (
	"api/internal/container"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, container *container.Container) {
	apiGroup := r.Group("/api")

	RegisterAuthRoutes(apiGroup.Group("/auth"), container)
}
