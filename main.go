package main

import (
	cmd "snow-auth-service/cmds"
)

// main This serves as the entry-point to the application--which simply
// calls the Start() method on the server package.
func main() {
	cmd.Execute()
}
