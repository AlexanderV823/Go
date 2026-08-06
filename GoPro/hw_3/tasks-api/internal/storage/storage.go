package storage

import (
	"errors"
	"tasks-api/internal/models"
)

// ErrTaskNotFound — доменная ошибка для сценария, когда задача не существует
var ErrTaskNotFound = errors.New("task not found")

type Storage interface {
	List() []models.Task
	Create(task models.Task) (models.Task, error)
	Get(id int) (models.Task, bool)
	Update(id int, task models.Task) (models.Task, error)
	Delete(id int) error
}