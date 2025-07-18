package service

import (
	"context"
	"errors"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	repository "github.com/Auxesia23/url_shortener/internal/repositories"
	"github.com/Auxesia23/url_shortener/internal/utils"
)

type UrlService interface {
	CreateShortUrl(ctx context.Context, url mapper.UrlInput, userEmail string) (mapper.UrlResponse, error)
	GetUrl(ctx context.Context, email, shortUrl string) (mapper.UrlResponse, error)
	GetUrlByEmail(ctx context.Context, email string) (mapper.UrlListResponse, error)
	DeleteUrl(ctx context.Context, email, shortUrl string) error
}

type urlService struct {
	urlRepo repository.UrlRepository
}

func NewUrlService(urlRepo repository.UrlRepository) UrlService {
	return &urlService{
		urlRepo: urlRepo,
	}
}

func (service *urlService) CreateShortUrl(ctx context.Context, url mapper.UrlInput, userEmail string) (mapper.UrlResponse, error) {
	var (
		input mapper.UrlInput
		err error
	)

	input.Original, err = utils.ValidateOriginalUrl(url.Original)
	if err != nil {
		return mapper.UrlResponse{}, err
	}

	ok := utils.VerifyShortenedUrl(url.Shortened)
	if ok {
		input.Shortened = url.Shortened
	} else {
		return mapper.UrlResponse{}, errors.New("shortened URL must be 5-20 alphanumeric characters or hyphens/underscores")
	}

	model := mapper.ParseUrlInput(input, userEmail)

	err = service.urlRepo.Create(ctx, model)
	if err != nil {
		return mapper.UrlResponse{}, errors.New("the provided URL is already in use")
	}

	createdUrl, err := service.urlRepo.Read(ctx, url.Shortened)
	if err != nil {
		return mapper.UrlResponse{}, errors.New("failed to retrieve the created URL")
	}

	response := mapper.ParseUrlResponse(createdUrl)

	return response, nil
}

func (service *urlService) GetUrl(ctx context.Context, email, shortUrl string) (mapper.UrlResponse, error) {
	url, err := service.urlRepo.ReadByEmail(ctx, email, shortUrl)
	if err != nil {
		return mapper.UrlResponse{}, errors.New("short URL not found or does not belong to the user")
	}

	response := mapper.ParseUrlResponse(url)

	return response, nil
}

func (service *urlService) GetUrlByEmail(ctx context.Context, email string) (mapper.UrlListResponse, error) {
	urls, err := service.urlRepo.ReadListByEmail(ctx, email)
	if err != nil {
		return mapper.UrlListResponse{}, errors.New("failed to retrieve URLs for the user")
	}

	response := mapper.ParseUrlListResponse(urls)

	return response, nil
}

func (service *urlService) DeleteUrl(ctx context.Context, email, shortUrl string) error {
	err := service.urlRepo.Delete(ctx, email, shortUrl)
	if err != nil {
		return errors.New("failed to delete the URL or it does not belong to the user")
	}

	return nil
}
