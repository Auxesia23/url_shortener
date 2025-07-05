package handler

import (
	"net/http"

	"github.com/Auxesia23/url_shortener/internal/mapper"
	service "github.com/Auxesia23/url_shortener/internal/services"
	"github.com/gin-gonic/gin"
)

type UrlHandler interface {
	HandleCreateUrl(c *gin.Context)
	HandleGetUrl(c *gin.Context)
	HandleGetUrlByEmail(c *gin.Context)
	HandleDeleteUrl(c *gin.Context)
}

type urlHandler struct {
	urlService      service.UrlService
	analyticService service.AnalyticService
}

func NewUrlHandler(urlService service.UrlService, analyticServive service.AnalyticService) UrlHandler {
	return &urlHandler{
		urlService:      urlService,
		analyticService: analyticServive,
	}
}

// HandleCreateUrl godoc
// @Summary Create a new short URL
// @Description Create a new short URL for the authenticated user
// @Tags urls
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param url body mapper.UrlInput true "URL input"
// @Success 201 {object} mapper.UrlResponse
// @Failure 400 {object} ErrorMessage
// @Failure 401 {object} ErrorMessage
// @Failure 409 {object} ErrorMessage
// @Router /urls/ [post]
// @Security BearerAuth
func (handler *urlHandler) HandleCreateUrl(c *gin.Context) {
	user, ok := c.Get("user")
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorMessage{Message: "Not Authorized"})
		return
	}

	var urlInput mapper.UrlInput
	err := c.Bind(&urlInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorMessage{Message: "Invalid JSON Format"})
		return
	}

	url, err := handler.urlService.CreateShortUrl(c.Request.Context(), urlInput, user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, ErrorMessage{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, url)
}

// HandleGetUrl godoc
// @Summary Get a short URL detail
// @Description Get detail and analytics for a short URL owned by the authenticated user
// @Tags urls
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Short URL ID"
// @Success 200 {object} mapper.UrlAnalyticResponse
// @Failure 401 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage
// @Router /urls/{id} [get]
// @Security BearerAuth
func (handler *urlHandler) HandleGetUrl(c *gin.Context) {
	shortUrl := c.Param("id")
	user := c.MustGet("user")

	url, err := handler.urlService.GetUrl(c.Request.Context(), user.(string), shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
		return
	}

	analytic, err := handler.analyticService.Get(c.Request.Context(), shortUrl)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
		return
	}

	response := mapper.UrlAnalyticResponse{
		Url:      url,
		Analytic: analytic,
	}

	c.JSON(http.StatusOK, response)
}

// HandleGetUrlByEmail godoc
// @Summary Get all short URLs by user
// @Description Get all short URLs created by the authenticated user
// @Tags urls
// @Produce json
// @Param Authorization header string true "Bearer token"
// @Success 200 {array} mapper.UrlListResponse
// @Failure 401 {object} ErrorMessage
// @Failure 500 {object} ErrorMessage
// @Router /urls/ [get]
// @Security BearerAuth
func (handler *urlHandler) HandleGetUrlByEmail(c *gin.Context) {
	user := c.MustGet("user")
	response, err := handler.urlService.GetUrlByEmail(c.Request.Context(), user.(string))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorMessage{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}

// HandleDeleteUrl godoc
// @Summary Delete a short URL
// @Description Delete a short URL owned by the authenticated user
// @Tags urls
// @Param Authorization header string true "Bearer token"
// @Param id path string true "Short URL ID"
// @Success 204
// @Failure 401 {object} ErrorMessage
// @Failure 404 {object} ErrorMessage
// @Router /urls/{id} [delete]
// @Security BearerAuth
func (handler *urlHandler) HandleDeleteUrl(c *gin.Context) {
	user := c.MustGet("user")
	shortUrl := c.Param("id")

	err := handler.urlService.DeleteUrl(c.Request.Context(), user.(string), shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorMessage{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
