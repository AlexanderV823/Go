package db

import (
	"database/sql"
	"encoding/json"
	"hw_1/internal/domain"
)

// SQLiteOrderRepository инкапсулирует работу с конкретной БД SQLite
type SQLiteOrderRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteOrderRepository {
	return &SQLiteOrderRepository{db: db}
}

// Реализация интерфейса service.DBInitializer
func (r *SQLiteOrderRepository) InitSchema() error {
	_, err := r.db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)
	return err
}

// Реализация интерфейса service.OrderWriter
func (r *SQLiteOrderRepository) Save(order *domain.Order) error {
	productsJSON, err := json.Marshal(order.Products)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(
		"INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
		order.Customer, string(productsJSON), order.Total, order.Status,
	)
	return err
}
