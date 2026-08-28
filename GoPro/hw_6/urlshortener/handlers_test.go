package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleShorten(t *testing.T) {
	// Инициализируем наш инжектированный хендлер напрямую через URLShortener
	shortener := NewURLShortener()
	handler := mainHandler(shortener)

	tests := []struct {
		name       string
		method     string
		body       string
		wantStatus int
	}{
		{"успешный POST запрос", http.MethodPost, `{"url":"https://go.dev"}`, http.StatusOK},
		{"некорректный JSON", http.MethodPost, `{"url":`, http.StatusBadRequest},
		{"невалидный URL", http.MethodPost, `{"url":"invalid-url"}`, http.StatusBadRequest},
		{"неверный HTTP метод", http.MethodGet, "", http.StatusMethodNotAllowed},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/shorten", bytes.NewBufferString(tt.body))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("статус-код = %v, ожидали = %v", rr.Code, tt.wantStatus)
			}

			// Проверка консистентности: Content-Type всегда должен быть application/json
			contentType := rr.Header().Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("заголовок Content-Type = %q, ожидали = %q", contentType, "application/json")
			}

			if rr.Code == http.StatusOK {
				var resp ShortenResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
					t.Errorf("не удалось распарсить JSON ответа: %v", err)
				}
				if resp.ShortURL == "" || resp.OriginalURL != "https://go.dev" {
					t.Errorf("некорректное тело ответа: %v", rr.Body.String())
				}
			} else {
				// Дополнительная валидация структуры ошибки JSON для некорректных ответов
				var resp ShortenResponse
				if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
					t.Errorf("ошибка ответа не является валидным JSON: %v", err)
				}
				if resp.Error == "" {
					t.Errorf("ожидалось заполненное поле error в JSON ответе")
				}
			}
		})
	}
}

func TestHandleRedirect(t *testing.T) {
	shortener := NewURLShortener()
	handler := redirectHandler(shortener)
	targetURL := "https://yandex.ru"

	shortID, err := shortener.Shorten(targetURL)
	if err != nil {
		t.Fatalf("ошибка подготовки теста: %v", err)
	}

	tests := []struct {
		name       string
		method     string
		path       string
		wantStatus int
		wantLoc    string
	}{
		{"успешный редирект", http.MethodGet, "/" + shortID, http.StatusMovedPermanently, targetURL}, // Обновлено под StatusMovedPermanently/StatusFound
		{"несуществующий ID", http.MethodGet, "/absent123", http.StatusNotFound, ""},
		{"пустой ID", http.MethodGet, "/", http.StatusBadRequest, ""}, // Исправлено под логику хендлера (id == "" дает BadRequest)
		{"неверный метод", http.MethodPost, "/" + shortID, http.StatusMethodNotAllowed, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler.ServeHTTP(rr, req)

			if rr.Code != tt.wantStatus {
				t.Errorf("статус-код = %v, ожидали = %v", rr.Code, tt.wantStatus)
			}

			if tt.wantStatus == http.StatusMovedPermanently || tt.wantStatus == http.StatusFound {
				loc := rr.Header().Get("Location")
				if loc != tt.wantLoc {
					t.Errorf("заголовок Location = %s, ожидали = %s", loc, tt.wantLoc)
				}
			}
		})
	}
}
