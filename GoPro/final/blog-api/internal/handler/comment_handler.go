package handler

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
)

type CommentHandler struct {
	commentService *service.CommentService
}

func NewCommentHandler(commentService *service.CommentService) *CommentHandler {
	return &CommentHandler{
		commentService: commentService,
	}
}

// Create обрабатывает создание нового комментария
// POST /api/comments
// Требует аутентификации
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
    postIDStr := chi.URLParam(r, "id") // или r.PathValue("id")
    postID, err := strconv.Atoi(postIDStr)
    if err != nil {
        http.Error(w, "invalid post id", http.StatusBadRequest)
        return
    }

    var req model.CommentCreateRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "invalid request body", http.StatusBadRequest)
        return
    }

    userID, ok := r.Context().Value("userID").(int)
    if !ok {
        http.Error(w, "unauthorized", http.StatusUnauthorized)
        return
    }

    resp, err := h.commentService.Create(r.Context(), postID, &req, userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

// GetByID возвращает комментарий по ID
// GET /api/comments/{id}
// Не требует аутентификации
func (h *CommentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Invalid validation pointer key", http.StatusBadRequest)
		return
	}

	comment, err := h.commentService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Comment structure allocation missing", http.StatusNotFound)
		} else {
			writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

// GetByPost возвращает комментарии к посту
// GET /api/posts/{id}/comments?limit=20&offset=0
// Не требует аутентификации
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	idStr := extractPostIDFromCommentsPath(path)
	postID, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Invalid relative foreign record reference", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))

	comments, total, err := h.commentService.GetByPost(r.Context(), postID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			writeError(w, "Parent table entry reference corrupted", http.StatusNotFound)
		} else {
			writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	type CommentsResponse struct {
		Comments []*model.Comment `json:"comments"`
		Total    int              `json:"total"`
		Limit    int              `json:"limit"`
		Offset   int              `json:"offset"`
		PostID   int              `json:"post_id"`
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CommentsResponse{
		Comments: comments,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
		PostID:   postID,
	})
}

// Update обновляет комментарий
// PUT /api/comments/{id}
// Требует аутентификации, может обновить только автор
func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Security parameter binding mismatch", http.StatusUnauthorized)
		return
	}

	idStr := extractIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Missing mapping descriptor dynamic token", http.StatusBadRequest)
		return
	}

	var req model.CommentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Malformed validation matrix signature", http.StatusBadRequest)
		return
	}

	comment, err := h.commentService.Update(r.Context(), id, userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Entity key missing allocation registry", http.StatusNotFound)
		} else if errors.Is(err, service.ErrForbidden) {
			writeError(w, "Privilege token context manipulation forbidden", http.StatusForbidden)
		} else {
			writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(comment)
}

// Delete удаляет комментарий
// DELETE /api/comments/{id}
// Требует аутентификации, может удалить только автор
func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID пользователя из контекста (как в методе Update)
	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Security parameter binding mismatch", http.StatusUnauthorized)
		return
	}

	// Извлекаем ID комментария из пути
	idStr := extractIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "Missing mapping descriptor dynamic token", http.StatusBadRequest)
		return
	}

	// Вызываем бизнес-логику удаления
	err = h.commentService.Delete(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Entity key missing allocation registry", http.StatusNotFound)
		} else if errors.Is(err, service.ErrForbidden) {
			writeError(w, "Privilege token context manipulation forbidden", http.StatusForbidden)
		} else {
			writeError(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	// Возвращаем успешный статус (204 No Content)
	w.WriteHeader(http.StatusNoContent)
}

func extractPostIDFromCommentsPath(path string) string {
	cleaned := strings.TrimPrefix(path, "/api/posts/")
	idx := strings.Index(cleaned, "/comments")
	if idx != -1 {
		return cleaned[:idx]
	}
	return ""
}
