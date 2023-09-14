package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/gin-gonic/gin"
)

type TaleHandler struct {
	plotRepo      repositories.PlotRepository
	characterRepo repositories.CharacterRepository
}

func NewTaleHandler(
	plotRepo repositories.PlotRepository,
	characterRepo repositories.CharacterRepository,
) TaleHandler {
	return TaleHandler{plotRepo, characterRepo}
}

func (h *TaleHandler) createTale(c *gin.Context) {
	var input requests.TaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("auth_user").(models.User)
	targetCharacter, err := h.characterRepo.GetCharacterByID(input.TargetCharacterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "対象キャラクターが見つかりません"})
		return
	}

	heroCharacter, err := h.characterRepo.GetCharacterByID(input.HeroCharacterID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "主人公キャラクターが見つかりません"})
		return
	}

	plot, err := h.plotRepo.GetPlotByID(input.PlotID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "物語の舞台が見つかりません"})
		return
	}

	if err := h.plotRepo.CreatePlot(&plot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, plot)
}
