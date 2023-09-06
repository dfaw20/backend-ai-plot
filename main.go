package main

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/database"
	"github.com/dfaw20/backend-ai-plot/migrations"
	"github.com/dfaw20/backend-ai-plot/serve"
)

func main() {
	config := configuration.LoadConfig()
	db := database.ConnectDB(config.Postgres)
	defer database.CloseDB(db)
	migrations.AutoMigrate(db, config)
	serve.RunServer(db, config.Google, config.Frontend)
}
