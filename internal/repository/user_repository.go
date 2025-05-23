package repository

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface{
	Create(ctx context.Context, user models.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (repo *userRepository) Create(ctx context.Context,user models.User) error {
	err := repo.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}
