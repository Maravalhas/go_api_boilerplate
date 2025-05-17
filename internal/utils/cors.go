package utils

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCors(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("ORIGIN")}
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
}
