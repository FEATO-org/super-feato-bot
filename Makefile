schema = db/schema.sql

build: cmd/server/main.go
	CGO_ENABLED=0 go build -ldflags="-s -w" -mod=readonly -trimpath -o bin/server $^

start:
	./bin/server

dev:
	air -c .air.toml

clean:
	rm -rf bin tmp

schema.update: $(schema) db/query.sql
	sqlc generate

migrate: $(schema)
	psqldef -h $(PGHOST) $(APP_ENV) < $<
	make schema.update

migrate.dry:
	psqldef -h $(PGHOST) $(APP_ENV) --dry-run < $(schema)

setup:
	./tools/build-db.sh

.PHONY: start dev clean migrate.dry migrate setup
