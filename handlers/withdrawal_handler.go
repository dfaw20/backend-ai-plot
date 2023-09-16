package handlers

import (
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/requests"
	"github.com/dfaw20/backend-ai-plot/services"
	"github.com/gin-gonic/gin"
)

type WithdrawalHandler struct {
	withdrawalExecuter        services.WithdrawalExecuter
	withdrawalEmailRepository repositories.WithdrawalEmailRepository
}

func NewWithdrawalHandler(
	withdrawalExecuter services.WithdrawalExecuter,
	withdrawalEmailRepository repositories.WithdrawalEmailRepository,
) WithdrawalHandler {
	return WithdrawalHandler{
		withdrawalExecuter,
		withdrawalEmailRepository,
	}
}

func (h *WithdrawalHandler) DoWithdrawal(c *gin.Context) {
	authUser := c.Value("auth_user").(models.User)

	if err := h.withdrawalExecuter.DoWithdrawal(authUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

func (h *WithdrawalHandler) enableReRegister(c *gin.Context) {
	var input requests.ReRegisterEmail
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.withdrawalEmailRepository.DeleteByEmail(input.WithdrawalEmail).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "退会の解除に失敗しました"})
		return
	}

	c.JSON(http.StatusAccepted, nil)
}
