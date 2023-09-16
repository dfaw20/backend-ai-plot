package repositories

import (
	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"
)

type WithdrawalEmailRepository struct {
	db *gorm.DB
}

func NewWithdrawalEmailRepository(db *gorm.DB) WithdrawalEmailRepository {
	return WithdrawalEmailRepository{db: db}
}

func (r *WithdrawalEmailRepository) DeleteByEmail(email string) error {
	return r.db.Delete(&models.WithdrawalEmail{}, "email = ?", email).Error
}

func (r *WithdrawalEmailRepository) InsertEmail(email string) (models.WithdrawalEmail, error) {
	withdrawalEmail := models.WithdrawalEmail{
		Email: email,
	}

	err := r.db.Create(&withdrawalEmail).Error

	return withdrawalEmail, err
}

func (r *WithdrawalEmailRepository) FindByEmail(email string) ([]models.WithdrawalEmail, error) {
	var emails []models.WithdrawalEmail
	err := r.db.
		Where("email = ?", email).
		Find(&emails).Error

	return emails, err
}
