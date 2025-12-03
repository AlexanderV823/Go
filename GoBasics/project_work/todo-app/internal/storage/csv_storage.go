package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"todo-app/internal/todo"
)

// LoadCSV загружает данные о задачах из файла CSV
// path - путь к файлу
func LoadCSV(path string) ([]todo.Task, error) {

}

// SaveCSV сохраняет данные о задачах в файл CSV
// path - путь к файлу
func SaveCSV(path string, tasks []todo.Task) error {

}
