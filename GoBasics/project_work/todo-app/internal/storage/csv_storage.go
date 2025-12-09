package storage

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"todo-app/internal/todo"
)

// LoadCSV загружает данные о задачах из файла CSV.
// path - путь к файлу
func LoadCSV(path string) ([]todo.Task, error) {

	//tasksCSV - срез структур с задачами
	var tasksCSV []todo.Task

	// 1. Откройте файл через os.Open.
	fileCSV, err := os.Open(path)
	if err != nil {
		return tasksCSV, fmt.Errorf("ошибка чтения файла csv: %w", err)
	}
	defer fileCSV.Close()

	// 2. Считайте все записи через csv.NewReader(...).ReadAll().
	reader := csv.NewReader(fileCSV)
	reader.Comma = ','

	records, err := reader.ReadAll()
	if err != nil {
		return tasksCSV, fmt.Errorf("ошибка чтения файла csv: %w", err)
	}

	for i, row := range records {
		// 3. Пропустите заголовок (первая строка).
		if i == 0 {
			continue
		}
		// 4. Преобразуйте строки: strconv.Atoi для ID, strconv.ParseBool для Done.
		id, _ := strconv.Atoi(row[0])
		desc := row[1]
		done, _ := strconv.ParseBool(row[2])

		// 5. Сформируйте срез []todo.Task.
		newTask := todo.Task{ID: id, Description: desc, Done: done}
		tasksCSV = append(tasksCSV, newTask)
	}

	return tasksCSV, nil

}

// SaveCSV сохраняет данные о задачах в файл CSV.
// path - путь к файлу
func SaveCSV(path string, tasks []todo.Task) error {

	// 1. Создайте (или перезапишите) файл через os.Create.
	fileCSV, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("ошибка сохранения файла csv: %w", err)
	}
	defer fileCSV.Close()

	// 2. Инициализируйте csv.NewWriter и запишите заголовок []string{"ID","Description","Done"}.
	writer := csv.NewWriter(fileCSV)

	record := []string{"ID", "Description", "Done"}
	writer.Write(record) // Запись одной строки

	// 3. Для каждой задачи сформируйте строку []string{strconv.Itoa(ID), Description, strconv.FormatBool(Done)}.
	var data [][]string

	for _, task := range tasks {
		data = append(data, []string{strconv.Itoa(task.ID), task.Description, strconv.FormatBool(task.Done)})
	}
	writer.WriteAll(data)

	// 4. Не забудьте writer.Flush().
	writer.Flush()

	if err := writer.Error(); err != nil {
		return fmt.Errorf("ошибка сохранения файла csv: %w", err)
	}
	return nil
}
