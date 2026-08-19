package main

import (
	"testing"
)

func TestURLShortener_Shorten(t *testing.T) {
	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"валидный HTTP URL", "http://example.com", false},
		{"валидный HTTPS URL", "https://google.com", false},
		{"невалидный URL", "not-a-url", true},
		{"пустая строка", "", true},
		{"неподдерживаемая схема ftp", "ftp://files.com", true},
	}

	shortener := NewURLShortener()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shortID, err := shortener.Shorten(tt.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("ошибка = %v, ожидали ошибку = %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(shortID) < 6 || len(shortID) > 8 {
					t.Errorf("короткий ID имеет неверную длину: %s (длина %d)", shortID, len(shortID))
				}
			}
		})
	}
}

func TestURLShortener_GetOriginal(t *testing.T) {
	shortener := NewURLShortener()
	originalURL := "https://github.com"

	shortID, err := shortener.Shorten(originalURL)
	if err != nil {
		t.Fatalf("не удалось создать короткий URL для теста: %v", err)
	}

	tests := []struct {
		name    string
		id      string
		wantURL string
		wantErr bool
	}{
		{"успешное получение", shortID, originalURL, false},
		{"несуществующий ID", "nonexistent", "", true},
		{"пустой ID", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, err := shortener.GetOriginal(tt.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ошибка = %v, ожидали ошибку = %v", err, tt.wantErr)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("получили = %v, ожидали = %v", gotURL, tt.wantURL)
			}
		})
	}
}
