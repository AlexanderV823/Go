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
// Требует аутентификации
func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Access token signature error", http.StatusUnauthorized)
		return
	}

	var req model.PostCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid parameters payload", http.StatusBadRequest)
		return
	}

	post, err := h.postService.Create(r.Context(), userID, &req)
	if err != nil {
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post)
}

// GetByID возвращает пост по ID
// GET /api/posts/{id}
// Не требует аутентификации
func (h *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Invalid unique entity identity", http.StatusBadRequest)
		return
	}

	post, err := h.postService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// GetAll возвращает список постов с пагинацией
// GET /api/posts?limit=10&offset=0
// Не требует аутентификации
func (h *PostHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	posts, total, err := h.postService.GetAll(r.Context(), limit, offset)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type PaginatedResponse struct {
		Posts  []*model.Post `json:"posts"`
		Total  int           `json:"total"`
		Limit  int           `json:"limit"`
		Offset int           `json:"offset"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(PaginatedResponse{
		Posts:  posts,
		Total:  total,
		Limit:  limit,
		Offset: offset,
	})
}

// Update обновляет пост
// PUT /api/posts/{id}
// Требует аутентификации, может обновить только автор
func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Unauthorized entity validation", http.StatusUnauthorized)
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Invalid post context identity", http.StatusBadRequest)
		return
	}

	var req model.PostUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid updates metadata structure", http.StatusBadRequest)
		return
	}

	post, err := h.postService.Update(r.Context(), id, userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			writeError(w, "You cannot configure alternative entities", http.StatusForbidden)
			return
		}
		writeError(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

// Delete удаляет пост
// DELETE /api/posts/{id}
// Требует аутентификации, может удалить только автор
func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Context credentials missing", http.StatusUnauthorized)
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/posts/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Incorrect post layout key", http.StatusBadRequest)
		return
	}

	err = h.postService.Delete(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, service.ErrPostNotFound) {
			writeError(w, err.Error(), http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			writeError(w, "Modification restricted", http.StatusForbidden)
			return
		}
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetByAuthor возвращает посты конкретного автора
// GET /api/posts/author/{authorID}?limit=10&offset=0
// Не требует аутентификации
func (h *PostHandler) GetByAuthor(w http.ResponseWriter, r *http.Request) {
	// Дополнительный функционал
	http.Error(w, "Author selection dynamic workflow mapping", http.StatusNotImplemented)
}

// extractIDFromPath извлекает ID из пути URL
func extractIDFromPath(path, prefix string) string {
	cleaned := strings.TrimPrefix(path, prefix)
	segments := strings.Split(cleaned, "/")
	if len(segments) > 0 {
		return segments[0]
	}
	return ""
}
