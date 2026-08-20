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
func (h *CommentHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var req model.CommentCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	comment, err := h.commentService.Create(r.Context(), userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			writeError(w, "Post not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeError(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(comment)
}

// GetByID возвращает комментарий по ID
// GET /api/comments/{id}
func (h *CommentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := extractCommentIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	comment, err := h.commentService.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Comment not found", http.StatusNotFound)
			return
		}
		writeError(w, "Failed to get comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comment)
}

// GetByPost возвращает комментарии к посту
// GET /api/posts/{id}/comments?limit=20&offset=0
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := extractPostIDFromCommentsPath(r.URL.Path)
	postID, err := strconv.Atoi(idStr)
	if err != nil || postID <= 0 {
		writeError(w, "Invalid post ID", http.StatusBadRequest)
		return
	}

	query := r.URL.Query()
	limit, _ := strconv.Atoi(query.Get("limit"))
	if limit <= 0 {
		limit = 20
	}
	offset, _ := strconv.Atoi(query.Get("offset"))
	if offset < 0 {
		offset = 0
	}

	comments, total, err := h.commentService.GetByPost(r.Context(), postID, limit, offset)
	if err != nil {
		if errors.Is(err, service.ErrPostNotExists) {
			writeError(w, "Post not found", http.StatusNotFound)
			return
		}
		writeError(w, "Failed to get comments", http.StatusInternalServerError)
		return
	}

	type CommentsResponse struct {
		Comments []*model.Comment `json:"comments"`
		Total    int              `json:"total"`
		Limit    int              `json:"limit"`
		Offset   int              `json:"offset"`
		PostID   int              `json:"post_id"`
	}

	resp := CommentsResponse{
		Comments: comments,
		Total:    total,
		Limit:    limit,
		Offset:   offset,
		PostID:   postID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// Update обновляет комментарий
// PUT /api/comments/{id}
// Требует аутентификации, может обновить только автор
func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := extractCommentIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	var req model.CommentUpdateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	comment, err := h.commentService.Update(r.Context(), id, userID, &req)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Comment not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			writeError(w, "You can only update your own comments", http.StatusForbidden)
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeError(w, "Failed to update comment", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(comment)
}

// Delete удаляет комментарий
// DELETE /api/comments/{id}
// Требует аутентификации, может удалить только автор
func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	idStr := extractCommentIDFromPath(r.URL.Path, "/api/comments/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		writeError(w, "Invalid comment ID", http.StatusBadRequest)
		return
	}

	err = h.commentService.Delete(r.Context(), id, userID)
	if err != nil {
		if errors.Is(err, service.ErrCommentNotFound) {
			writeError(w, "Comment not found", http.StatusNotFound)
			return
		}
		if errors.Is(err, service.ErrForbidden) {
			writeError(w, "You can only delete your own comments", http.StatusForbidden)
			return
		}
		writeError(w, "Failed to delete comment", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// Вспомогательные функции парсинга путей и ответов

func extractCommentIDFromPath(path, prefix string) string {
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	idPart := strings.TrimPrefix(path, prefix)
	if idx := strings.Index(idPart, "/"); idx != -1 {
		idPart = idPart[:idx]
	}
	return idPart
}

func extractPostIDFromCommentsPath(path string) string {
	prefix := "/api/posts/"
	if !strings.HasPrefix(path, prefix) {
		return ""
	}
	idPart := strings.TrimPrefix(path, prefix)
	suffixIdx := strings.Index(idPart, "/comments")
	if suffixIdx == -1 {
		return ""
	}
	return idPart[:suffixIdx]
}
