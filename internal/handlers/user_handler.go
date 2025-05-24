package handler

import (
	"github.com/Auxesia23/url_shortener/internal/services"

)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
