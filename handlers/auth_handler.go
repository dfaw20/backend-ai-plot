package handlers

import (
	"context"
	"net/http"

	"github.com/dfaw20/backend-ai-plot/repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

type AuthHandler struct {
	oauth2Config        oauth2.Config
	userRepository      repositories.UserRepository
	userTokenRepository repositories.UserTokenRepository
}

func NewAuthHandler(
	oauth2Config oauth2.Config,
	userRepository repositories.UserRepository,
	userTokenRepository repositories.UserTokenRepository,
) AuthHandler {
	return AuthHandler{oauth2Config, userRepository, userTokenRepository}
}

func (h *AuthHandler) GetOAuthURL(c *gin.Context) {
	url := h.oauth2Config.AuthCodeURL("", oauth2.AccessTypeOffline)

	c.JSON(http.StatusOK, gin.H{"oauth_url": url})
}

func (h *AuthHandler) GetAuthGoogleCallback(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing code parameter"})
		return
	}

	// OAuth2トークンを取得し、ユーザー情報を取得できます。
	token, err := h.oauth2Config.Exchange(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザー情報の取得やデータベースへの保存などをここで行います。

	// トークンをサーバに保管
	h.userTokenRepository.StoreToken(*token)

	// アクセストークンを使用してユーザ情報を取得
	cxt := context.Background()
	oauth2Service, err := v2.NewService(cxt, option.WithTokenSource(h.oauth2Config.TokenSource(cxt, token)))
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
	user, err := h.userRepository.CreateOrSyncUser(*userInfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// ユーザー情報の表示
	c.JSON(http.StatusOK, gin.H{"token": token, "user": user})
}
