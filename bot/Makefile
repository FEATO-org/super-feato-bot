build: src/main.go
	go build -ldflags="-s -w" -trimpath -o bin/server $<

start:
	./bin/server

dev:
	air -c .air.toml

clean:
	rm -rf bin tmp

.PHONY: start dev clean
