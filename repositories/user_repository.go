package repositories

import (
	"errors"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/jinzhu/gorm"

	v2 "google.golang.org/api/oauth2/v2"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateOrSyncUser(userInfo v2.Userinfo) (models.User, error) {

	// バリデーション
	if len(userInfo.Email) == 0 {
		return models.User{}, errors.New("メールアドレスが取得できませんでした。")
	}

	// データベースからユーザ情報を検索
	var user models.User
	result := r.db.Where("email = ?", userInfo.Email).First(&user)

	if result.Error != nil {
		return models.User{}, result.Error
	}

	if user.ID == 0 {
		// ユーザが存在しない場合、新しいユーザを作成
		newUser := models.User{
			Email:       userInfo.Email,
			DisplayName: userInfo.Name,
		}

		r.db.Create(&newUser)

		return newUser, nil
	} else {
		// ユーザが存在する場合、ユーザ情報を更新
		r.db.Model(&user).Updates(models.User{
			DisplayName: userInfo.Name,
		})

		return user, nil
	}
}
