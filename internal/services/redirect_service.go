package service

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	repository "github.com/Auxesia23/url_shortener/internal/repositories"
)

type RedirectService interface{
	Redirect(ctx context.Context,shortUrl string)(mapper.UrlResponse, error)
}

type redirectService struct{
	urlRepo repository.UrlRepository
}

func NewRedirectService (urlRepository repository.UrlRepository) RedirectService{
	return &redirectService{
		urlRepo: urlRepository,
	}
}

func (service *redirectService) Redirect(ctx context.Context,shortUrl string)(mapper.UrlResponse, error){
	url,err := service.urlRepo.Read(ctx, shortUrl)
	if err != nil {
		return mapper.UrlResponse{}, err
	}
	
	response := mapper.ParseResponse(url)
	
	return response, nil
}