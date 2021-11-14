include .env

.PHONY: run
run:
	go run cmd/main.go -env .env

.PHONY: build
build:
	go build -o eventcalendar cmd/main.go

.PHONY: migration
migration:
	@migrate -path internal/migrations \
		-database=pgx://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable \
		$(filter-out $@,$(MAKECMDGOALS))