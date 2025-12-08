package storage

import (
	"fmt"
	// "os"
	// "path/filepath"
	"todo-app/internal/todo"
)

// LoadCSV загружает данные о задачах из файла CSV.
// path - путь к файлу
func LoadCSV(path string) ([]todo.Task, error) {
	fmt.Println("Загрузка из CSV в разработке...")
	return nil, nil
}

// SaveCSV сохраняет данные о задачах в файл CSV.
// path - путь к файлу
func SaveCSV(path string, tasks []todo.Task) error {
	fmt.Println("Сохранение в CSV в разработке...")
	return nil
}
