package logger

import (
	"log/slog"
	"os"
)

type Logger struct {
}

func NewLogger() *slog.Logger {
	// Создаем JSON-хендлер
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	// Устанавливаем как глобальный логгер
	slog.SetDefault(logger)
	return logger
}
