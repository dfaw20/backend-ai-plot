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
)

var (
	oauth2Config oauth2.Config
)

func diWire() {

}

func main() {
	config := configuration.LoadConfig()

	db := database.ConnectDB(config.Postgres)
	defer database.CloseDB(db)
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
	corsConfig.AllowOrigins = []string{
		config.Frontend.Origin,
	}
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

	di := initializeDIContainer(db, oauth2Config)

	r.GET("/auth/google", di.AuthHandler.GetOAuthURL)
	r.GET("/auth/google/callback", di.AuthHandler.GetAuthGoogleCallback)
	authorized.GET("/user_info", di.AuthHandler.GetUserInfo)
	authorized.GET("/chracters", di.CharacterHandler.GetCharacters)
	authorized.GET("/chracters/:id", di.CharacterHandler.GetCharacterByID)
	authorized.POST("/chracters", di.CharacterHandler.CreateCharacter)

	r.Run(":8080")
}
