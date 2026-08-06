package main

import (
	"log"
	"net/http"

	"tasks-api/internal/handlers"
	customMiddleware "tasks-api/internal/http" // Импортируем мидлваре
	"tasks-api/internal/storage/memory"
)

func main() {
	store := memory.New()
	h := handlers.New(store)

	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/tasks", h.TasksCollection)
	mux.HandleFunc("/tasks/", h.TaskItem)

	// Оборачиваем mux в цепочку мидлварей (выполняются снизу вверх/снаружи внутрь)
	var handler http.Handler = mux
	handler = customMiddleware.JsonContentType(handler) // Авто-заголовок JSON
	handler = customMiddleware.Logger(handler)          // Базовое логирование

	log.Println("server listening on :8080")
	// Запускаем сервер с нашей мидлварей вместо mux
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}