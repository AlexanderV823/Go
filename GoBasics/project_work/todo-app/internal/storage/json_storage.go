package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo-app/internal/todo"
)

// LoadJSON загружает данные о задачах из файла JSON.
// path - путь к файлу
func LoadJSON(path string) ([]todo.Task, error) {

	//tasksJSON - срез структур с задачами
	var tasksJSON []todo.Task

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {

			// Если ошибка равна ошибке отсутствия файла, то значит создаем новый
			err := SaveJSON(path, tasksJSON)

			if err != nil {
				return tasksJSON, fmt.Errorf("ошибка создания файла json: %w", err)
			}
		} else {
			return tasksJSON, fmt.Errorf("ошибка создания файла json: %w", err) //На случай иных ошибок
		}
	} else {
		if fileInfo.Size() == 0 {

			// Если в файле пусто, то просто записываем свои данные
			err := SaveJSON(path, tasksJSON)

			if err != nil {
				return tasksJSON, fmt.Errorf("ошибка создания файла json: %w", err)
			} else {

				//Сообщаем об успешной записи
				fmt.Println("Файл json успешно записан")
			}
		} else {
			// Если что-то записано, то прочитать из файла и загрузить в срез tasksJSON и вернуть в место вызова
			// Считайте содержимое os.ReadFile.
			// Открываем JSON-файл
			jsonData, err := os.ReadFile(path)

			if err != nil {
				return tasksJSON, fmt.Errorf("ошибка чтения файла json: %w", err)
			}

			// Десериализуем JSON-данные в структуру через json.Unmarshal
			err = json.Unmarshal(jsonData, &tasksJSON)

			if err != nil {
				// В случае наличия каких-то других данных, сообщить пользователю
				return tasksJSON, fmt.Errorf("ошибка чтения данных из файла json: %w", err)
			}

			//если все нормально, пробуем записать
			err = SaveJSON(path, tasksJSON)

			if err != nil {
				return tasksJSON, fmt.Errorf("ошибка создания файла json: %w", err)
			}
		}
	}
	return tasksJSON, nil
}

// SaveJSON сохраняет данные о задачах в файл JSON.
// path - путь к файлу,
// tasks - срез структур с задачами
func SaveJSON(path string, tasks []todo.Task) error {

	//Форматируем JSON в срез байтов
	tasksData, err := json.MarshalIndent(tasks, "", "\t")

	if err != nil {
		return fmt.Errorf("ошибка записи в файл json: %w", err)
	}

	//Записываем полученный срез байтов в файл
	err = os.WriteFile(path, tasksData, 0644)

	if err != nil {
		return fmt.Errorf("ошибка записи в файл json: %w", err)
	}

	//Сообщаем об успешной записи
	//fmt.Println("Данные успешно записаны в json файл")

	return nil
}
