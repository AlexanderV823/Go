package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
	"sync"
)

type URLShortener struct {
	urls map[string]string
	mu   sync.RWMutex
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}

func (us *URLShortener) Shorten(originalURL string) (string, error) {
	if !isValidURL(originalURL) {
		return "", errors.New("invalid or unsupported URL scheme")
	}

	us.mu.Lock()
	defer us.mu.Unlock()

	for {
		shortID := generateShortID()
		if _, exists := us.urls[shortID]; !exists {
			us.urls[shortID] = originalURL
			return shortID, nil
		}
	}
}

func (us *URLShortener) GetOriginal(shortID string) (string, error) {
	us.mu.RLock()
	defer us.mu.RUnlock()

	original, exists := us.urls[shortID]
	if !exists {
		return "", errors.New("url not found")
	}
	return original, nil
}

func generateShortID() string {
	b := make([]byte, 5)
	_, err := rand.Read(b)
	if err != nil {
		return "fallback"
	}
	id := base64.RawURLEncoding.EncodeToString(b)
	return id
}

func isValidURL(str string) bool {
	u, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	scheme := strings.ToLower(u.Scheme)
	return scheme == "http" || scheme == "https"
}
