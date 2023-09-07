//go:build wireinject

package main

import (
	"github.com/dfaw20/backend-ai-plot/dependency"
	"github.com/dfaw20/backend-ai-plot/handlers"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"golang.org/x/oauth2"
)

func initializeDIContainer(db *gorm.DB, oauth2Config oauth2.Config) dependency.DIContainer {
	wire.Build(
		dependency.NewDIContainer,
		handlers.NewAuthHandler,
		repositories.NewUserRepository,
		repositories.NewUserTokenRepository,
	)
	return dependency.DIContainer{}
}
