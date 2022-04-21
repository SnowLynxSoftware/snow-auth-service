package server

import (
	"auth-service/internal/controllers"
	"fmt"
	"log"
	"net/http"
	"os"
)

// Start Configures everything and starts the web server.
// This is where all the initialization of the server actually happens.
func Start() {
	http.HandleFunc("/", controllers.HealthController)
	appPort := os.Getenv("APP_PORT")
	s := &http.Server{
		Addr: ":" + appPort,
	}
	fmt.Printf("App is listening on port: [%s]", appPort)
	log.Fatal(s.ListenAndServe())
}
