package middlewares

import (
	"context"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	"github.com/jinzhu/gorm"

	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/dfaw20/backend-ai-plot/repositories"
)

func AuthMiddleware(db *gorm.DB, oauth2Config oauth2.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダーからAuthorizationヘッダーを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// "Bearer " のプレフィックスを取り除いてアクセストークンを取得
		accessToken := authHeader[len("Bearer "):]

		isValid, user, _ := isValidAccessToken(db, accessToken, oauth2Config)

		if isValid {
			c.Set("auth_user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}
	}
}

func isValidAccessToken(db *gorm.DB, accessToken string, oauth2Config oauth2.Config) (bool, models.User, error) {
	userTokenRepository := repositories.NewUserTokenRepository(db)
	token, err := userTokenRepository.FindByAccessToken(accessToken)
	if err != nil {
		return false, models.User{}, err
	}

	cxt := context.Background()
	oauth2Service, err := v2.NewService(cxt, option.WithTokenSource(oauth2Config.TokenSource(cxt, &token)))
	if err != nil {
		return false, models.User{}, err
	}

	userInfo, err := oauth2Service.Userinfo.V2.Me.Get().Do()
	if err != nil {
		return false, models.User{}, err
	}

	userRepository := repositories.NewUserRepository(db)
	user, err := userRepository.FindByUserInfo(*userInfo)
	if err != nil {
		return false, models.User{}, err
	}

	if user.ID == 0 {
		return false, models.User{}, err
	}

	return true, user, nil
}
