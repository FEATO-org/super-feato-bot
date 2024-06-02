schema = db/schema.sql

build: cmd/server/main.go
	CGO_ENABLED=0 go build -ldflags="-s -w" -mod=readonly -trimpath -o bin/server $^

start:
	./bin/server

dev:
	air -c .air.toml

clean:
	rm -rf bin tmp

schema.update:
	make migrate model.build

model.build: $(schema) db/query.sql
	sqlc generate

migrate: $(schema)
	mysqldef -h $(DBHOST) sfs -u ${DBUSER} -p ${DBPASSWORD} < $<

migrate.dry:
	mysqldef -h $(DBHOST) sfs -u ${DBUSER} -p ${DBPASSWORD} --dry-run < $(schema)

db.console:
	mysql -h $(DBHOST) sfs -u ${DBUSER} -p${DBPASSWORD}

.PHONY: start dev clean migrate.dry db.console
