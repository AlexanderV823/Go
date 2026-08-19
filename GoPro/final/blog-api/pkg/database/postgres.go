package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// Config содержит параметры подключения к PostgreSQL
type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewPostgresDB создает новое подключение к PostgreSQL
func NewPostgresDB(cfg Config) (*sql.DB, error) {
	// 1. Сформировать строку подключения (DSN) из параметров конфигурации
	dsn := GetDSN(cfg)

	// 2. Открыть соединение с БД используя sql.Open("postgres", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия базы данных: %w", err)
	}

	// 3. Проверить соединение методом Ping()
	if err := CheckConnection(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("база данных недоступна: %w", err)
	}

	// 4. Настроить пул соединений (SetMaxOpenConns, SetMaxIdleConns)
	db.SetMaxOpenConns(25)                 // Максимальное количество открытых соединений
	db.SetMaxIdleConns(25)                 // Максимальное количество простаивающих соединений
	db.SetConnMaxLifetime(5 * time.Minute) // Максимальное время жизни соединения

	// 5. Вернуть подключение или ошибку
	return db, nil
}

// Migrate выполняет миграции базы данных в транзакции
func Migrate(db *sql.DB) error {
	// 1-4. SQL запросы для создания таблиц и необходимых индексов
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) UNIQUE NOT NULL,
			email VARCHAR(100) UNIQUE NOT NULL,
			password_hash VARCHAR(255) NOT NULL,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,
		`CREATE TABLE IF NOT EXISTS posts (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			author_id INT REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,
		`CREATE TABLE IF NOT EXISTS comments (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			post_id INT REFERENCES posts(id) ON DELETE CASCADE,
			author_id INT REFERENCES users(id) ON DELETE CASCADE,
			created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
			updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
		);`,
		// Дополнительные индексы для ускорения выборок и пагинации
		`CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC);`,
		`CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);`,
	}

	// 5. Выполнить каждый запрос в транзакции
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию миграций: %w", err)
	}

	// Гарантируем откат транзакции в случае паники или ошибки до коммита
	defer tx.Rollback()

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			return fmt.Errorf("ошибка при выполнении миграции: %w", err)
		}
	}

	// Коммитим транзакцию, если все запросы прошли успешно
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("не удалось применить транзакцию миграций: %w", err)
	}

	log.Println("Миграции базы данных успешно применены")
	return nil
}

// CheckConnection проверяет соединение с базой данных
func CheckConnection(db *sql.DB) error {
	// Использовать db.Ping() для проверки
	return db.Ping()
}

// GetDSN формирует строку подключения к PostgreSQL
func GetDSN(cfg Config) string {
	// Формат: "host=%s port=%d user=%s password=%s dbname=%s sslmode=%s"
	if cfg.SSLMode == "" {
		cfg.SSLMode = "disable"
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)
}

// Close закрывает соединение с базой данных
func Close(db *sql.DB) error {
	// Корректно закрыть соединение
	if db != nil {
		return db.Close()
	}
	return nil
}

// TestConnection выполняет тестовый запрос к БД
func TestConnection(db *sql.DB) error {
	// Выполнить простой запрос для проверки работы БД
	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		return fmt.Errorf("тестовый запрос SELECT 1 провалился: %w", err)
	}
	return nil
}
