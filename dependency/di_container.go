package dependency

import "github.com/dfaw20/backend-ai-plot/handlers"

type DIContainer struct {
	AuthHandler      handlers.AuthHandler
	UserHandler      handlers.UserHandler
	PlayerHandler    handlers.PlayerHandler
	CharacterHandler handlers.CharacterHandler
	PlotHandler      handlers.PlotHandler
	TaleHandler      handlers.TaleHandler
	StoryHandler     handlers.StoryHandler
}

func NewDIContainer(
	authHandler handlers.AuthHandler,
	userHandler handlers.UserHandler,
	playerHandler handlers.PlayerHandler,
	CharacterHandler handlers.CharacterHandler,
	plotHandler handlers.PlotHandler,
	TaleHandler handlers.TaleHandler,
	StoryHandler handlers.StoryHandler,
) DIContainer {
	return DIContainer{
		authHandler,
		userHandler,
		playerHandler,
		CharacterHandler,
		plotHandler,
		TaleHandler,
		StoryHandler,
	}
}
