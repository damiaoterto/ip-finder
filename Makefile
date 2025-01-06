APP_NAME="ip-finder"

build:
	@CGO_ENABLED=0 go build -o $(APP_NAME) main.go

clear:
	@go clean ./...
	@rm -f $(APP_NAME)