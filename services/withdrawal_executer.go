package services

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"gorm.io/gorm"
)

type WithdrawalExecuter struct {
	db                  *gorm.DB
	userRepository      repositories.UserRepository
	plotRepository      repositories.PlotRepository
	characterRepository repositories.CharacterRepository
	storyRepository     repositories.StoryRepository
}

func NewWithdrawalExecuter(
	db *gorm.DB,
	userRepository repositories.UserRepository,
	plotRepository repositories.PlotRepository,
	characterRepository repositories.CharacterRepository,
	storyRepository repositories.StoryRepository,
) WithdrawalExecuter {
	return WithdrawalExecuter{
		db:                  db,
		userRepository:      userRepository,
		plotRepository:      plotRepository,
		characterRepository: characterRepository,
		storyRepository:     storyRepository,
	}
}

func (e *WithdrawalExecuter) DoWithdrawal(user models.User) error {

	tx := e.db.Begin()

	if err := e.plotRepository.DeletePlotsByUser(user); err != nil {
		tx.Rollback()
		return err
	}

	if err := e.characterRepository.DeleteCharactersByUser(user); err != nil {
		tx.Rollback()
		return err
	}
	if err := e.storyRepository.DeleteStoriesByUser(user); err != nil {
		tx.Rollback()
		return err
	}
	if err := e.userRepository.DeleteByUserID(user.ID); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
