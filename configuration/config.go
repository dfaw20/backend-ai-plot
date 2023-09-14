package configuration

import (
	"encoding/json"
	"os"
)

type FrontendConfig struct {
	Origins []string `json:"origins"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Dbname   string `json:"dbname"`
	Password string `json:"password"`
}

type GoogleConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

type Config struct {
	Frontend   FrontendConfig `json:"frontend"`
	Postgres   PostgresConfig `json:"postgres"`
	Google     GoogleConfig   `json:"google"`
	Production bool           `json:"production"`
}

func LoadConfig() Config {
	// 設定ファイルから設定情報を読み込む
	configData, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	if err := json.Unmarshal(configData, &config); err != nil {
		panic(err)
	}

	return config
}
