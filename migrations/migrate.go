package migrations

import (
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/core"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/models"
	"snow-auth-service/internal/database"
)

var migrationsData = []models.DBMigrationData{
	Get20220427(),
}

func MigrateDB() {
	client := database.SnowDatabaseManager{}.InitializeDatabaseConnection()
	core.MigrateDB(client.DB, "snow-auth", migrationsData)
}
