package handler

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

type PostHandler struct {
	postService *service.PostService
}

func NewPostHandler(postService *service.PostService) *PostHandler {
	return &PostHandler{
		postService: postService,
	}
}

// Create обрабатывает создание нового поста
// POST /api/posts
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	var req model.PostCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	post, err := h.postService.Create(r.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidInput) {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, "failed to create post")
		return
	}

	respondWithJSON(w, http.StatusCreated, post)
}

// GetByID возвращает пост по ID
// GET /api/posts/{id}
func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	post, err := h.postService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			respondWithError(w, http.StatusNotFound, "post not found")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "failed to get post")
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}

// GetAll возвращает список постов с пагинацией
// GET /api/posts?limit=10&offset=0
func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	limit := parseQueryInt(r, "limit", 10)
	offset := parseQueryInt(r, "offset", 0)

	posts, total, err := h.postService.GetAll(r.Context(), limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to fetch posts")
		return
	}

	// Формируем ответ с метаданными пагинации
	response := map[string]interface{}{
		"posts":  posts,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// Update обновляет пост
// PUT /api/posts/{id}
func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	var req model.PostUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		respondWithError(w, http.StatusBadRequest, "invalid json body")
		return
	}

	post, err := h.postService.Update(r.Context(), id, userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			respondWithError(w, http.StatusNotFound, "post not found")
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			respondWithError(w, http.StatusForbidden, "you are not the author of this post")
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			respondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, "failed to update post")
		return
	}

	respondWithJSON(w, http.StatusOK, post)
}

// Delete удаляет пост
// DELETE /api/posts/{id}
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		respondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		respondWithError(w, http.StatusBadRequest, "invalid post id")
		return
	}

	err = h.postService.Delete(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			respondWithError(w, http.StatusNotFound, "post not found")
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			respondWithError(w, http.StatusForbidden, "you are not the author of this post")
			return
		}
		respondWithError(w, http.StatusInternalServerError, "failed to delete post")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetByAuthor возвращает посты конкретного автора
// GET /api/posts/author/{authorID}?limit=10&offset=0
func (h *PostHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/author/")
	authorID, err := strconv.Atoi(idStr)
	if err != nil || authorID <= 0 {
		respondWithError(w, http.StatusBadRequest, "invalid author id")
		return
	}

	limit := parseQueryInt(r, "limit", 10)
	offset := parseQueryInt(r, "offset", 0)

	posts, total, err := h.postService.GetByAuthor(r.Context(), authorID, limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to fetch author posts")
		return
	}

	response := map[string]interface{}{
		"posts":  posts,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	}

	respondWithJSON(w, http.StatusOK, response)
}

// extractIDFromPath извлекает ID из пути URL
func extractIDFromPath(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	// Отрезаем префикс и берем следующую часть пути до возможного слэша
	idPart := strings.TrimPrefix(path, prefix)
	if idx := strings.Index(idPart, "/"); idx != -1 {
		idPart = idPart[:idx]
	}
	return idPart
}

// Вспомогательные функции хендлера

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error":"failed to marshal response"}`))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func parseQueryInt(r *http.Request, key string, defaultVal int) int {
	valStr := r.URL.Query().Get(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}
