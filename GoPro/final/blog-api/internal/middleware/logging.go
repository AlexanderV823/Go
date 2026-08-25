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

// Ключ для контекста, чтобы безопасно хранить Request ID
type ctxKey string

const RequestIDKey ctxKey = "requestID"

// LoggingMiddleware предоставляет логирование запросов, CORS, восстановление и другие утилиты
type LoggingMiddleware struct {
	logger *log.Logger
}

// NewLoggingMiddleware создает новый экземпляр logging middleware
func NewLoggingMiddleware(logger *log.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

// Logger логирует все HTTP запросы с учетом Request ID
func (m *LoggingMiddleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapped := newResponseWriter(w)

		// Извлекаем requestID для красивого логирования (если он был добавлен ранее в цепочке)
		reqID, _ := r.Context().Value(RequestIDKey).(string)
		if reqID == "" {
			reqID = "-"
		}

		next.ServeHTTP(wrapped, r)

		duration := time.Since(start)
		ip := getClientIP(r)

		m.logger.Printf(
			"[API] %s | %s | %-7s | %s | %d | %d bytes | %v",
			ip,
			reqID,
			r.Method,
			r.URL.Path,
			wrapped.statusCode,
			wrapped.size,
			duration,
		)
	})
}

// Recovery восстанавливается после паник, пишет стек вызовов в логи и возвращает 500 ошибку
func (m *LoggingMiddleware) Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// Получаем стек вызовов (stack trace)
				stack := debug.Stack()

				// Логируем критическую ошибку
				m.logger.Printf("[PANIC RECOVER] %v\nStack trace:\n%s", err, string(stack))

				// Возвращаем клиенту чистый JSON с 500 статусом
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(`{"error":"Internal Server Error","message":"A critical error occurred"}`))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

// CORS добавляет CORS заголовки и корректно отвечает на PREFLIGHT (OPTIONS) запросы
func (m *LoggingMiddleware) CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем запросы с любых доменов (для продакшена лучше указать конкретные домены)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Request-ID")
		w.Header().Set("Access-Control-Max-Age", "86400") // Кэширование CORS-ответа на 24 часа

		// Если это Preflight-запрос (OPTIONS), завершаем его со статусом 204 No Content
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// RequestID генерирует уникальный ID для каждого запроса, прокидывает его в контекст и заголовок ответа
func (m *LoggingMiddleware) RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Пытаемся взять ID из входящего заголовка (если запрос идет через микросервисы)
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			// Если заголовка нет, генерируем быстрый криптографически стойкий псевдо-UUID
			b := make([]byte, 16)
			_, err := rand.Read(b)
			if err == nil {
				reqID = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
			} else {
				// Фолбэк на случай непредвиденных проблем с энтропией
				reqID = fmt.Sprintf("%d", time.Now().UnixNano())
			}
		}

		// Добавляем ID в заголовок ответа, чтобы клиент мог на него сослаться при ошибке
		w.Header().Set("X-Request-ID", reqID)

		// Помещаем ID в контекст запроса
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ipRateInfo хранит данные о лимитах конкретного IP адреса
type ipRateInfo struct {
	tokens     int
	lastUpdate time.Time
}

// RateLimiter реализует алгоритм Token Bucket для защиты от DDoS и брутфорса
func (m *LoggingMiddleware) RateLimiter(maxRequests int, window time.Duration) func(http.Handler) http.Handler {
	// sync.Map идеально подходит для конкурентного чтения/записи клиентов по IP
	var clients sync.Map

	// Скорость восполнения токенов (сколько токенов добавляется за секунду)
	refillRate := float64(maxRequests) / window.Seconds()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := getClientIP(r)
			now := time.Now()

			// Загружаем или инициализируем данные для текущего IP
			val, _ := clients.LoadOrStore(ip, &ipRateInfo{
				tokens:     maxRequests,
				lastUpdate: now,
			})
			info := val.(*ipRateInfo)

			// Вычисляем, сколько токенов успело натечь с момента последнего запроса
			elapsed := now.Sub(info.lastUpdate).Seconds()
			generatedTokens := elapsed * refillRate

			if generatedTokens > 0 {
				info.tokens += int(generatedTokens)
				if info.tokens > maxRequests {
					info.tokens = maxRequests
				}
				info.lastUpdate = now
			}

			// Если токенов нет — возвращаем 429 Too Many Requests
			if info.tokens <= 0 {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusTooManyRequests)
				_, _ = w.Write([]byte(`{"error":"Too Many Requests","message":"Rate limit exceeded. Please try again later."}`))
				return
			}

			// Списываем один токен и пропускаем запрос дальше
			info.tokens--
			next.ServeHTTP(w, r)
		})
	}
}

// ContentTypeJSON устанавливает Content-Type: application/json для всех ответов
func (m *LoggingMiddleware) ContentTypeJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

// getClientIP извлекает реальный IP адрес клиента
func getClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		parts := strings.Split(xff, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return xri
	}
	// Отсекаем порт (:8080), если он присутствует в RemoteAddr
	if idx := strings.LastIndex(r.RemoteAddr, ":"); idx != -1 {
		return r.RemoteAddr[:idx]
	}
	return r.RemoteAddr
}

// responseWriter обертка для захвата статус кода и размера ответа
type responseWriter struct {
	http.ResponseWriter
	statusCode int
	size       int
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
	n, err := rw.ResponseWriter.Write(b)
	rw.size += n
	return n, err
}

// newResponseWriter создает новую обертку
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		size:           0,
		written:        false,
	}
}
