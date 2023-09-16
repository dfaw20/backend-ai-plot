package repositories

import (
	"log"

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

func (r *StoryRepository) DeleteStoriesByUser(user models.User) error {
	if err := r.db.Delete(&models.Story{}, "user_id = ?", user.ID).Error; err != nil {
		return err
	}
	return nil
}

func (r *StoryRepository) CreateStory(story *models.Story) error {
	if err := r.db.Create(story).Error; err != nil {
		log.Print(story)
		return err
	}
	return nil
}

func (r *StoryRepository) UpdateText(storyID uint, text string) (models.Story, error) {
	story, err := r.GetStoryByID(storyID)

	if err != nil {
		return models.Story{}, err
	}

	story.Text = text

	if err := r.db.Save(&story).Error; err != nil {
		return models.Story{}, err
	}

	return *story, nil
}
