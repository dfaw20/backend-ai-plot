//go:build wireinject

package main

import (
	"github.com/dfaw20/backend-ai-plot/dependency"
	"github.com/dfaw20/backend-ai-plot/handlers"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/google/wire"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

func initializeDIContainer(db *gorm.DB, oauth2Config oauth2.Config) dependency.DIContainer {
	wire.Build(
		dependency.NewDIContainer,
		handlers.NewAuthHandler,
		handlers.NewUserHandler,
		handlers.NewPlayerHandler,
		handlers.NewCharacterHandler,
		repositories.NewUserRepository,
		repositories.NewPlayerRepository,
		repositories.NewUserTokenRepository,
		repositories.NewCharacterRepository,
	)
	return dependency.DIContainer{}
}
