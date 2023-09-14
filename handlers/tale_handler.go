package handlers

import (
	"log"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/entities"
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/gin-gonic/gin"
)

type TaleHandler struct {
	plotRepo      repositories.PlotRepository
	characterRepo repositories.CharacterRepository
	storyRepo     repositories.StoryRepository
}

func NewTaleHandler(
	plotRepo repositories.PlotRepository,
	characterRepo repositories.CharacterRepository,
	storyRepo repositories.StoryRepository,
) TaleHandler {
	return TaleHandler{plotRepo, characterRepo, storyRepo}
}

func (h *TaleHandler) CreateTale(c *gin.Context) {
	user := c.Value("auth_user").(models.User)

	var input requests.TaleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	talePrompt := entities.NewTalePrompt(
		*targetCharacter, *heroCharacter, *plot,
	)

	fullPrompt := talePrompt.BuildFullPrompt()

	story := models.Story{
		UserID: user.ID,
		PlotID: plot.ID,
		Prompt: fullPrompt,
		Text:   "",
	}
	err = h.storyRepo.CreateCharacter(&story)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ストーリーの保存に失敗しました"})
		return
	}

	log.Println(story)

	c.JSON(http.StatusCreated, story)
}
