package handler

import (
	"context"
	"net/http"

	service "github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

type RedirectHandler interface{
	HandleRedirect(c *gin.Context)
}

type redirectHandler struct{
	redirectService service.RedirectService
}

func NewRedirectHandler(redirectService service.RedirectService) RedirectHandler{
	return &redirectHandler{
		redirectService: redirectService,
	}
}

func (handler *redirectHandler)HandleRedirect(c *gin.Context){
	id := c.Param("id")
	url, err := handler.redirectService.Redirect(context.Background(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error" : "Short url not found"})
		return
	}
	
	c.Redirect(http.StatusPermanentRedirect, url.Original)
}