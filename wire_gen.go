// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dfaw20/backend-ai-plot/dependency"
	"github.com/dfaw20/backend-ai-plot/handlers"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initializeDIContainer(db *gorm.DB, oauth2Config2 oauth2.Config) dependency.DIContainer {
	userRepository := repositories.NewUserRepository(db)
	userTokenRepository := repositories.NewUserTokenRepository(db)
	authHandler := handlers.NewAuthHandler(oauth2Config2, userRepository, userTokenRepository)
	characterRepository := repositories.NewCharacterRepository(db)
	userHandler := handlers.NewUserHandler(oauth2Config2, userRepository, userTokenRepository, characterRepository)
	playerRepository := repositories.NewPlayerRepository(userRepository)
	plotRepository := repositories.NewPlotRepository(db)
	playerHandler := handlers.NewPlayerHandler(playerRepository, characterRepository, plotRepository)
	characterHandler := handlers.NewCharacterHandler(characterRepository)
	plotHandler := handlers.NewPlotHandler(plotRepository)
	diContainer := dependency.NewDIContainer(authHandler, userHandler, playerHandler, characterHandler, plotHandler)
	return diContainer
}
