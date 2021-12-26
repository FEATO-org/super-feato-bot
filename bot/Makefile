build: src/main.go
	make clean
	go mod vendor
	go build -ldflags="-s -w" -trimpath -o bin/bot-server $<

start:
	./bin/bot-server

dev:
	go install github.com/cosmtrek/air@latest
	air

clean:
	rm -rf bin tmp

.PHONY: start dev clean
