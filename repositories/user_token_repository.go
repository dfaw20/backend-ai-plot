package repositories

import (
	"errors"

	"github.com/dfaw20/backend-ai-plot/models"
	"gorm.io/gorm"

	"golang.org/x/oauth2"
)

type UserTokenRepository struct {
	db *gorm.DB
}

func NewUserTokenRepository(db *gorm.DB) UserTokenRepository {
	return UserTokenRepository{db: db}
}

func (r *UserTokenRepository) FindByAccessToken(accessToken string) (oauth2.Token, error) {
	var userToken models.UserToken
	result := r.db.Where("access_token = ?", accessToken).First(&userToken)

	if result.Error != nil {
		return oauth2.Token{}, result.Error
	}

	if userToken.ID == 0 {
		return oauth2.Token{}, errors.New("トークンが見つかりません")
	}

	return oauth2.Token{
		AccessToken:  userToken.AccessToken,
		TokenType:    userToken.TokenType,
		RefreshToken: userToken.RefreshToken,
		Expiry:       userToken.Expiry,
	}, nil
}

func (r *UserTokenRepository) StoreToken(token oauth2.Token) {

	// 新しいトークンを保存
	newToken := models.UserToken{
		AccessToken:  token.AccessToken,
		TokenType:    token.TokenType,
		RefreshToken: token.RefreshToken,
		Expiry:       token.Expiry,
	}
	r.db.Create(&newToken)
}
