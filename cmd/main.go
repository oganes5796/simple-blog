package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/oganes5796/simple-blog/pkg/handler"
	"github.com/oganes5796/simple-blog/pkg/logger"
	"github.com/oganes5796/simple-blog/pkg/repository"
	"github.com/oganes5796/simple-blog/pkg/server"
	"github.com/oganes5796/simple-blog/pkg/service"
)

func main() {
	logger := logger.NewLogger()

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file:" + err.Error())
		return
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("SSLMode"),
	})
	if err != nil {
		logger.Error("Error connecting to database:" + err.Error())
		return
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
			logger.Error("Error starting server:" + err.Error())
		}
	}()

	logger.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logger.Info("Shutting down server")

	if err := srv.Shutdown(); err != nil {
		logger.Error("Error shutting down server:" + err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Error("Error closing database connection:" + err.Error())
	}
}
