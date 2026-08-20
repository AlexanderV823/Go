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

	dsn := GetDSN(cfg)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("ошибка открытия базы данных: %w", err)
	}

	if err := CheckConnection(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("база данных недоступна: %w", err)
	}

	db.SetMaxOpenConns(25)                 // Максимальное количество открытых соединений
	db.SetMaxIdleConns(25)                 // Максимальное количество простаивающих соединений
	db.SetConnMaxLifetime(5 * time.Minute) // Максимальное время жизни соединения

	return db, nil
}

// Migrate выполняет миграции базы данных в транзакции
func Migrate(db *sql.DB) error {

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
		`CREATE INDEX IF NOT EXISTS idx_posts_created_at ON posts(created_at DESC);`,
		`CREATE INDEX IF NOT EXISTS idx_comments_post_id ON comments(post_id);`,
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("не удалось начать транзакцию миграций: %w", err)
	}

	defer tx.Rollback()

	for _, query := range queries {
		if _, err := tx.Exec(query); err != nil {
			return fmt.Errorf("ошибка при выполнении миграции: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("не удалось применить транзакцию миграций: %w", err)
	}

	log.Println("Миграции базы данных успешно применены")
	return nil
}

// CheckConnection проверяет соединение с базой данных
func CheckConnection(db *sql.DB) error {

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

	if db != nil {
		return db.Close()
	}
	return nil
}

// TestConnection выполняет тестовый запрос к БД
func TestConnection(db *sql.DB) error {

	var result int
	err := db.QueryRow("SELECT 1").Scan(&result)
	if err != nil {
		return fmt.Errorf("тестовый запрос SELECT 1 провалился: %w", err)
	}
	return nil
}
