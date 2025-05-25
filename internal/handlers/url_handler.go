package handler

import (
	"context"
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
}

func NewUrlHandler(urlService service.UrlService) UrlHandler{
	return &urlHandler{
		urlService: urlService,
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
	
	url, err := handler.urlService.CreateShortUrl(context.Background(), urlInput, user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error":"The requested short url has already in use"})
		return
	}
	
	c.JSON(http.StatusCreated, url)
}

func(handler *urlHandler) HandleGetUrl(c *gin.Context){
	shortUrl := c.Param("id")
	
	url, err := handler.urlService.GetUrl(context.Background(), shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error":"Short url not found"})
		return
	}
	
	c.JSON(http.StatusOK, url)
}

func (handler *urlHandler) HandleGetUrlByEmail(c *gin.Context){
	user := c.MustGet("user")
	response, err := handler.urlService.GetUrlByEmail(context.Background(), user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

func (handler *urlHandler)HandleDeleteUrl(c *gin.Context){
	user := c.MustGet("user")
	shortUrl := c.Param("id")
	
	err := handler.urlService.DeleteUrl(context.Background(), user.(string), shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error":err.Error()})
		return
	}
	
	c.Status(http.StatusNoContent)
}