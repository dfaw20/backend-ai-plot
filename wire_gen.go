// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/dependency"
	"github.com/dfaw20/backend-ai-plot/handlers"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/dfaw20/backend-ai-plot/services"
	"golang.org/x/oauth2"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func initializeDIContainer(db *gorm.DB, oauth2Config2 oauth2.Config, config configuration.Config) dependency.DIContainer {
	userRepository := repositories.NewUserRepository(db)
	userTokenRepository := repositories.NewUserTokenRepository(db)
	withdrawalEmailRepository := repositories.NewWithdrawalEmailRepository(db)
	authHandler := handlers.NewAuthHandler(oauth2Config2, userRepository, userTokenRepository, withdrawalEmailRepository)
	characterRepository := repositories.NewCharacterRepository(db)
	userHandler := handlers.NewUserHandler(oauth2Config2, userRepository, userTokenRepository, characterRepository)
	playerRepository := repositories.NewPlayerRepository(userRepository)
	plotRepository := repositories.NewPlotRepository(db)
	playerHandler := handlers.NewPlayerHandler(playerRepository, characterRepository, plotRepository)
	characterHandler := handlers.NewCharacterHandler(characterRepository)
	plotHandler := handlers.NewPlotHandler(plotRepository)
	storyRepository := repositories.NewStoryRepository(db)
	taleHandler := handlers.NewTaleHandler(plotRepository, characterRepository, storyRepository)
	chatGenerator := services.NewChatGenerator(config)
	storyHandler := handlers.NewStoryHandler(storyRepository, chatGenerator)
	withdrawalExecuter := services.NewWithdrawalExecuter(db, userRepository, plotRepository, characterRepository, storyRepository, withdrawalEmailRepository)
	withdrawalHandler := handlers.NewWithdrawalHandler(withdrawalExecuter, withdrawalEmailRepository)
	diContainer := dependency.NewDIContainer(authHandler, userHandler, playerHandler, characterHandler, plotHandler, taleHandler, storyHandler, withdrawalHandler)
	return diContainer
}
