package server

import (
	"fmt"
	database2 "github.com/SnowLynxSoftware/go-mysql-data-core/pkg/database"
	"log"
	"net/http"
	"os"
	"snow-auth-service/internal/database"
	"snow-auth-service/internal/routers"
)

// Start Configures everything and starts the web server.
// This is where all the initialization of the server actually happens.
func Start() {
	// Create database connection
	client := database.SnowDatabaseManager{}.InitializeDatabaseConnection()

	appPort := os.Getenv("APP_PORT")
	s := &http.Server{
		Addr: ":" + appPort,
	}

	// Setup Routers
	defineRoutes()

	fmt.Printf("App is listening on port: [%s]", appPort)
	// defer the close till after the Start function has finished
	// executing
	defer func(client database2.MySQLDB) {
		err := client.CloseConnection()
		if err != nil {
			panic(err)
		}
	}(client)
	log.Fatal(s.ListenAndServe())
}

func defineRoutes() {
	//http.HandleFunc("/health", controllers.HealthController)
	//http.HandleFunc("/auth/register", controllers.AuthControllerRegister)
	http.HandleFunc("/", routers.Serve)
}
