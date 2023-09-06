package database

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func buildConnectInfo(config configuration.PostgresConfig) string {
	return " host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.User +
		" dbname=" + config.Dbname +
		" password=" + config.Password +
		" sslmode=disable"
}

func ConnectDB(config configuration.PostgresConfig) *gorm.DB {
	info := buildConnectInfo(config)

	// データベースに接続するコード
	db, err := gorm.Open("postgres", info)
	if err != nil {
		panic(err)
	}

	return db
}

func CloseDB(db *gorm.DB) {
	db.Close()
}
