package handler

import (
	"context"
	"fmt"
	"net/http"
	"os"

	service "github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	HandleGoogleLogin(c *gin.Context)
	HandleGoogleCallback(c *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

// HandleGoogleLogin godoc
// @Summary Google OAuth login
// @Description Get Google OAuth login URL
// @Tags auth
// @Produce json
// @Success 200 {object} GoogleUrlResponse
// @Router /auth/google [get]
func (handler *userHandler) HandleGoogleLogin(c *gin.Context) {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	redirectUri := os.Getenv("GOOGLE_REDIRECT_URI")

	oauthUri := fmt.Sprintf(
		"https://accounts.google.com/o/oauth2/v2/auth?scope=email%%20profile&redirect_uri=%s&response_type=code&client_id=%s&access_type=offline",
		redirectUri,
		clientId,
	)
	c.PureJSON(http.StatusOK, GoogleUrlResponse{Url: oauthUri})
}

// HandleGoogleCallback godoc
// @Summary Google OAuth callback
// @Description Handle Google OAuth callback and return JWT token
// @Tags auth
// @Produce json
// @Param code query string true "OAuth code"
// @Success 200 {object} TokenResponse
// @Failure 401 {object} ErrorMessage
// @Router /auth/google/callback [get]
func (handler *userHandler) HandleGoogleCallback(c *gin.Context) {
	code := c.Query("code")

	jwt, err := handler.userService.GoogleLogin(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorMessage{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, TokenResponse{AccesToken: jwt})
}
