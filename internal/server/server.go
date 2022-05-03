package server

import (
	"database/sql"
	"fmt"
	"github.com/SnowLynxSoftware/go-mysql-data-core/pkg/core"
	"log"
	"net/http"
	"os"
	"snow-auth-service/internal/controllers"
)

// Start Configures everything and starts the web server.
// This is where all the initialization of the server actually happens.
func Start() {

	// Create database connection
	client := core.CreateMySQLClient()
	connectionString := os.Getenv("DB_CONNECTION_STRING")
	db, err := client.Connect(connectionString, true)
	if err != nil {
		panic(err)
	}

	appPort := os.Getenv("APP_PORT")
	s := &http.Server{
		Addr: ":" + appPort,
	}
	fmt.Printf("App is listening on port: [%s]", appPort)
	// defer the close till after the Start function has finished
	// executing
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	log.Fatal(s.ListenAndServe())
}

func defineRoutes() {
	http.HandleFunc("/health", controllers.HealthController)
	http.HandleFunc("/auth/register", controllers.AuthControllerRegister)
}
