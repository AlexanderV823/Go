package memory

import (
	"errors"
	"sort"
	"sync"
	"time"

	"tasks-api/internal/models"
)

type MemoryStorage struct {
	mu     sync.RWMutex
	tasks  map[int]models.Task
	nextID int
}

func New() *MemoryStorage {
	return &MemoryStorage{
		tasks:  make(map[int]models.Task),
		nextID: 1,
	}
}

// List возвращает копию списка задач, гарантированно отсортированную по ID
func (s *MemoryStorage) List() []models.Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	list := make([]models.Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		// Создаем явную копию структуры, чтобы избежать передачи ссылок
		taskCopy := models.Task{
			ID:        task.ID,
			Title:     task.Title,
			Done:      task.Done,
			CreatedAt: task.CreatedAt,
		}
		list = append(list, taskCopy)
	}

	// Гарантируем стабильный порядок обхода (сортировка по ID возрастанию)
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})

	return list
}

func (s *MemoryStorage) Create(t models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t.ID = s.nextID
	s.nextID++
	if t.CreatedAt == "" {
		t.CreatedAt = time.Now().Format(time.RFC3339)
	}

	// Сохраняем копию
	s.tasks[t.ID] = models.Task{
		ID:        t.ID,
		Title:     t.Title,
		Done:      t.Done,
		CreatedAt: t.CreatedAt,
	}

	// Возвращаем копию наружу
	return s.tasks[t.ID], nil
}

func (s *MemoryStorage) Get(id int) (models.Task, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, false
	}

	// Возвращаем изолированную копию структуры
	return models.Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
	}, true
}

func (s *MemoryStorage) Update(id int, t models.Task) (models.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	task, exists := s.tasks[id]
	if !exists {
		return models.Task{}, errors.New("task not found")
	}

	// Модифицируем внутреннее состояние в рамках блокировки
	task.Title = t.Title
	task.Done = t.Done
	s.tasks[id] = task

	// Возвращаем изолированную копию
	return models.Task{
		ID:        task.ID,
		Title:     task.Title,
		Done:      task.Done,
		CreatedAt: task.CreatedAt,
	}, nil
}

func (s *MemoryStorage) Delete(id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(s.tasks, id)
	return nil
}