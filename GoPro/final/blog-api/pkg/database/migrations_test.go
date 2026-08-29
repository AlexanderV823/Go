package database_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

// Имитируем внутреннее устройство функции Migrate на базе sqlmock
func runMigrationMock(ctx context.Context, db *sql.DB, sqlSchema string) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Выполняем сырой SQL-скрипт схемы в транзакции
	if _, err := tx.ExecContext(ctx, sqlSchema); err != nil {
		return err
	}

	return tx.Commit()
}

func TestMigrate_SuccessTransaction(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка инициализации sqlmock: %v", err)
	}
	defer db.Close()

	// Ожидаем открытие транзакции, выполнение DDL и успешный коммит
	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	schema := "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY);"
	err = runMigrationMock(context.Background(), db, schema)
	if err != nil {
		t.Fatalf("ожидалась успешная миграция, получен сбой: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все ожидания транзакции СУБД были выполнены: %v", err)
	}
}

func TestMigrate_RollbackOnError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка инициализации sqlmock: %v", err)
	}
	defer db.Close()

	// Ожидаем открытие транзакции, синтаксический сбой при Exec и автоматический Rollback
	mock.ExpectBegin()
	mock.ExpectExec("BROKEN SQL SYNTAX").WillReturnError(errors.New("postgres: syntax error near BROKEN"))
	mock.ExpectRollback()

	brokenSchema := "BROKEN SQL SYNTAX;"
	err = runMigrationMock(context.Background(), db, brokenSchema)
	if err == nil {
		t.Error("мигратор пропустил сломанный SQL синтаксис без вызова Rollback")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("транзакция не была откачена корректно: %v", err)
	}
}
