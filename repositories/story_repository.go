package repositories

import (
	"github.com/dfaw20/backend-ai-plot/entities"
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

type StoryRepository struct {
	db *gorm.DB
}

func NewStoryRepository(db *gorm.DB) StoryRepository {
	return StoryRepository{db}
}

func (r *StoryRepository) GetStoriesByPlayer(player entities.Player) ([]models.Story, error) {
	var stories []models.Story
	if err := r.db.Where("user_id = ?", player.ID).Find(&stories).Error; err != nil {
		return nil, err
	}
	return stories, nil
}

func (r *StoryRepository) GetStoryByID(id uint) (*models.Story, error) {
	var story models.Story
	if err := r.db.First(&story, id).Error; err != nil {
		return nil, err
	}
	return &story, nil
}

func (r *StoryRepository) CreateStory(story *models.Story) error {
	if err := r.db.Create(story).Error; err != nil {
		return err
	}
	return nil
}
