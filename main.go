package main

import (
	"api/db"
	"api/internal/config"
	"api/internal/container"
	"api/internal/controllers"
	"api/internal/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	gormDb, err := db.InitializeDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	appContainer := container.NewContainer(gormDb, cfg)

	router := gin.Default()

	router.Use(utils.Cors)

	controllers.RegisterRoutes(router, appContainer)

	err = router.Run(":" + cfg.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
