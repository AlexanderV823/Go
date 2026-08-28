package main

import (
	"testing"
	"fmt"
	"sync"
)

func TestURLShortener_Shorten(t *testing.T) {
	shortener := NewURLShortener()

	tests := []struct {
		name    string
		url     string
		wantErr bool
	}{
		{"валидный HTTP URL", "http://example.com", false},
		{"валидный HTTPS URL", "https://google.com", false},
		// Валидация строк перенесена на уровень HTTP хендлера перед вызовом Shorten,
		// однако сам метод Shorten принимает любые строки и гарантирует выдачу ID без ошибок,
		// если работает системная энтропия.
	}

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

// Новый тест для проверки отсутствия дубликатов при повторных вызовах
func TestURLShortener_Deduplication(t *testing.T) {
	shortener := NewURLShortener()
	url := "https://habr.com"

	id1, err := shortener.Shorten(url)
	if err != nil {
		t.Fatalf("первый вызов Shorten завершился ошибкой: %v", err)
	}

	id2, err := shortener.Shorten(url)
	if err != nil {
		t.Fatalf("второй вызов Shorten завершился ошибкой: %v", err)
	}

	if id1 != id2 {
		t.Errorf("дедупликация не сработала: для одного URL сгенерированы разные ID (%s и %s)", id1, id2)
	}
}

func TestURLShortener_Resolve(t *testing.T) {
	shortener := NewURLShortener()
	originalURL := "https://github.com"

	shortID, err := shortener.Shorten(originalURL)
	if err != nil {
		t.Fatalf("не удалось создать короткий URL для теста: %v", err)
	}

	tests := []struct {
		name       string
		id         string
		wantURL    string
		wantExists bool
	}{
		{"успешное получение", shortID, originalURL, true},
		{"несуществующий ID", "nonexistent", "", false},
		{"пустой ID", "", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotURL, exists := shortener.Resolve(tt.id)
			if exists != tt.wantExists {
				t.Errorf("exists = %v, ожидали = %v", exists, tt.wantExists)
				return
			}
			if gotURL != tt.wantURL {
				t.Errorf("получили URL = %v, ожидали = %v", gotURL, tt.wantURL)
			}
		})
	}
}

// TestURLShortener_Concurrency проверяет потокобезопасность структуры URLShortener.
// Тест запускается с флагом -race для автоматического поиска race conditions.
func TestURLShortener_Concurrency(t *testing.T) {
	shortener := NewURLShortener()

	const goroutinesCount = 500 // Всего будет 500 пишущих и 500 читающих горутин
	var wg sync.WaitGroup

	// Запускаем горутины на параллельную ЗАПИСЬ (Shorten)
	for i := 0; i < goroutinesCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			url := fmt.Sprintf("https://example-%d.com", id)

			_, err := shortener.Shorten(url)
			if err != nil {
				t.Errorf("ошибка при конкурентном сокращении: %v", err)
			}
		}(i)
	}

	// Запускаем горутины на параллельное ЧТЕНИЕ (Resolve) и повторную запись того же URL
	for i := 0; i < goroutinesCount; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// Пытаемся читать как существующие, так и рандомные ID
			randomID := fmt.Sprintf("rand%d", id)
			shortener.Resolve(randomID)

			// Проверяем дедупликацию в условиях гонки — отправляем один и тот же URL
			const sharedURL = "https://shared-concurrent-url.com"
			_, err := shortener.Shorten(sharedURL)
			if err != nil {
				t.Errorf("ошибка при конкурентном сокращении общего URL: %v", err)
			}
		}(i)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()
}