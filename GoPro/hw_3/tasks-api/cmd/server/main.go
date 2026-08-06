package main

import (
	"log"
	"net/http"

	"tasks-api/internal/handlers"
	"tasks-api/internal/storage/memory"
)

func main() {
	// Подключаем потокобезопасную in-memory реализацию интерфейса Storage
	store := memory.New()

	h := handlers.New(store)

	mux := http.NewServeMux()

	// РЕГИСТРАЦИЯ ЗДЕСЬ: Добавляем маршрут для проверки здоровья
	mux.HandleFunc("/health", h.Health)

	mux.HandleFunc("/tasks", h.TasksCollection) // GET, POST
	mux.HandleFunc("/tasks/", h.TaskItem)       // GET, PUT, DELETE

	log.Println("server listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}