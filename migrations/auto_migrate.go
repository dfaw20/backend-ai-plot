package migrations

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB, config configuration.Config) {
	db.AutoMigrate(&models.User{})
}
