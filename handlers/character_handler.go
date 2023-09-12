package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/dfaw20/backend-ai-plot/utils"
	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	characterRepo repositories.CharacterRepository
}

func NewCharacterHandler(characterRepo repositories.CharacterRepository) CharacterHandler {
	return CharacterHandler{characterRepo}
}

func (h *CharacterHandler) GetCharacterByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ParseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid character ID"})
		return
	}

	character, err := h.characterRepo.GetCharacterByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
		return
	}
	c.JSON(http.StatusOK, character)
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	var input requests.CharacterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("auth_user").(models.User)

	gender, err := models.ChoiceGender(input.Gender)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var character = models.Character{
		UserID:      user.ID,
		Name:        input.Name,
		Nickname:    input.Nickname,
		Gender:      string(gender),
		Outfit:      input.Outfit,
		Hairstyle:   input.Hairstyle,
		Personality: input.Personality,
		Tone:        input.Tone,
		Profile:     input.Profile,
	}

	if err := h.characterRepo.CreateCharacter(&character); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, character)
}
