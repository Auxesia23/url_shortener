package mapper

import (
	"fmt"
	"os"

	"github.com/Auxesia23/url_shortener/internal/models"
)


type UrlResponse struct{
	Original string `json:"original"`
	Shortened string `json:"shortened"`
}

type UrlInput struct{
	Original string `json:"original"`
	Shortened string `json:"shortened"`
}

type UrlListResponse struct{
	Urls []UrlResponse `json:"urls"`
}

func ParseResponse(url models.Url) UrlResponse{
	baseUrl := os.Getenv("BASE_URL")
	return UrlResponse{
		Original: url.Original,
		Shortened: fmt.Sprintf(baseUrl+"/%s", url.Shortened),
	}
}

func ParseInput(url UrlInput, email string) models.Url{
	return models.Url{
		Original: url.Original,
		Shortened: url.Shortened,
		UserEmail: email,
	}
}

func ParseListResponse(urls []models.Url)UrlListResponse{
	baseUrl := os.Getenv("BASE_URL")
	var urlList []UrlResponse
	for _,url := range(urls){
		urlList = append(urlList, UrlResponse{
			Original: url.Original,
			Shortened: fmt.Sprintf(baseUrl+"/%s", url.Shortened),
		})
	}
	return UrlListResponse{
		Urls: urlList,
	}
}



