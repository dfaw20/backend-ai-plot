package dependency

import "github.com/dfaw20/backend-ai-plot/handlers"

type DIContainer struct {
	HomeHandler       handlers.HomeHandler
	AuthHandler       handlers.AuthHandler
	UserHandler       handlers.UserHandler
	PlayerHandler     handlers.PlayerHandler
	CharacterHandler  handlers.CharacterHandler
	PlotHandler       handlers.PlotHandler
	TaleHandler       handlers.TaleHandler
	StoryHandler      handlers.StoryHandler
	WithdrawalHandler handlers.WithdrawalHandler
}

func NewDIContainer(
	homeHandler handlers.HomeHandler,
	authHandler handlers.AuthHandler,
	userHandler handlers.UserHandler,
	playerHandler handlers.PlayerHandler,
	CharacterHandler handlers.CharacterHandler,
	plotHandler handlers.PlotHandler,
	TaleHandler handlers.TaleHandler,
	StoryHandler handlers.StoryHandler,
	WithdrawalHandler handlers.WithdrawalHandler,
) DIContainer {
	return DIContainer{
		homeHandler,
		authHandler,
		userHandler,
		playerHandler,
		CharacterHandler,
		plotHandler,
		TaleHandler,
		StoryHandler,
		WithdrawalHandler,
	}
}
