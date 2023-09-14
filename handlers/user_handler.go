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
	characterRepository repositories.CharacterRepository
}

func NewUserHandler(
	oauth2Config oauth2.Config,
	userRepository repositories.UserRepository,
	userTokenRepository repositories.UserTokenRepository,
	characterRepository repositories.CharacterRepository,
) UserHandler {
	return UserHandler{
		oauth2Config,
		userRepository,
		userTokenRepository,
		characterRepository,
	}
}

func (h *UserHandler) GetUserInfo(c *gin.Context) {
	authUser := c.Value("auth_user").(models.User)

	c.JSON(http.StatusOK, gin.H{
		"user": authUser,
	})
}

func (h *UserHandler) UpdateUserDisplayName(c *gin.Context) {
	var input requests.UserDisplayNameEdit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authUser := c.Value("auth_user").(models.User)

	if len(input.GetTrimDisplayName()) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("表示名が未入力です")})
		return
	}

	if err := h.userRepository.
		UpdateUserDisplayName(authUser.ID, input.GetTrimDisplayName()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, authUser)
}

func (h *UserHandler) UpdateUserSensitiveOption(c *gin.Context) {
	var input requests.UserDisplayNameEdit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authUser := c.Value("auth_user").(models.User)

	if len(input.GetTrimDisplayName()) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("表示名が未入力です")})
		return
	}

	if err := h.userRepository.
		UpdateUserDisplayName(authUser.ID, input.GetTrimDisplayName()); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, authUser)
}
