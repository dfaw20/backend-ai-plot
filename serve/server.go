package serve

import (
	"context"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/jinzhu/gorm"

	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/dfaw20/backend-ai-plot/repositories"
)

var (
	oauth2Config oauth2.Config
)

func RunServer(db *gorm.DB) {
	config := configuration.LoadConfig()

	// OAuth2設定を構築
	oauth2Config = oauth2.Config{
		ClientID:     config.Google.ClientID,
		ClientSecret: config.Google.ClientSecret,
		RedirectURL:  config.Google.RedirectURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	r := gin.Default()

	r.GET("/auth/google", func(c *gin.Context) {
		url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)
		c.Redirect(http.StatusFound, url)
	})

	r.GET("/auth/google/callback", func(c *gin.Context) {
		code := c.DefaultQuery("code", "")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
			return
		}

		// OAuth2トークンを取得し、ユーザー情報を取得できます。
		token, err := oauth2Config.Exchange(c, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ユーザー情報の取得やデータベースへの保存などをここで行います。

		// アクセストークンを使用してユーザ情報を取得
		cxt := context.Background()
		oauth2Service, err := v2.NewService(cxt, option.WithTokenSource(oauth2Config.TokenSource(cxt, token)))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userInfo, err := oauth2Service.Userinfo.V2.Me.Get().Do()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ユーザ情報の作成or同期
		userRepository := repositories.NewUserRepository(db)
		user, err := userRepository.CreateOrSyncUser(*userInfo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ユーザー情報の表示
		c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
	})

	r.Run(":8080")
}
