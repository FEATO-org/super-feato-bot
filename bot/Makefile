build: main.go
	go mod vendor
	go build -ldflags="-s -w" -trimpath -o bin/bot-server main.go

start:
	./bin/bot-server

dev:
	go install github.com/cosmtrek/air@latest
	air

.PHONY: start dev
