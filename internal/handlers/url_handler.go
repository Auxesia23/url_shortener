package handler

import (
	"net/http"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	service "github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

type UrlHandler interface{
	HandleCreateUrl(c *gin.Context)
	HandleGetUrl(c *gin.Context)
	HandleGetUrlByEmail(c *gin.Context)
	HandleDeleteUrl(c *gin.Context)
}

type urlHandler struct{
	urlService service.UrlService
	analyticService service.AnalyticService
}

func NewUrlHandler(urlService service.UrlService, analyticServive service.AnalyticService) UrlHandler{
	return &urlHandler{
		urlService: urlService,
		analyticService: analyticServive,
	}
}

func(handler *urlHandler) HandleCreateUrl(c *gin.Context){
	user, ok := c.Get("user")
	if !ok{
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error":"Not Authorized"})
		return
	}
	
	var urlInput mapper.UrlInput
	err := c.Bind(&urlInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error":"Invalid JSON format"})
		return
	}
	
	url, err := handler.urlService.CreateShortUrl(c.Request.Context(), urlInput, user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error":err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, url)
}

func(handler *urlHandler) HandleGetUrl(c *gin.Context){
	shortUrl := c.Param("id")
	user := c.MustGet("user")
	
	url, err := handler.urlService.GetUrl(c.Request.Context(), user.(string), shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}
	
	analytic, err := handler.analyticService.Get(c.Request.Context(), shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}
	
	
	
	c.JSON(http.StatusOK, gin.H{"url":url,"analytic":analytic})
}

func (handler *urlHandler) HandleGetUrlByEmail(c *gin.Context){
	user := c.MustGet("user")
	response, err := handler.urlService.GetUrlByEmail(c.Request.Context(), user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler *urlHandler)HandleDeleteUrl(c *gin.Context){
	user := c.MustGet("user")
	shortUrl := c.Param("id")
	
	err := handler.urlService.DeleteUrl(c.Request.Context(), user.(string), shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}