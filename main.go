package main

import (
	"github.com/dfaw20/backend-ai-plot/database"
	"github.com/dfaw20/backend-ai-plot/migrations"
	"github.com/dfaw20/backend-ai-plot/serve"
)

func main() {
	db := database.ConnectDB()
	migrations.AutoMigrate(db)
	serve.RunServer()
}
