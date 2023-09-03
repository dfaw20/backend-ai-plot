package migrations

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}
