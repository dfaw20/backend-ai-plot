package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/gin-gonic/gin"
)

type StoryHandler struct {
	storyRepo repositories.StoryRepository
}

func NewStoryHandler(
	storyRepo repositories.StoryRepository,
) StoryHandler {
	return StoryHandler{storyRepo}
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

	story.Prompt

	c.JSON(http.StatusCreated, story)
}
