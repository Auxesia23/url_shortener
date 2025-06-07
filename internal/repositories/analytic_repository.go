package repository

import (
	"context"
	"log"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	"github.com/Auxesia23/url_shortener/internal/models"
	"gorm.io/gorm"
)

type AnalyticRepository interface {
	Create(ctx context.Context, analytic models.Analytic)error
	GetTotalClicks(ctx context.Context, shortUrl string)(int64, error)
	GetClicksPerDay(ctx context.Context, shortUrl string)([]mapper.DailyClickStat, error)
	GetClicksPerCountry(ctx context.Context, shortUrl string)([]mapper.ClickStat, error)
	GetClicksPerUserAgent(ctx context.Context, shortUrl string)([]mapper.ClickStat, error)
}

type analyticRepository struct {
	db *gorm.DB
}

func NewAnalyticRepository(db *gorm.DB) AnalyticRepository{
	return &analyticRepository{
		db: db,
	}
}

func(repo *analyticRepository)Create(ctx context.Context, analytic models.Analytic)error{
	err := repo.db.WithContext(ctx).Create(&analytic).Error
	if err != nil {
		log.Printf("Error : %s", err.Error())
		return err
	}

	return nil
}

func(repo *analyticRepository)GetTotalClicks(ctx context.Context, shortUrl string)(int64, error){
	var count int64
	err := repo.db.WithContext(ctx).Model(&models.Analytic{}).Where("shortened_url = ?", shortUrl).Count(&count).Error
	if err != nil {
	    return 0, err
	}
	return count, nil
}

func(repo *analyticRepository)GetClicksPerDay(ctx context.Context, shortUrl string)([]mapper.DailyClickStat, error){
	var clicksPerDay []mapper.DailyClickStat
	err := repo.db.WithContext(ctx).Model(&models.Analytic{}).
		Select("DATE(created_at) as date,COUNT(*) as count").
		Where("shortened_url = ?",shortUrl).
		Group("DATE(created_at)").
		Order("DATE(created_at) DESC").
		Limit(7).
		Scan(&clicksPerDay).Error
	
	if err != nil {
		return []mapper.DailyClickStat{}, err
	}
	
	return clicksPerDay, nil
}

func(repo *analyticRepository)GetClicksPerCountry(ctx context.Context, shortUrl string)([]mapper.ClickStat, error){
	var clickPerCountry []mapper.ClickStat
	err := repo.db.WithContext(ctx).Model(&models.Analytic{}).
		Select("country as name, COUNT(*) as count").
		Where("shortened_url = ?", shortUrl).
		Group("country").
		Scan(&clickPerCountry).Error
	if err != nil {
		return []mapper.ClickStat{}, err
	}
	
	return clickPerCountry, nil
}

func(repo *analyticRepository)GetClicksPerUserAgent(ctx context.Context, shortUrl string)([]mapper.ClickStat, error){
	var clickPerUserAgent []mapper.ClickStat
	
	err := repo.db.WithContext(ctx).Model(&models.Analytic{}).
		Select("user_agent as name, COUNT(*) as count").
		Where("shortened_url = ?", shortUrl).
		Group("user_agent").
		Scan(&clickPerUserAgent).Error
	if err != nil {
		return []mapper.ClickStat{}, err
	}
	
	return clickPerUserAgent, nil
}

