package handlers

import (
	"errors"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input requests.UserEdit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("user").(models.User)

	if len(input.DisplayName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("表示名が未入力です")})
		return
	}

	user.DisplayName = input.DisplayName

	if err := h.userRepository.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}
