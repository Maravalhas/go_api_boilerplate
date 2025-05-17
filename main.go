package main

import (
	"api/db"
	"api/internal/container"
	"api/internal/routes"
	"api/internal/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	godotenv.Load()

	// Initialize and connect to the database
	db, err := db.InitializeDB()

	if err != nil {
		panic(err)
	}

	container := container.NewContainer(db)

	router := gin.Default()

	// Enable CORS
	utils.UseCors(router)

	// Register routes
	routes.RegisterRoutes(router, container)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start the server
	router.Run(":" + port)
}
