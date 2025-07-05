package handler

import (
	"context"
	"net/http"

	service "github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

// RedirectHandler interface
// @Summary Redirect to original URL
// @Description Redirect short URL to the original URL and save analytics
// @Tags redirect
// @Param id path string true "Short URL ID"
// @Success 307
// @Failure 404 {object} map[string]string
// @Router /{id} [get]
type RedirectHandler interface {
	HandleRedirect(c *gin.Context)
}

type redirectHandler struct {
	redirectService service.RedirectService
	analyticServive service.AnalyticService
}

func NewRedirectHandler(redirectService service.RedirectService, analyticServive service.AnalyticService) RedirectHandler {
	return &redirectHandler{
		redirectService: redirectService,
		analyticServive: analyticServive,
	}
}

func (handler *redirectHandler) HandleRedirect(c *gin.Context) {
	id := c.Param("id")
	ip := c.ClientIP()
	agent := c.Request.UserAgent()

	go handler.analyticServive.Save(context.Background(), id, ip, agent)

	url, err := handler.redirectService.Redirect(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, url.Original)
}
