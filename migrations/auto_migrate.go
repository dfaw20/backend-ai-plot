package migrations

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB, config configuration.Config) {
	db.Exec(`
		DO $$ 
		BEGIN 
			IF EXISTS (SELECT constraint_name 
						FROM information_schema.table_constraints 
						WHERE table_name = 'user_tokens' 
						AND constraint_type = 'UNIQUE' 
						AND constraint_name = 'user_tokens_refresh_token_key') 
			THEN 
				ALTER TABLE user_tokens DROP CONSTRAINT user_tokens_refresh_token_key; 
			END IF; 
		END $$;
	`)

	db.AutoMigrate(
		&models.User{},
		&models.UserToken{},
		&models.Character{},
		&models.Plot{},
		&models.Story{},
		&models.StoryCharacter{},
	)
}
