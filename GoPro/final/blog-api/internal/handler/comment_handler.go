package handler

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"blog-api/internal/middleware"
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

// CreateComment обрабатывает создание нового комментария
// POST /api/posts/{id}/comments
func (h *CommentHandler) CreateComment(w http.ResponseWriter, r *http.Request) {
	postIDStr := chi.URLParam(r, "id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		writeError(w, "invalid post id", http.StatusBadRequest)
		return
	}

	var req model.CommentCreateRequest
	// Лимит 256 КБ на комментарий
	if err := decodeJSONStrict(w, r, &req, 256<<10); err != nil {
		writeError(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	// Валидация содержимого в рунах
	req.Content, ok = cleanAndValidateString(req.Content, 1, 2000)
	if !ok {
		writeError(w, "Comment content must be between 1 and 2000 characters and cannot consist of spaces", http.StatusBadRequest)
		return
	}

	resp, err := h.commentService.Create(r.Context(), postID, &req, userID)
	if err != nil {
		writeError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// GetByID возвращает комментарий по ID
func (h *CommentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
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
func (h *CommentHandler) GetByPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
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
func (h *CommentHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
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
	if err := decodeJSONStrict(w, r, &req, 256<<10); err != nil {
		writeError(w, "Malformed validation matrix signature: "+err.Error(), http.StatusBadRequest)
		return
	}

	req.Content, ok = cleanAndValidateString(req.Content, 1, 2000)
	if !ok {
		writeError(w, "Comment content must be between 1 and 2000 characters and cannot consist of spaces", http.StatusBadRequest)
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
func (h *CommentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := middleware.GetUserIDFromContext(r.Context())
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
