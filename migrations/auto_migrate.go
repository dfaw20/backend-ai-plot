package migrations

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, config configuration.Config) {
	db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
		&models.Character{},
		&models.Plot{},
		&models.Story{},
		&models.StoryCharacter{},
	)
}
