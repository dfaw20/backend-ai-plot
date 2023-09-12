package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/dfaw20/backend-ai-plot/utils"
	"github.com/gin-gonic/gin"
)

type PlotHandler struct {
	plotRepo   repositories.PlotRepository
	playerRepo repositories.PlayerRepository
}

func NewPlotHandler(
	plotRepo repositories.PlotRepository,
	playerRepo repositories.PlayerRepository,
) PlotHandler {
	return PlotHandler{plotRepo, playerRepo}
}

func (h *PlotHandler) GetPlotByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := utils.ParseUint(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid plot ID"})
		return
	}

	plot, err := h.plotRepo.GetPlotByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Plot not found"})
		return
	}
	c.JSON(http.StatusOK, plot)
}

func (h *PlotHandler) CreatePlot(c *gin.Context) {
	var input requests.PlotInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := c.Value("auth_user").(models.User)

	var plot = models.Plot{
		UserID:       user.ID,
		Title:        input.Title,
		Description:  input.Description,
		Prompt:       input.Prompt,
		Location:     input.Location,
		Season:       input.Season,
		Genre:        input.Genre,
		OutputFormat: input.OutputFormat,
		ShowWarning:  input.ShowWarning,
	}

	if err := h.plotRepo.CreatePlot(&plot); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, plot)
}

func (h *PlotHandler) GetPlotsRecently(c *gin.Context) {
	plots, err := h.plotRepo.GetPlotsOrderByUpdatedAtDescLimit100()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plots)
}

func (h *PlotHandler) GetPlotsByPlayer(c *gin.Context) {
	playerIdStr := c.Param("player_id")
	playerId, err := utils.ParseUint(playerIdStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	player, err := h.playerRepo.FindByPlayerID(uint(playerId))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	plots, err := h.plotRepo.GetPlotsByPlayer(player)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, plots)
}
