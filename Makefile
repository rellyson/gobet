set-env:
	go install github.com/cosmtrek/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.50.0
	sh ./scripts/install-git-hooks.sh

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/gobet cmd/gobet/main.go

dev:
	air -c .air.toml

update-swagger:
	swag init -g pkg/http/server.go -o api/openapi
