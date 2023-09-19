package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HomeHandler struct {
}

func NewHomeHandler() HomeHandler {
	return HomeHandler{}
}

func (h *HomeHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"HealthCheckResult": "OK"})
}
