package database

import (
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/core"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/database"
	"os"
)

type SnowDatabaseManager struct{}

// InitializeDatabaseConnection Initialize a connection to the database and return an instance
// of the active client.
func (databaseManager SnowDatabaseManager) InitializeDatabaseConnection() database.MySQLDB {
	client := core.CreateMySQLClient()
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	_, err := client.Connect(connectionString, true)
	if err != nil {
		panic(err)
	} else {
		return client
	}
}
