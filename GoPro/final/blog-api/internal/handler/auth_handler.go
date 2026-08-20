package handler

import (
	"blog-api/internal/model"
	"blog-api/internal/service"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

type AuthHandler struct {
	userService *service.UserService
}

func NewAuthHandler(userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		userService: userService,
	}
}

// Register обрабатывает запрос на регистрацию нового пользователя
// POST /api/register
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.UserCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.userService.Register(r.Context(), &req)
	if err != nil {
		if errors.Is(err, service.ErrUserAlreadyExists) {
			writeError(w, "User already exists", http.StatusConflict)
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeError(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(resp)
}

// Login обрабатывает запрос на вход пользователя
// POST /api/login
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req model.UserLoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.userService.Login(r.Context(), &req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidCredentials) {
			writeError(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}
		if errors.Is(err, service.ErrInvalidInput) {
			writeError(w, err.Error(), http.StatusBadRequest)
			return
		}
		writeError(w, "Failed to login", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

// GetProfile возвращает профиль текущего пользователя
// GET /api/profile
func (h *AuthHandler) GetProfile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	userID, ok := getUserIDFromContext(r.Context())
	if !ok {
		writeError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := h.userService.GetByID(r.Context(), userID)
	if err != nil {
		if errors.Is(err, service.ErrUserNotFound) {
			writeError(w, "User not found", http.StatusNotFound)
			return
		}
		writeError(w, "Failed to get profile", http.StatusInternalServerError)
		return
	}

	// Формируем безопасный ответ без хэша пароля
	userResp := model.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userResp)
}

func writeError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// getUserIDFromContext извлекает ID пользователя из контекста
func getUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value("userID").(int)
	return userID, ok
}
