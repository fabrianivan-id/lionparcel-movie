package config

import (
	"fmt"
	"movie-festival-app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitConfig() error {
	// dsn := os.Getenv("DATABASE_URL")
	// if dsn == "" {
	// 	return errors.New("DATABASE_URL not set in .env file")
	// }

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open("postgresql://postgres:Ivan2102@localhost:5432/lp-movie"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}

	// Run migrations
	return DB.AutoMigrate(&models.Movie{}, &models.User{}, &models.Vote{}, &models.ViewLog{})
}
