package main

import (
	"auth-service/internal/server"
)

// main This serves as the entry-point to the application--which simply
// calls the Start() method on the server package.
func main() {
	server.Start()
}
