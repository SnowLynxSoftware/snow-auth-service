# Run the .env file to set the local ENV variables.
./.env

# Run the local dev server.
env $(cat ./.env | xargs) go run main.go
