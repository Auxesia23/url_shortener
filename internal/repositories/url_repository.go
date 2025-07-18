package repository

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/gorm"
)

type UrlRepository interface {
	Create(ctx context.Context, url models.Url) error
	Read(ctx context.Context, shortUrl string) (models.Url, error)
	ReadByEmail(ctx context.Context, email, shortenedUrl string) (models.Url, error)
	ReadListByEmail(ctx context.Context, email string) ([]models.Url, error)
	Delete(ctx context.Context, email, shortUrl string) error
}

type urlRepository struct {
	db *gorm.DB
}

func NewUrlRepository(db *gorm.DB) UrlRepository {
	return &urlRepository{
		db: db,
	}
}

func (repo *urlRepository) Create(ctx context.Context, url models.Url) error {
	err := repo.db.WithContext(ctx).Create(&url).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *urlRepository) Read(ctx context.Context, shortUrl string) (models.Url, error) {
	var url models.Url
	err := repo.db.WithContext(ctx).Where("shortened = ?", shortUrl).First(&url).Error
	if err != nil {
		return models.Url{}, err
	}
	return url, nil
}

func (repo *urlRepository) ReadByEmail(ctx context.Context, email, shortenedUrl string) (models.Url, error) {
	var url models.Url
	err := repo.db.WithContext(ctx).Where("user_email =?  AND shortened = ?", email, shortenedUrl).First(&url).Error
	if err != nil {
		return models.Url{}, err
	}
	return url, nil
}

func (repo *urlRepository) ReadListByEmail(ctx context.Context, email string) ([]models.Url, error) {
	var urls []models.Url
	err := repo.db.WithContext(ctx).Where("user_email = ?", email).Order("created_at DESC").Find(&urls).Error
	if err != nil {
		return []models.Url{}, err
	}
	return urls, nil
}

func (repo *urlRepository) Delete(ctx context.Context, email, shortUrl string) error {
	result := repo.db.WithContext(ctx).Unscoped().Where("user_email = ? AND shortened = ?", email, shortUrl).Delete(&models.Url{})
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
