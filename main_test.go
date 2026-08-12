package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

// Глобальная переменная для управления моками в тестах
var mock sqlmock.Sqlmock

func init() {

	var err error
	var mockDB *sql.DB
	// Инициализируем sqlmock
	mockDB, mock, err = sqlmock.New(sqlmock.MonitorPingsOption(true))
	if err != nil {
		panic("failed to open a stub database connection: " + err.Error())
	}
	db = mockDB
}

// Тест эндпоинта /health
func TestHealthHandler(t *testing.T) {

	// Настраиваем ожидание Ping()
	mock.ExpectPing()

	req, err := http.NewRequest("GET", "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("HealthHandler returned wrong status: got %v want %v", rr.Code, http.StatusOK)
	}

	if !strings.Contains(rr.Body.String(), `"status":"ok"`) {
		t.Errorf("HealthHandler returned unexpected body: %s", rr.Body.String())
	}
}

// Тест эндпоинта /register
func TestRegisterHandler(t *testing.T) {

	email := "testuser@example.com"
	username := "test_user"

	// 1. Имитируем запрос из UserExistsByEmail (SELECT EXISTS...)
	// Возвращаем false (пользователь не существует)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)")).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"exists"}).AddRow(false))

	// 2. Имитируем запрос из CreateUser (INSERT INTO... RETURNING id, created_at)
	// Возвращаем сгенерированный ID = 1 и текущее время
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO users (email, username, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at")).
		WithArgs(email, username, sqlmock.AnyArg()). // Пароль будет хэширован, проверяем через AnyArg
		WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(1, time.Now()))

	payload := RegisterRequest{
		Email:    email,
		Username: username,
		Password: "SecurePass123",
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/register", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(RegisterHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("RegisterHandler returned wrong status: got %v want %v. Body: %s", rr.Code, http.StatusCreated, rr.Body.String())
	}

	var resp AuthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v. Body was: %s", err, rr.Body.String())
	}

	if resp.Token == "" {
		t.Error("RegisterHandler did not return a token")
	}
}

// Тест эндпоинта /login
func TestLoginHandler(t *testing.T) {

	email := "testuser@example.com"

	// Хэшируем тестовый пароль, чтобы функция CheckPassword внутри хендлера выдала true
	hashedPassword, _ := HashPassword("SecurePass123")

	// Имитируем запрос из GetUserByEmail
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email, username, password_hash, created_at FROM users WHERE email = $1")).
		WithArgs(email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "password_hash", "created_at"}).
			AddRow(1, email, "test_user", hashedPassword, time.Now()))

	payload := LoginRequest{
		Email:    email,
		Password: "SecurePass123",
	}
	body, _ := json.Marshal(payload)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LoginHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("LoginHandler returned wrong status: got %v want %v. Body: %s", rr.Code, http.StatusOK, rr.Body.String())
	}

	var resp AuthResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}

	if resp.Token == "" {
		t.Error("LoginHandler did not return JWT token")
	}
}

// Тест эндпоинта /profile
func TestProfileHandler(t *testing.T) {

	userID := 1

	// Имитируем запрос из GetUserByID (обратите внимание, что password_hash не запрашивается)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, email, username, created_at FROM users WHERE id = $1")).
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "created_at"}).
			AddRow(userID, "testuser@example.com", "test_user", time.Now()))

	req, err := http.NewRequest("GET", "/profile", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Передаем правильный ключ контекста для прохождения авторизации в тесте
	ctx := context.WithValue(req.Context(), userIDKey, userID)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(ProfileHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("ProfileHandler returned wrong status: got %v want %v. Body: %s", rr.Code, http.StatusOK, rr.Body.String())
	}

	var resp ProfileResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatalf("Failed to unmarshal profile response: %v", err)
	}

	if resp.User.ID != userID {
		t.Errorf("Expected user ID %d, got %d", userID, resp.User.ID)
	}
}
