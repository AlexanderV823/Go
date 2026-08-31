package database_test

import (
	"blog-api/pkg/database"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestMigrate_RealTransaction_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка инициализации sqlmock: %v", err)
	}
	defer db.Close()

	// Настраиваем строгую последовательность вызовов для реального метода Migrate согласно схеме
	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS posts").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS comments").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS idx_posts_author_id").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS idx_comments_post_id").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectExec("CREATE INDEX IF NOT EXISTS idx_posts_created_at").WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	err = database.Migrate(db)
	if err != nil {
		t.Fatalf("настоящая миграция упала со сбоем: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("не все DDL шаги мигратора выполнены внутри транзакции: %v", err)
	}
}

func TestMigrate_RealTransaction_Rollback(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("ошибка инициализации sqlmock: %v", err)
	}
	defer db.Close()

	// Имитируем сбой СУБД на первом же шаге создания таблицы users
	mock.ExpectBegin()
	mock.ExpectExec("CREATE TABLE IF NOT EXISTS users").
		WillReturnError(errors.New("postgres: connection lost or syntax error"))
	mock.ExpectRollback()

	err = database.Migrate(db)
	if err == nil {
		t.Error("мигратор пропустил ошибку СУБД и не вызвал Rollback транзакции")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("транзакция принудительного отката повреждена: %v", err)
	}
}
