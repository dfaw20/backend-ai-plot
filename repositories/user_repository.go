package repositories

import (
	"errors"

	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"

	v2 "google.golang.org/api/oauth2/v2"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{db: db}
}

func (r *UserRepository) CreateUserIfNotExist(userInfo v2.Userinfo) (models.User, error) {

	// バリデーション
	if len(userInfo.Email) == 0 {
		return models.User{}, errors.New("メールアドレスが取得できませんでした。")
	}

	// データベースからユーザ情報を検索
	var user models.User
	err := r.db.Where("email = ?", userInfo.Email).First(&user).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, err
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		// ユーザが存在しない場合、新しいユーザを作成
		newUser := models.User{
			Email:           userInfo.Email,
			DisplayName:     userInfo.Name,
			SensitiveDirect: false,
		}

		r.db.Create(&newUser)

		return newUser, nil
	} else {
		return user, nil
	}
}

func (r *UserRepository) FindByUserInfo(userInfo v2.Userinfo) (models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", userInfo.Email).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if user.ID == 0 {
		return models.User{}, errors.New("ユーザが見つかりません")
	}

	return user, nil
}

func (r *UserRepository) FindByUserID(userID uint) (models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", userID).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if user.ID == 0 {
		return models.User{}, errors.New("ユーザが見つかりません")
	}

	return user, nil
}

func (r *UserRepository) DeleteByUser(user models.User) error {
	return r.db.Delete(&models.User{}, "id = ?", user.ID).Error
}

func (r *UserRepository) UpdateUserEmail(userID uint, email string) (models.User, error) {
	user, err := r.FindByUserID(userID)

	if err != nil {
		return models.User{}, err
	}

	user.Email = email

	if err := r.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUserDisplayName(userID uint, displayName string) (models.User, error) {
	user, err := r.FindByUserID(userID)

	if err != nil {
		return models.User{}, err
	}

	user.DisplayName = displayName

	if err := r.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepository) UpdateUserSensitiveOption(userID uint, sensitiveDirect bool) (models.User, error) {
	user, err := r.FindByUserID(userID)

	if err != nil {
		return models.User{}, err
	}

	user.SensitiveDirect = sensitiveDirect

	if err := r.db.Save(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}
