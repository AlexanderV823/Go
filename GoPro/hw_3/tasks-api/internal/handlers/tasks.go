package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"tasks-api/internal/models"
	"tasks-api/internal/storage"
)

// Handler инкапсулирует слой бизнес-логики и работу с хранилищем
type Handler struct {
	Store storage.Storage
}

// New создает новый экземпляр обработчика задач
func New(s storage.Storage) *Handler {
	return &Handler{Store: s}
}

// Вспомогательный метод для отправки JSON-ошибок клиенту
func (h *Handler) sendError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

// Вспомогательный метод для отправки успешных ответов в формате JSON
func (h *Handler) sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// Health — Эндпоинт GET /health для проверки работоспособности сервиса
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	h.sendJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

// TasksCollection обрабатывает запросы к корню коллекции: /tasks (GET, POST)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		// Возвращает детерминированный, отсортированный по ID срез-копию задач
		tasks := h.Store.List()
		h.sendJSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		// Закрываем тело запроса сразу после выхода из этой ветки для экономии ресурсов
		defer r.Body.Close()

		var req models.Task
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.sendError(w, http.StatusBadRequest, "Неверный формат JSON")
			return
		}

		req.Title = strings.TrimSpace(req.Title)

		if req.Title == "" {
			h.sendError(w, http.StatusBadRequest, "Поле 'title' обязательно")
			return
		}

		created, err := h.Store.Create(req)
		if err != nil {
			log.Printf("Непредвиденная ошибка при создании задачи: %v", err)
			h.sendError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			return
		}

		h.sendJSON(w, http.StatusCreated, created)

	default:
		// REST-стандарт: заголовок Allow обязателен для статуса 405
		w.Header().Set("Allow", "GET, POST")
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// TaskItem обрабатывает операции с конкретной задачей по ID: /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {

	// Извлекаем и валидируем числовой ID из пути URL
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		h.sendError(w, http.StatusBadRequest, "Неверный или отсутствующий ID задачи")
		return
	}

	switch r.Method {
	case http.MethodGet:
		task, exists := h.Store.Get(id)
		if !exists {
			h.sendError(w, http.StatusNotFound, "Задача не найдена")
			return
		}
		h.sendJSON(w, http.StatusOK, task)

	case http.MethodPut:
		defer r.Body.Close()

		var req models.Task
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.sendError(w, http.StatusBadRequest, "Неверный формат JSON")
			return
		}

		req.Title = strings.TrimSpace(req.Title)

		if req.Title == "" {
			h.sendError(w, http.StatusBadRequest, "Поле 'title' обязательно")
			return
		}

		updated, err := h.Store.Update(id, req)
		if err != nil {
			// Дифференцируем ожидаемую доменную ошибку и непредвиденный сбой сервера
			if errors.Is(err, storage.ErrTaskNotFound) {
				h.sendError(w, http.StatusNotFound, "Задача не найдена")
			} else {
				log.Printf("Критическая ошибка при обновлении задачи %d: %v", id, err)
				h.sendError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			}
			return
		}
		h.sendJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		err := h.Store.Delete(id)
		if err != nil {
			// Дифференцируем ошибки для DELETE ресурса
			if errors.Is(err, storage.ErrTaskNotFound) {
				h.sendError(w, http.StatusNotFound, "Задача не найдена")
			} else {
				log.Printf("Критическая ошибка при удалении задачи %d: %v", id, err)
				h.sendError(w, http.StatusInternalServerError, "Внутренняя ошибка сервера")
			}
			return
		}
		// Статус 204 No Content возвращается без какого-либо тела JSON
		w.WriteHeader(http.StatusNoContent)

	default:
		// Список разрешенных методов для эндпоинта элемента коллекции
		w.Header().Set("Allow", "GET, PUT, DELETE")
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}