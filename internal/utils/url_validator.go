package utils

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
)

func ValidateOriginalUrl(inputUrl string) (string, error) {

	parsedUrl, err := url.Parse(inputUrl)
	if err != nil {
		return "", fmt.Errorf("invalid original URL format: %v", err)
	}

	if parsedUrl.Scheme == "" {
		inputUrl = "https://" + inputUrl
		parsedUrl, err = url.Parse(inputUrl)
		if err != nil {
			return "", fmt.Errorf("invalid original URL after adding scheme: %v", err)
		}
	}

	if parsedUrl.Host == "" {
		return "", fmt.Errorf("URL must have a host")
	}

	return parsedUrl.String(), nil
}

func VerifyShortenedUrl(inputUrl string) bool {
	shortenedPattern := `^[a-zA-Z0-9_-]{5,20}$`
	reShortened, err := regexp.Compile(shortenedPattern)
	if err != nil {
		log.Println("Regex pattern error")
	}

	return reShortened.MatchString(inputUrl)
}
