package database

import (
	"github.com/dfaw20/backend-ai-plot/configuration"
	//_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func buildConnectDsn(config configuration.PostgresConfig) string {
	return " host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.User +
		" dbname=" + config.Dbname +
		" password=" + config.Password +
		" sslmode=disable"
}

func ConnectDB(config configuration.PostgresConfig) *gorm.DB {
	dsn := buildConnectDsn(config)

	// データベースに接続するコード
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	return db
}
