include bot/Makefile

production:
	cd "$(PWD)/bot" && make build
	cd "$(PWD)/bot" && make start

.PHONY: production
