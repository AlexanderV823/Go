package http

import (
	"log"
	"net/http"
	"time"
)

// Logger — мидлварь для базового логирования метода, пути и времени выполнения запроса
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Передаем управление следующему обработчику в цепочке
		next.ServeHTTP(w, r)

		// Логируем параметры запроса после его выполнения
		log.Printf("[%s] %s | Время выполнения: %v", r.Method, r.URL.Path, time.Since(start))
	})
}

// JsonContentType — мидлварь, которая автоматически добавляет заголовок JSON ко всем ответам
func JsonContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}