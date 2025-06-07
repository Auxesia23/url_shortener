package service

import (
	"context"
	"errors"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	repository "github.com/Auxesia23/url_shortener/internal/repositories"
	"github.com/Auxesia23/url_shortener/internal/utils"
)

type UrlService interface{
	CreateShortUrl(ctx context.Context, url mapper.UrlInput, userEmail string) (mapper.UrlResponse, error)
	GetUrl(ctx context.Context, email, shortUrl string)(mapper.UrlResponse, error)
	GetUrlByEmail(ctx context.Context, email string)(mapper.UrlListResponse, error)
	DeleteUrl(ctx context.Context, email,shortUrl string) error
}

type urlService struct{
	urlRepo repository.UrlRepository
}

func NewUrlService (urlRepo repository.UrlRepository) UrlService{
	return &urlService{
		urlRepo: urlRepo,
	}
}

func (service *urlService) CreateShortUrl(ctx context.Context, url mapper.UrlInput, userEmail string) (mapper.UrlResponse, error){
	input := mapper.ParseUrlInput(url, userEmail)
	if (!utils.ValidateUrl(input.Shortened)){
		return mapper.UrlResponse{}, errors.New("invalid short URL. Only aplhanumeric characters, hypens, and underscore are allowed")
	}
	err := service.urlRepo.Create(ctx, input)
	if err != nil {
		return mapper.UrlResponse{}, err
	}
	
	createdUrl, err := service.urlRepo.Read(ctx, url.Shortened)
	if err != nil {
		return mapper.UrlResponse{}, err
	}
	
	response := mapper.ParseUrlResponse(createdUrl)
	
	return response, nil
}

func(service *urlService) GetUrl(ctx context.Context, email,shortUrl string) (mapper.UrlResponse, error) {
	url, err := service.urlRepo.ReadByEmail(ctx, email,shortUrl)
	if err != nil {
		return mapper.UrlResponse{}, err
	}
	
	response := mapper.ParseUrlResponse(url)
	
	return response, nil
}

func (service *urlService) GetUrlByEmail(ctx context.Context, email string)(mapper.UrlListResponse, error){
	urls, err := service.urlRepo.ReadListByEmail(ctx, email)
	if err != nil {
		return mapper.UrlListResponse{}, err
	}
	
	response := mapper.ParseUrlListResponse(urls)
	
	return response, nil
}

func(service *urlService) DeleteUrl(ctx context.Context, email,shortUrl string) error{
	err := service.urlRepo.Delete(ctx, email, shortUrl)
	if err != nil {
		return err
	}
	
	return nil
}