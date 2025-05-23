package handler

import (
	"context"

	"github.com/Auxesia23/url_shortener/internal/models"
	"github.com/Auxesia23/url_shortener/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	err = h.userService.CreateUser(context.Background(), user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(201, gin.H{"message": "User created successfully"})
}
