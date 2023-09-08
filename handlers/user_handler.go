package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type UserHandler struct {
	oauth2Config        oauth2.Config
	userRepository      repositories.UserRepository
	userTokenRepository repositories.UserTokenRepository
}

func NewUserHandler() UserHandler {
	return UserHandler{}
}

func (h *AuthHandler) GetUserInfo(c *gin.Context) {
	user, exists := c.Get("user")

	if exists {
		c.JSON(http.StatusOK, gin.H{
			"user": user,
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
}
