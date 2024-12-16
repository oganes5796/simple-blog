package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/oganes5796/simple-blog/pkg/logger"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

const (
	usersTable = "users"
)

func NewPostgresDB(config Config) (*sqlx.DB, error) {
	logger := logger.NewLogger()
	// Формируем строку подключения
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)

	// Ожидаем доступности базы данных
	db, err := waitForDB(dsn, 10)
	if err != nil {
		logger.Error("Ошибка подключения к базе данных:" + err.Error())
		return nil, err

	}

	// Проверяем подключение
	if err := db.Ping(); err != nil {
		logger.Error("База данных недоступна:" + err.Error())
		return nil, err

	}

	return db, nil
}

// waitForDB ждет доступности базы данных
func waitForDB(dsn string, maxRetries int) (*sqlx.DB, error) {
	var db *sqlx.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		db, err = sqlx.Open("postgres", dsn)
		if err == nil && db.Ping() == nil {
			return db, nil
		}

		fmt.Printf("Попытка %d: не удалось подключиться к базе данных, повтор через 2 секунды...\n", i+1)
		time.Sleep(2 * time.Second)
	}
	return nil, fmt.Errorf("не удалось подключиться к базе данных: %w", err)
}
