package handlers

import (
	"errors"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/dfaw20/backend-ai-plot/utils"
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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var input requests.UserEdit
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authUser := c.Value("auth_user").(models.User)

	if len(input.DisplayName) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("表示名が未入力です")})
		return
	}

	authUser.DisplayName = input.DisplayName

	if err := h.userRepository.UpdateUser(&authUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, authUser)
}

func (h *UserHandler) GetUserCharacters(c *gin.Context) {
	userIdStr := c.Param("user_id")
	userId, err := utils.ParseUint(userIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepository.FindByUserID(uint(userId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	characters, err := h.characterRepository.GetCharactersByUser(user)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, characters)
}
