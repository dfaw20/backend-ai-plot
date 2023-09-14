package main

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/database"
	"github.com/dfaw20/backend-ai-plot/middlewares"
	"github.com/dfaw20/backend-ai-plot/migrations"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm/logger"
)

var (
	oauth2Config oauth2.Config
)

func main() {
	config := configuration.LoadConfig()

	db := database.ConnectDB(config.Postgres)
	db.Logger = db.Logger.LogMode(logger.Info)

	migrations.AutoMigrate(db, config)

	// OAuth2設定を構築
	oauth2Config = oauth2.Config{
		ClientID:     config.Google.ClientID,
		ClientSecret: config.Google.ClientSecret,
		RedirectURL:  config.Google.RedirectURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	r := gin.Default()

	// CORS 対応
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = config.Frontend.Origins
	corsConfig.AllowMethods = []string{
		"POST", "GET", "OPTIONS",
	}
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
	}
	r.Use(cors.New(corsConfig))

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware(db, oauth2Config))

	di := initializeDIContainer(db, oauth2Config, config)

	r.GET("/auth/google", di.AuthHandler.GetOAuthURL)
	r.GET("/auth/google/callback", di.AuthHandler.GetAuthGoogleCallback)

	authorized.GET("/user", di.UserHandler.GetUserInfo)
	authorized.POST("/user/update/display_name", di.UserHandler.UpdateUserDisplayName)
	authorized.POST("/user/update/sensitive_option", di.UserHandler.UpdateUserSensitiveOption)

	r.GET("/players/:player_id", di.PlayerHandler.GetPlayer)
	r.GET("/players/:player_id/characters", di.PlayerHandler.GetPlayerCharacters)
	r.GET("/players/:player_id/plots", di.PlayerHandler.GetPlayerPlots)

	r.GET("/characters/:id", di.CharacterHandler.GetCharacterByID)
	authorized.POST("/characters/create", di.CharacterHandler.CreateCharacter)

	r.GET("/plots/:id", di.PlotHandler.GetPlotByID)
	authorized.POST("/plots/create", di.PlotHandler.CreatePlot)
	r.GET("/plots/recent", di.PlotHandler.GetPlotsRecently)

	authorized.POST("/tale/create", di.TaleHandler.CreateTale)
	authorized.POST("/story/create", di.StoryHandler.GenerateChat)

	r.Run(":8080")
}
