package db

import (
	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
  db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
  db.AutoMigrate(&models.Url{},&models.User{})
  if err != nil {
    return nil, err
  }
  return db, nil
}