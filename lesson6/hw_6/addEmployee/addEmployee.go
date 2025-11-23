package addEmployee

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/findEmployee"
	"io"
	"os"
	"path/filepath"
)

// Добавляет структуру сотрудника в срез. Если сотрудник с таким ID уже есть в срезе, данные о нем перезаписываются.
// emp - структура с данными сотрудника.
// arr - срез для поиска и записи.
func addEmployee(emp Employee, arr []Employees) (err error) {

	if err != nil {
		return fmt.Errorf("Ошибка вызова addEmployee: %w", err)
	}

	// Сначала, нужно проверить нет ли уже такого сотрудника
	findEmployee, err := findEmployee(arr, emp.ID)
	if err != nil {
		return fmt.Errorf("Ошибка поиска: %w", err)
	}
	fmt.Printf("Таб.номер: %d, Фамилия: %s, Имя: %s, Должность:%s Возраст: %d\n", foundEmployee.ID, foundEmployee.Surname, foundEmployee.Name, foundEmployee.Job.Name, foundEmployee.Age)
	// Ищем сотрудника по ID, если найден, то обновляем данные

	//Сообщаем об успешной записи
	fmt.Println("Файл успешно записан")
	return nil

}
