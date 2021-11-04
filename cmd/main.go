package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/event-calendar/internal/pgdb"
	"github.com/grum261/event-calendar/internal/rest"
	"github.com/grum261/event-calendar/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	var envPath string

	flag.StringVar(&envPath, "env", "", "")
	flag.Parse()

	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	pool, err := newDB(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	})

	s := pgdb.NewStore(pool)
	svc := service.NewRepositories(s.Tag, s.City, s.Event, s.EventPart)
	h := rest.NewHandlers(svc.Tag, svc.City, svc.Event)

	h.RegisterRoutes(app.Group("/api/v1"))

	log.Fatal(app.Listen(":8000"))
}

func newDB(ctx context.Context) (*pgxpool.Pool, error) {
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connectionURL := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s", username, password, host, port, dbName)

	config, err := pgxpool.ParseConfig(connectionURL)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		return nil, err
	}

	return pool, nil
}
