package service

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"github.com/Auxesia23/url_shortener/internal/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context,user models.User) error {
	return s.userRepo.Create(ctx, user)
}
