include .env

run:
	go run cmd/main.go -env .env

build:
	go build -o eventcalendar cmd/main.go

migrate_up:
	migrate -path internal/migrations -database=pgx://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable up