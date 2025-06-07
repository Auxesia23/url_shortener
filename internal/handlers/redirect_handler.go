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
	analyticServive service.AnalyticService
}

func NewRedirectHandler(redirectService service.RedirectService, analyticServive service.AnalyticService) RedirectHandler{
	return &redirectHandler{
		redirectService: redirectService,
		analyticServive: analyticServive,
	}
}

func (handler *redirectHandler)HandleRedirect(c *gin.Context){
	id := c.Param("id")
	ip := c.ClientIP()
	agent := c.Request.UserAgent()
	
	go handler.analyticServive.Save(context.Background(), id, ip, agent)

	url, err := handler.redirectService.Redirect(c.Request.Context(), id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error" : err.Error()})
		return
	}
	
	c.Redirect(http.StatusTemporaryRedirect, url.Original)
}