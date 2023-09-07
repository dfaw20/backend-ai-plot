package dependency

import "github.com/dfaw20/backend-ai-plot/handlers"

type DIContainer struct {
	AuthHandler      handlers.AuthHandler
	CharacterHandler handlers.CharacterHandler
}

func NewDIContainer(
	authHandler handlers.AuthHandler,
	CharacterHandler handlers.CharacterHandler,
) DIContainer {
	return DIContainer{
		authHandler,
		CharacterHandler,
	}
}
