package main

import (
	"log"
	"movie-festival-app/config"
	"movie-festival-app/models"
)

func main() {
	// Initialize configuration
	if err := config.InitConfig(); err != nil {
		log.Fatalf("Could not initialize config: %v", err)
	}

	// Run migrations
	err := config.DB.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Vote{},
		&models.ViewLog{},
	)

	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	log.Println("Migration completed successfully!")
}
