package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"todo-app/internal/todo"
)

// MakePath - кроссплатформенно формирует путь для сохранения файла относительно домашнего каталога пользователя.
// name 	- имя файла,
// ext 		- расширение файла.
func MakePath(name string, ext string) (filePath string, err error) {

	// Находим домашний каталог
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}

	// Объединяем путь к домашнему каталогу и имя папки
	folderPath := filepath.Join(homeDir, "MyTasks")

	// Создаем папку, если она не существует.
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		return "", fmt.Errorf("ошибка создания файла: %w", err)
	}

	// Соединяем папку и имя файла в путь
	filePath = filepath.Join(folderPath, name+ext)

	// Получаем информацию о файле
	_, err = os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			// Если ошибка равна ошибке отсутствия файла, то значит создаем новый и записываем пустую структуру, иначе будет ошибка при попытке десериализировать в JSON
			var tasks []todo.Task
			switch ext {
			case ".json":
				err = SaveJSON(filePath, tasks)
				if err != nil {
					return "", fmt.Errorf("ошибка создания json-файла: %w", err)
				}
			case ".csv":
				err = SaveCSV(filePath, tasks)
				if err != nil {
					return "", fmt.Errorf("ошибка создания csv-файла: %w", err)
				}
			default:
				return "", fmt.Errorf("указано неподдерживаемое расширение")
			}
		} else {
			return "", fmt.Errorf("ошибка создания файла: %w", err) //На случай иных ошибок
		}
	}
	//Сообщаем об успешном создании (нахождении) файла
	return filePath, nil
}
