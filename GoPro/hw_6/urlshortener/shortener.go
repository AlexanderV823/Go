package main

import (
	"crypto/rand"
	"errors"
	"math/big"
	"sync"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idLength = 6
)

type URLShortener struct {
	mu         sync.RWMutex
	shortToURL map[string]string
	urlToShort map[string]string // Новая мапа для предотвращения дубликатов
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		shortToURL: make(map[string]string),
		urlToShort: make(map[string]string),
	}
}

// Измененная функция теперь возвращает ошибку, если crypto/rand дал сбой
func (s *URLShortener) generateShortID() (string, error) {
	b := make([]byte, idLength)
	for i := range b {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
		if err != nil {
			return "", err // Вместо "fallback" возвращаем реальную ошибку
		}
		b[i] = alphabet[num.Int64()]
	}
	return string(b), nil
}

func (s *URLShortener) Shorten(longURL string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Если URL уже сокращался, возвращаем существующий ID
	if existingID, exists := s.urlToShort[longURL]; exists {
		return existingID, nil
	}

	// Генерируем новый ID с обработкой ошибки
	shortID, err := s.generateShortID()
	if err != nil {
		return "", errors.New("failed to generate unique ID due to system entropy issue")
	}

	// Защита от коллизий в мапе
	for {
		if _, exists := s.shortToURL[shortID]; !exists {
			break
		}
		shortID, err = s.generateShortID()
		if err != nil {
			return "", errors.New("failed to generate unique ID due to system entropy issue")
		}
	}

	s.shortToURL[shortID] = longURL
	s.urlToShort[longURL] = shortID // Сохраняем обратное соответствие

	return shortID, nil
}

func (s *URLShortener) Resolve(shortID string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	url, exists := s.shortToURL[shortID]
	return url, exists
}
