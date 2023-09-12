package repositories

import (
	"github.com/dfaw20/backend-ai-plot/entities"
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
	return CharacterRepository{db}
}

func (r *CharacterRepository) GetCharactersByPlayer(player entities.Player) ([]models.Character, error) {
	var characters []models.Character
	if err := r.db.Where("user_id = ?", player.ID).Find(&characters).Error; err != nil {
		return nil, err
	}
	return characters, nil
}

func (r *CharacterRepository) GetCharacterByID(id uint) (*models.Character, error) {
	var character models.Character
	if err := r.db.First(&character, id).Error; err != nil {
		return nil, err
	}
	return &character, nil
}

func (r *CharacterRepository) CreateCharacter(character *models.Character) error {
	if err := r.db.Create(character).Error; err != nil {
		return err
	}
	return nil
}
