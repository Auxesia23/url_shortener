package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
}

func NewHealthCheck() *HealthCheck {
	return &HealthCheck{}
}

func (hc *HealthCheck) Check(c *gin.Context){
	c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "Health check successful"})
}