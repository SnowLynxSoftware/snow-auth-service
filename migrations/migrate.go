package migrations

import (
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/core"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/models"
	"os"
)

var migrationsData = []models.DBMigrationData{
	Get20220427(),
}

func MigrateDB() {
	client := core.CreateMySQLClient()
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := client.Connect(connectionString, true)
	if err != nil {
		panic(err)
	}
	core.MigrateDB(db, migrationsData)
}
