package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler interface{
	GoogleLogin(c *gin.Context)
	GoogleCallback(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}


func (handler *userHandler) GoogleLogin(c *gin.Context){
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	redirectUri := os.Getenv("GOOGLE_REDIRECT_URI")
	
	oauthUri := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/v2/auth?scope=email%%20profile&redirect_uri=%s&response_type=code&client_id=%s&access_type=offline",
		redirectUri,
		clientId,
	)
	c.PureJSON(http.StatusOK, gin.H{"url" : oauthUri})
}

func (handler *userHandler) GoogleCallback(c *gin.Context){
	code := c.Query("code")
	
	jwt, err := handler.userService.GoogleLogin(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error" : err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"acces_token" : jwt})
}
