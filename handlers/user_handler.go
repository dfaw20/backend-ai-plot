package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
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

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	user := c.Value("user").(models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
