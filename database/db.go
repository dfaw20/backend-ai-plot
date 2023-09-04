package database

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func buildConnectInfo(config configuration.Config) string {
	return " host=" + config.Postgres.Host +
		" port=" + config.Postgres.Port +
		" user=" + config.Postgres.User +
		" dbname=" + config.Postgres.Dbname +
		" password=" + config.Postgres.Password +
		" sslmode=disable"
}

func ConnectDB(config configuration.Config) *gorm.DB {
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
