package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/grum261/event-calendar/internal/postgresql"
	"github.com/grum261/event-calendar/internal/rest"
	"github.com/grum261/event-calendar/internal/service"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	var envPath string

	flag.StringVar(&envPath, "env", "", "")
	flag.Parse()

	errCh, err := run(envPath)
	if err != nil {
		log.Fatal(err)
	}

	if err := <-errCh; err != nil {
		log.Fatal(err)
	}
}

type serverConfig struct {
	Address     string
	DB          *pgxpool.Pool
	Logger      *zap.Logger
	Middlewares []func() fiber.Handler
}

func newServer(conf serverConfig) (*fiber.App, error) {
	app := fiber.New(fiber.Config{
		ReadTimeout:  time.Second,
		WriteTimeout: time.Second,
		IdleTimeout:  time.Second,
	})

	r := app.Group("/api/v1")

	for _, mw := range conf.Middlewares {
		app.Use(mw())
	}

	store := postgresql.NewStore(conf.DB)
	svc := service.NewServices(store.Tag, store.Event, store.EventPart)
	h := rest.NewHandlers(svc.Tag, svc.Event, svc.EventPart)

	h.RegisterRoutes(r)

	return app, nil
}

func run(env string) (<-chan error, error) {
	if err := godotenv.Load(env); err != nil {
		return nil, err
	}

	logger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	pool, err := newDB()
	if err != nil {
		return nil, err
	}

	logging := func() fiber.Handler {
		return func(c *fiber.Ctx) error {
			fields := []zapcore.Field{
				zap.Time("time", time.Now()),
				zap.String("method", c.Method()),
				zap.String("uri", c.OriginalURL()),
			}

			if err := c.Next(); err != nil {
				return err
			}

			status := c.Response().StatusCode()
			switch {
			case status >= 500:
				fields = append(fields, zap.Int("status", status))
				logger.Error("Ошибка сервера", fields...)
			case status >= 400:
				fields = append(fields, zap.Int("status", status))
				logger.Warn("Ошибка клиента", fields...)
			case status >= 300:
				fields = append(fields, zap.Int("status", status))
				logger.Info("Редирект", fields...)
			default:
				fields = append(fields, zap.Int("status", status))
				logger.Info("OK", fields...)
			}

			return nil
		}
	}

	srv, err := newServer(serverConfig{
		Address:     ":8000",
		DB:          pool,
		Logger:      logger,
		Middlewares: []func() func(*fiber.Ctx) error{logging},
	})
	if err != nil {
		return nil, err
	}

	errCh := make(chan error, 1)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-ctx.Done()

		logger.Info("Получен сигнал остановки сервера")

		ctxTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)

		defer func() {
			_ = logger.Sync()

			pool.Close()
			stop()
			cancel()
			close(errCh)
		}()

		<-ctxTimeout.Done()

		if err := srv.Shutdown(); err != nil {
			errCh <- err
		}

		logger.Info("Сервер остановлен")
	}()

	go func() {
		logger.Info("Начинаем слушать", zap.String("address", ":8000"))

		if err := srv.Listen(":8000"); err != nil {
			errCh <- err
		}
	}()

	return errCh, nil
}

func newDB() (*pgxpool.Pool, error) {
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

	pool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	return pool, nil
}
