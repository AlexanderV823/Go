package middleware

import (
	"blog-api/pkg/auth"
	"context"
	"encoding/json"
	"net/http"
	"strings"
)

// contextKey is a custom type for context keys to avoid collisions
type contextKey string

const (
	// UserIDKey is the key for storing user ID in context
	UserIDKey contextKey = "userID"
	// UserEmailKey is the key for storing user email in context
	UserEmailKey contextKey = "userEmail"
	// UserNameKey is the key for storing username in context
	UserNameKey contextKey = "username"
)

// AuthMiddleware provides JWT authentication
type AuthMiddleware struct {
	jwtManager *auth.JWTManager
}

// NewAuthMiddleware creates a new auth middleware instance
func NewAuthMiddleware(jwtManager *auth.JWTManager) *AuthMiddleware {
	return &AuthMiddleware{
		jwtManager: jwtManager,
	}
}

// RequireAuth is a middleware that requires valid JWT token
func (m *AuthMiddleware) RequireAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := extractToken(r)
		if tokenStr == "" {
			writeJSONError(w, "missing or invalid authorization header", http.StatusUnauthorized)
			return
		}

		// Используем ваш родной метод ValidateToken из пакета auth
		claims, err := m.jwtManager.ValidateToken(tokenStr)
		if err != nil {
			writeJSONError(w, "invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Добавляем данные пользователя в контекст с использованием кастомных типов ключей
		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		ctx = context.WithValue(ctx, UserNameKey, claims.Username)

		// Сохраняем обратную совместимость с хендлерами, которые используют строковый ключ "userID"
		ctx = context.WithValue(ctx, "userID", claims.UserID)

		next(w, r.WithContext(ctx))
	}
}

// OptionalAuth is a middleware that extracts JWT token if present, but doesn't require it
func (m *AuthMiddleware) OptionalAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr := extractToken(r)
		if tokenStr == "" {
			next(w, r)
			return
		}

		claims, err := m.jwtManager.ValidateToken(tokenStr)
		if err != nil {
			// Если токен передан, но он невалидный — продолжаем работу как анонимный пользователь
			next(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		ctx = context.WithValue(ctx, UserNameKey, claims.Username)
		ctx = context.WithValue(ctx, "userID", claims.UserID)

		next(w, r.WithContext(ctx))
	}
}

// extractToken извлекает JWT токен из заголовка Authorization
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Заголовок должен строго начинаться с "Bearer "
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || !strings.EqualFold(parts[0], "Bearer") {
		return ""
	}

	return parts[1]
}

// GetUserIDFromContext извлекает ID пользователя из контекста
func GetUserIDFromContext(ctx context.Context) (int, bool) {
	userID, ok := ctx.Value(UserIDKey).(int)
	return userID, ok
}

// GetUserEmailFromContext извлекает email пользователя из контекста
func GetUserEmailFromContext(ctx context.Context) (string, bool) {
	email, ok := ctx.Value(UserEmailKey).(string)
	return email, ok
}

// GetUsernameFromContext извлекает username из контекста
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UserNameKey).(string)
	return username, ok
}

// writeJSONError отправляет ошибку в формате JSON
func writeJSONError(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// Chain позволяет объединить несколько middleware в цепочку
func Chain(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
