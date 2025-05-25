package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"encoding/json"

	"github.com/Auxesia23/url_shortener/internal/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var GoogleOauthConfig *oauth2.Config

func InitOauth(){
	GoogleOauthConfig = &oauth2.Config{
		ClientID:    os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret:os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}
	
	fmt.Println("Google Oauth Initialized")
}

func ExchangeToken(code string)(*oauth2.Token, error){
	if GoogleOauthConfig == nil {
		return nil, errors.New("Google Oauth config is not initialized")
	}
	
	token, err := GoogleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, errors.New("Error exchanging code")
	}
	
	return token, nil
}

func FetchUserInfo(token *oauth2.Token)(*models.GoogleUser, error) {
	req, err := http.NewRequest("GET", "https://www.googleapis.com/oauth2/v2/userinfo", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization","Bearer"+ token.AccessToken)
	
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get user info: status %d", resp.StatusCode)
	}

	var userInfo models.GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}