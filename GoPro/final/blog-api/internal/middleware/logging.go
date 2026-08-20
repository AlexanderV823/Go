package middleware

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

const (
	// RequestIDKey используется для сохранения ID запроса в контексте
	RequestIDKey contextKey = "requestID"
)

// LoggingMiddleware provides request logging, CORS, recovery and other utility middleware
type LoggingMiddleware struct {
	logger *log.Logger

	// Хранилище для Rate Limiting
	mu       sync.Mutex
	visitors map[string][]time.Time
}

// NewLoggingMiddleware creates a new logging middleware instance
func NewLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger:   logger,
		visitors: make(map[string][]time.Time),
	}
}

// Logger logs all HTTP requests
func (m *LoggingMiddleware) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Оборачиваем стандартный ResponseWriter
		rw := newResponseWriter(w)

		// Извлекаем Request ID из контекста, если он там есть
		reqID, _ := r.Context().Value(RequestIDKey).(string)
		if reqID != "" {
			reqID = "[" + reqID + "] "
		}

		next(rw, r)

		// Логируем по окончании выполнения запроса
		m.logger.Printf("%s%s %s %s - %d %s",
			reqID,
			r.Method,
			r.URL.Path,
			getClientIP(r),
			rw.statusCode,
			time.Since(start),
		)
	}
}

// Recovery восстанавливается после паник
func (m *LoggingMiddleware) Recovery(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Захватываем stack trace для отладки
				stack := debug.Stack()
				m.logger.Printf("[PANIC RECOVERY] Error: %v\nStack Trace:\n%s", err, string(stack))

				// Отдаем клиенту 500 ошибку
				writeJSONError(w, "internal server error", http.StatusInternalServerError)
			}
		}()

		next(w, r)
	}
}

// CORS добавляет CORS заголовки
func (m *LoggingMiddleware) CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем базовые CORS заголовки
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
		w.Header().Set("Access-Control-Max-Age", "86400") // 24 часа кэширования префлайта

		// Если это Preflight запрос — завершаем его со статусом 204
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}

// RequestID добавляет уникальный ID к каждому запросу
func (m *LoggingMiddleware) RequestID(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Генерируем псевдо-UUID или простой случайный хэш
		b := make([]byte, 8)
		_, err := rand.Read(b)
		reqID := fmt.Sprintf("%x-%d", b, time.Now().UnixNano())
		if err != nil {
			reqID = fmt.Sprintf("%d", time.Now().UnixNano())
		}

		// Добавляем ID в контекст запроса и в заголовок ответа
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)
		w.Header().Set("X-Request-ID", reqID)

		next(w, r.WithContext(ctx))
	}
}

// RateLimiter ограничивает количество запросов от одного клиента
func (m *LoggingMiddleware) RateLimiter(maxRequests int, window time.Duration) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			ip := getClientIP(r)
			now := time.Now()

			m.mu.Lock()

			// Ленивая очистка старых таймстампов для текущего IP
			var validRequests []time.Time
			for _, t := range m.visitors[ip] {
				if now.Sub(t) < window {
					validRequests = append(validRequests, t)
				}
			}

			// Проверяем лимиты
			if len(validRequests) >= maxRequests {
				m.mu.Unlock()
				writeJSONError(w, "too many requests", http.StatusTooManyRequests)
				return
			}

			// Добавляем текущий запрос в историю
			validRequests = append(validRequests, now)
			m.visitors[ip] = validRequests
			m.mu.Unlock()

			next(w, r)
		}
	}
}

// ContentTypeJSON устанавливает Content-Type: application/json для всех ответов
func (m *LoggingMiddleware) ContentTypeJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// getClientIP извлекает IP адрес клиента
func getClientIP(r *http.Request) string {
	// Проверяем заголовок прокси X-Forwarded-For
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		ips := strings.Split(xff, ",")
		if len(ips) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// Проверяем Nginx заголовок X-Real-IP
	xri := r.Header.Get("X-Real-IP")
	if xri != "" {
		return strings.TrimSpace(xri)
	}

	// Если заголовков нет, берем RemoteAddr и отсекаем порт
	ip := r.RemoteAddr
	if idx := strings.LastIndex(ip, ":"); idx != -1 {
		ip = ip[:idx]
	}
	return ip
}

// responseWriter обертка для захвата статус кода
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	written    bool
}

// WriteHeader сохраняет статус код
func (rw *responseWriter) WriteHeader(code int) {
	if !rw.written {
		rw.statusCode = code
		rw.ResponseWriter.WriteHeader(code)
		rw.written = true
	}
}

// Write вызывает WriteHeader если еще не был вызван
func (rw *responseWriter) Write(b []byte) (int, error) {
	if !rw.written {
		rw.WriteHeader(http.StatusOK)
	}
	return rw.ResponseWriter.Write(b)
}

// newResponseWriter создает новую обертку
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		written:        false,
	}
}
