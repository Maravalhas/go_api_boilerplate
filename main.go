package main

import (
	"api/db"
	"api/internal/config"
	"api/internal/container"
	"api/internal/controllers"
	"api/internal/logger"
	"api/internal/utils"

	"github.com/gin-gonic/gin"
)

var log = logger.New("main")

func main() {
	config.LoadConfig()

	gormDb, err := db.InitializeDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	appContainer := container.NewContainer(gormDb)

	router := gin.Default()

	router.Use(utils.Cors)

	controllers.RegisterRoutes(router, appContainer)

	err = router.Run(":" + config.Current.Port)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
