package configuration

import (
	"encoding/json"
	"os"
)

type Config struct {
	PostgresHost       string `json:"postgres.host"`
	PostgresPort       string `json:"postgres.port"`
	PostgresUser       string `json:"postgres.user"`
	PostgresDbname     string `json:"postgres.dbname"`
	GoogleClientID     string `json:"google.client_id"`
	GoogleClientSecret string `json:"google.client_secret"`
	GoogleRedirectURL  string `json:"google.redirect_url"`
}

func LoadConfig() Config {
	// 設定ファイルから設定情報を読み込む
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	return config
}
