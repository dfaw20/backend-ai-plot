package database

import (
	"log"

	"github.com/dfaw20/backend-ai-plot/configuration"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func makeConnectInfo(config configuration.Config) string {
	return " host=" + config.Postgres.Host +
		" port=" + config.Postgres.Port +
		" user=" + config.Postgres.User +
		" dbname=" + config.Postgres.Dbname +
		" password=" + config.Postgres.Password +
		" sslmode=disable"
}

func ConnectDB() *gorm.DB {
	config := configuration.LoadConfig()
	info := makeConnectInfo(config)

	log.Print(info)
	log.Print(config)

	// データベースに接続するコード
	db, err := gorm.Open("postgres", info)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return db
}
