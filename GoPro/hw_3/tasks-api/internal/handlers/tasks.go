package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"tasks-api/internal/models"
	"tasks-api/internal/storage"
)

type Handler struct{ Store storage.Storage }

func New(s storage.Storage) *Handler { return &Handler{Store: s} }

func (h *Handler) sendError(w http.ResponseWriter, status int, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{"error": msg})
}

func (h *Handler) sendJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

// PRO: Health-check с ограничением по методу и заголовком Allow
func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET") // Указываем разрешенный метод
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
		return
	}

	h.sendJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

// /tasks (GET, POST)
func (h *Handler) TasksCollection(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		tasks := h.Store.List()
		h.sendJSON(w, http.StatusOK, tasks)

	case http.MethodPost:
		var req models.Task
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.sendError(w, http.StatusBadRequest, "Неверный формат JSON")
			return
		}

		if strings.TrimSpace(req.Title) == "" {
			h.sendError(w, http.StatusBadRequest, "Поле 'title' обязательно")
			return
		}

		created, _ := h.Store.Create(req)
		h.sendJSON(w, http.StatusCreated, created)

	default:
		w.Header().Set("Allow", "GET, POST") // Указываем разрешенные методы
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}

// /tasks/{id} (GET, PUT, DELETE)
func (h *Handler) TaskItem(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)

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
		var req models.Task
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			h.sendError(w, http.StatusBadRequest, "Неверный формат JSON")
			return
		}

		if strings.TrimSpace(req.Title) == "" {
			h.sendError(w, http.StatusBadRequest, "Поле 'title' обязательно")
			return
		}

		updated, err := h.Store.Update(id, req)
		if err != nil {
			h.sendError(w, http.StatusNotFound, "Задача не найдена")
			return
		}
		h.sendJSON(w, http.StatusOK, updated)

	case http.MethodDelete:
		err := h.Store.Delete(id)
		if err != nil {
			h.sendError(w, http.StatusNotFound, "Задача не найдена")
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		w.Header().Set("Allow", "GET, PUT, DELETE") // Указываем разрешенные методы
		h.sendError(w, http.StatusMethodNotAllowed, "Метод не поддерживается")
	}
}