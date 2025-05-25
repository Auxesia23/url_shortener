package repository

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface{
	Create(ctx context.Context, user models.User) error
	Read(ctx context.Context, email string) (models.User, error)
	Delete(ctx context.Context, email string) error
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

func (repo *userRepository) Read(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := repo.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (repo *userRepository) Delete(ctx context.Context, email string) error {
	err := repo.db.WithContext(ctx).Where("email = ?", email).Delete(&models.User{}).Error
	if err != nil {
		return err
	}
	return nil
}
