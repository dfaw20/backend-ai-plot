package services

import (
	"time"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/dfaw20/backend-ai-plot/repositories"
	"gorm.io/gorm"
)

type WithdrawalExecuter struct {
	db                        *gorm.DB
	userRepository            repositories.UserRepository
	plotRepository            repositories.PlotRepository
	characterRepository       repositories.CharacterRepository
	storyRepository           repositories.StoryRepository
	withdrawalEmailRepository repositories.WithdrawalEmailRepository
}

func NewWithdrawalExecuter(
	db *gorm.DB,
	userRepository repositories.UserRepository,
	plotRepository repositories.PlotRepository,
	characterRepository repositories.CharacterRepository,
	storyRepository repositories.StoryRepository,
	withdrawalEmailRepository repositories.WithdrawalEmailRepository,
) WithdrawalExecuter {
	return WithdrawalExecuter{
		db:                        db,
		userRepository:            userRepository,
		plotRepository:            plotRepository,
		characterRepository:       characterRepository,
		storyRepository:           storyRepository,
		withdrawalEmailRepository: withdrawalEmailRepository,
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
	// 退会Eメールリストに追加する
	if _, err := e.withdrawalEmailRepository.InsertEmail(user.Email); err != nil {
		tx.Rollback()
		return err
	}
	// 削除前にメールアドレスに日時文字列をつける(再登録時に重複エラー回避の為)
	if _, err := e.userRepository.UpdateUserEmail(user.ID, user.Email+time.Now().String()); err != nil {
		tx.Rollback()
		return err
	}
	if err := e.userRepository.DeleteByUser(user); err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
