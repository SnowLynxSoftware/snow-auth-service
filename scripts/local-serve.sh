# Run the local dev server.
# shellcheck disable=SC2046
# shellcheck disable=SC2002
env $(cat ./.env | xargs) go run main.go serve
