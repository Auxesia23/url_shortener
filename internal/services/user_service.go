package service

import (
	"context"
	"errors"

	"github.com/Auxesia23/url_shortener/internal/auth"
	"github.com/Auxesia23/url_shortener/internal/models"
	"github.com/Auxesia23/url_shortener/internal/repositories"
)

type UserService interface {
	GoogleLogin(ctx context.Context,code string)(string, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (service *userService) GoogleLogin(ctx context.Context,code string)(string, error){
	//Exchange code for oauth2 token
	token, err := auth.ExchangeToken(code)
	if err != nil {
		return "",errors.New("failed to exchange code for token: " + err.Error())
	}
	
	//Fetching google user info
	googleUser, err := auth.FetchUserInfo(token)
	if err != nil {
		return "",errors.New("failed to fetch user info: " + err.Error())
	}
	
	//User for creating jwt
	var userForJWT models.User 
	
	//Check if user already exist
	userForJWT, err = service.userRepo.Read(ctx, googleUser.Email)
	if err != nil {
		//If not exist, create new user
		newUser := models.User{
			Email: googleUser.Email,
			Name: googleUser.Name,
			Picture: googleUser.Picture,
		}
		
		//Creating new user from google user info
		err := service.userRepo.Create(ctx, newUser)
		if err != nil {
			return "", errors.New("failed to create user: " + err.Error())
		}
		
		//Getting created user data
		userForJWT, err = service.userRepo.Read(ctx, googleUser.Email)
		if err != nil {
			return "", errors.New("failed to retrieve created user: " + err.Error())
		}
	}
	
	//Generate jwt
	jwt, err := auth.GenerateToken(&userForJWT)
	if err != nil {
		return "", errors.New("failed to generate JWT: " + err.Error())
	}
	
	return jwt, nil
}
	


