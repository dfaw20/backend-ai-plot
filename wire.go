//go:build wireinject

package main

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/dependency"
	"github.com/dfaw20/backend-ai-plot/handlers"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/services"
	"github.com/google/wire"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func initializeDIContainer(
	db *gorm.DB,
	oauth2Config oauth2.Config,
	config configuration.Config,
) dependency.DIContainer {
	wire.Build(
		dependency.NewDIContainer,
		handlers.NewAuthHandler,
		handlers.NewUserHandler,
		handlers.NewPlayerHandler,
		handlers.NewCharacterHandler,
		handlers.NewPlotHandler,
		handlers.NewTaleHandler,
		handlers.NewStoryHandler,
		repositories.NewUserRepository,
		repositories.NewPlayerRepository,
		repositories.NewUserTokenRepository,
		repositories.NewCharacterRepository,
		repositories.NewPlotRepository,
		repositories.NewStoryRepository,
		services.NewChatGenerator,
		services.NewWithdrawalExecuter,
	)
	return dependency.DIContainer{}
}
