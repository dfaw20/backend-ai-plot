package dependency

import "github.com/dfaw20/backend-ai-plot/handlers"

type DIContainer struct {
	AuthHandler      handlers.AuthHandler
	UserHandler      handlers.UserHandler
	CharacterHandler handlers.CharacterHandler
}

func NewDIContainer(
	authHandler handlers.AuthHandler,
	userHandler handlers.UserHandler,
	CharacterHandler handlers.CharacterHandler,
) DIContainer {
	return DIContainer{
		authHandler,
		userHandler,
		CharacterHandler,
	}
}
