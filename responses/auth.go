package responses

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"golang.org/x/oauth2"
)

type TokenResult struct {
	Token           oauth2.Token
	User            models.User
	IsWithdrawal    bool
	WithdrawalEmail string
}
