package dependency

import "github.com/dfaw20/backend-ai-plot/handlers"

type DIContainer struct {
	AuthHandler handlers.AuthHandler
}

func NewDIContainer(authHandler handlers.AuthHandler) DIContainer {
	return DIContainer{authHandler}
}
