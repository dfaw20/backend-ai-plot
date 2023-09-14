package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/dfaw20/backend-ai-plot/services"
	"github.com/dfaw20/backend-ai-plot/utils"
	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	storyRepo     repositories.StoryRepository
	chatGenerator services.ChatGenerator
}

func NewStoryHandler(
	storyRepo repositories.StoryRepository,
	chatGenerator services.ChatGenerator,
) StoryHandler {
	return StoryHandler{storyRepo, chatGenerator}
}

func (h *StoryHandler) GetStory(c *gin.Context) {
	storyIdStr := c.Param("story_id")
	storyId, err := utils.ParseUint(storyIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	story, err := h.storyRepo.GetStoryByID(uint(storyId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, story)
}

func (h *StoryHandler) GenerateChat(c *gin.Context) {
	var input requests.StoryGenerateReq
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	story, err := h.storyRepo.GetStoryByID(input.StoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ストーリーが見つかりません"})
		return
	}

	if len(story.Text) > 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ストーリーが生成済みです"})
		return
	}

	novelText, err := h.chatGenerator.Generate(story.Prompt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ストーリーの生成に失敗しました"})
		return
	}

	savedStory, err := h.storyRepo.UpdateText(story.ID, novelText)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ストーリーの保存に失敗しました"})
		return
	}

	c.JSON(http.StatusCreated, savedStory)
}
