package controllers

import (
	"api/internal/container"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, container *container.Container) {
	router := r.Group("/api")

	RegisterAuthRoutes(router.Group("/auth"), container)
}
