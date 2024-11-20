package main

import (
	"log"
	"movie-festival-app/config"
	"movie-festival-app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize configuration
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Could not initialize config: %v", err)
	}

	// Initialize router
	router := gin.Default()

	// Register routes
	routes.RegisterRoutes(router)

	// Start the server
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
