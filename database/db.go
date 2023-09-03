package database

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/jinzhu/gorm"
)

func makeConnectInfo(config configuration.Config) string {
	return " host=" + config.PostgresHost +
		" port=" + config.PostgresPort +
		" user=" + config.PostgresUser +
		" dbname=" + config.PostgresDbname +
		" sslmode=disable"
}

func ConnectDB() *gorm.DB {
	config := configuration.LoadConfig()

	// データベースに接続するコード
	db, err := gorm.Open("postgres", makeConnectInfo(config))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
