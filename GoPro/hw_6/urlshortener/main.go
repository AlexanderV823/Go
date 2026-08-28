package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct {
	ShortURL    string `json:"short_url,omitempty"`
	OriginalURL string `json:"original_url,omitempty"`
	Error       string `json:"error,omitempty"` // Добавлено поле для красивых JSON-ошибок
}

// Хелпер для отправки JSON-ответов (включая ошибки)
func respondWithError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(ShortenResponse{Error: message})
}

func mainHandler(shortener *URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			respondWithError(w, http.StatusMethodNotAllowed, "Method not allowed")
			return
		}

		var req ShortenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid JSON payload")
			return
		}

		req.URL = strings.TrimSpace(req.URL)
		if req.URL == "" {
			respondWithError(w, http.StatusBadRequest, "URL cannot be empty")
			return
		}

		// Простая валидация схемы
		if !strings.HasPrefix(req.URL, "http://") && !strings.HasPrefix(req.URL, "https://") {
			respondWithError(w, http.StatusBadRequest, "URL must start with http:// or https://")
			return
		}

		shortID, err := shortener.Shorten(req.URL)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ShortenResponse{
			ShortURL:    shortID,
			OriginalURL: req.URL,
		})
	}
}

func redirectHandler(shortener *URLShortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		id := strings.TrimPrefix(r.URL.Path, "/")
		if id == "" {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if url, exists := shortener.Resolve(id); exists {
			http.Redirect(w, r, url, http.StatusMovedPermanently) // Или StatusFound (302)
			return
		}

		http.Error(w, "Link not found", http.StatusNotFound)
	}
}
