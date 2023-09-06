package serve

import (
	"context"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/dfaw20/backend-ai-plot/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/jinzhu/gorm"

	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"

	"github.com/dfaw20/backend-ai-plot/repositories"

	"github.com/gin-contrib/cors"
)

var (
	oauth2Config oauth2.Config
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
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

		isValid, user, _ := isValidAccessToken(db, accessToken)

		if isValid {
			c.Set("user", user)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid access token"})
			c.Abort()
			return
		}
	}
}

func RunServer(db *gorm.DB,
	googleConfig configuration.GoogleConfig,
	frontendConfig configuration.FrontendConfig) {
	// OAuth2設定を構築
	oauth2Config = oauth2.Config{
		ClientID:     googleConfig.ClientID,
		ClientSecret: googleConfig.ClientSecret,
		RedirectURL:  googleConfig.RedirectURL,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}

	r := gin.Default()

	// CORS 対応
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		frontendConfig.Origin,
	}
	corsConfig.AllowMethods = []string{
		"POST", "GET", "OPTIONS",
	}
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
	}
	r.Use(cors.New(corsConfig))

	r.GET("/auth/google", func(c *gin.Context) {
		url := oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)

		c.JSON(http.StatusOK, gin.H{"oauth_url": url})
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

		// トークンをサーバに保管
		userTokenRepository := repositories.NewUserTokenRepository(db)
		userTokenRepository.StoreToken(*token)

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

	authorized := r.Group("/")
	authorized.Use(AuthMiddleware(db))

	authorized.GET("/user_info", func(c *gin.Context) {
		user, exists := c.Get("user")

		if exists {
			c.JSON(http.StatusOK, gin.H{
				"user": user,
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
	})

	r.Run(":8080")
}

func isValidAccessToken(db *gorm.DB, accessToken string) (bool, models.User, error) {
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
