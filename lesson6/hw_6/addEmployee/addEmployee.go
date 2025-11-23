package addEmployee

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/writeData"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/findEmployee"
	"io"
	"os"
	"path/filepath"
)

var (
	ErrJSONRead = errors.New("не удалось прочитать JSON")
)

func addEmployee(emp Employee) (err error) {

	if err != nil {
		return fmt.Errorf("Ошибка вызова addEmployee: %w", err)
	}

	// Преобразуем структуру emp в JSON
	jsonData, err := json.Marshal(emp)
	if err != nil {
		return fmt.Errorf("Ошибка формирования json: %w", err)
	}

	//Создаем нашу мини-БД для хранения сотрудников
	homeDir, err := os.UserHomeDir()

	if err != nil {
		return fmt.Errorf("Ошибка создания БД: %w", err)
	}

	// Объединяем путь к домашнему каталогу и имя папки
	folderPath := filepath.Join(homeDir, "employeesDB")

	// Создаем папку, если она не существует.
	err = os.MkdirAll(folderPath, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка создания БД: %w", err)
	}

	// Создание файла
	filePath := filepath.Join(folderPath, "employeesDB.txt")

	fileInfo, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			// Если ошибка равна ошибке отсутствия файла, то значит создаем новый
			err := writeData(filePath, jsonData)
			if err != nil {
				return fmt.Errorf("Ошибка создания БД: %w", err)
			}
		} else {
			return fmt.Errorf("Ошибка создания БД: %w", err) //На случай иных ошибок
		}
	} else {
		// Сначала открываем файл, только для чтения, чтобы проверить наличие в нем уже таких данных
		file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
		if err != nil {
			return fmt.Errorf("Ошибка чтения файла: %w", err)
		}
		defer file.Close() // Обязательно сразу объявляем закрытие файла

		if fileInfo.Size() == 0 {
			// Если в файле пусто, то просто записываем свои данные
			err := writeData(filePath, jsonData)
			if err != nil {
				return fmt.Errorf("Ошибка создания БД: %w", err)
			}
		} else {
			// Если что=то записано, то проверить содержимой файла на соответствие нашей структуре Employee
			// Открываем JSON-файл
			jsonFile, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
			if err != nil {
				return fmt.Errorf("Ошибка чтения файла: %w", err)
			}
			defer jsonFile.Close()

			// Считываем содержимое файла
			byteValue, err := io.ReadAll(jsonFile)
			if err != nil {
				return fmt.Errorf("Ошибка чтения файла: %w", err)
			}

			// Создаем переменную структуры
			var e Employee

			// Десериализуем JSON-данные в структуру
			err = json.Unmarshal(byteValue, &e)
			if err != nil {
				// В случае наличия каких-то других данных, сообщить пользователю
				return fmt.Errorf("Ошибка чтения данных из файла: %w (%w)", ErrJSONRead, err)
			}
			
			// Выводим данные из структуры.
			fmt.Printf("Таб.номер: %d, Фамилия: %s, Имя: %s, Должность:%s Возраст: %d\n",
				e.ID, e.Surname, e.Name, e.Job.Name, e.Age)
			
			// Если данные соответствуют нашей структуре, то проверять, нет ли уже такого сотрудника
			_, err := findEmployee(byteValue, e.ID)
			
			// Ищем сотрудника по ID, если найден, то обновляем данные
			// создать универсальную функцию поиска в массиве Find(array, value) (element, error)

		}

	}

	//Сообщаем об успешной записи
	fmt.Println("Файл успешно записан")
	return nil

}
