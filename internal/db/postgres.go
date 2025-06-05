package db

import (
	"fmt"
	"os"

	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres() (*gorm.DB, error) {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Url{}, &models.User{}, &models.Analytic{})
	if err != nil {
		return nil, err
	}

	return db, nil
}