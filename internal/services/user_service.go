package service

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"github.com/Auxesia23/url_shortener/internal/repositories"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) UserGoogleLogin(ctx context.Context, user models.GoogleUser) (string, error) {
	_, err := s.userRepo.Read(ctx, user.Email)
	if err == nil {
		//Make jwt if user already exist
	}
	
	newUser := models.User{
		Email: user.Email,
		Name: user.GivenName,
		Picture: user.Picture,
	}
	
	err = s.userRepo.Create(ctx, newUser)
	if err != nil {
		return "", err
	}
	
	
	return "", nil
}


