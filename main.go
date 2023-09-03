package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

var (
	oauth2Config oauth2.Config
)

func main() {
	// 設定ファイルから設定情報を読み込む
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	// OAuth2設定を構築
	oauth2Config = oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  config.RedirectURL,
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

		// ユーザー情報の表示
		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	r.Run(":8080")
}
