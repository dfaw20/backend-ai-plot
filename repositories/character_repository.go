package repositories

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/jinzhu/gorm"
)

type CharacterRepository struct {
	db *gorm.DB
}

func NewCharacterRepository(db *gorm.DB) CharacterRepository {
	return CharacterRepository{db}
}

func (r *CharacterRepository) GetCharactersByUser(user models.User) ([]models.Character, error) {
	var characters []models.Character
	if err := r.db.Where("user_id = ?", user.ID).Find(&characters).Error; err != nil {
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
