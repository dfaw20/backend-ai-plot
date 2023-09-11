package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/utils"
	"github.com/gin-gonic/gin"
)

type PlayerHandler struct {
	userRepository      repositories.UserRepository
	characterRepository repositories.CharacterRepository
}

func NewPlayerHandler(
	userRepository repositories.UserRepository,
	characterRepository repositories.CharacterRepository,
) PlayerHandler {
	return PlayerHandler{
		userRepository,
		characterRepository,
	}
}

func (h *PlayerHandler) GetPlayerCharacters(c *gin.Context) {
	playerIdStr := c.Param("player_id")
	playerId, err := utils.ParseUint(playerIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	player, err := h.userRepository.FindByUserID(uint(playerId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	characters, err := h.characterRepository.GetCharactersByUser(player)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, characters)
}
