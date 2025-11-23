package addEmployee

import (
	"errors"
	"fmt"
	"github.com/AlexanderV823/Go/tree/hw6/lesson6/hw_6/findEmployee"
)

// Добавляет структуру сотрудника в срез. Если сотрудник с таким ID уже есть в срезе, данные о нем перезаписываются.
// emp - структура с данными сотрудника.
// arr - срез для поиска и записи.
func addEmployee(emp Employee, arr []Employees) (err error) {
	
	if err != nil {
		return fmt.Errorf("Ошибка вызова addEmployee: %w", err)
	}

	// Сначала, нужно проверить нет ли уже такого сотрудника
	// Ищем сотрудника по ID
	findIndex, err := findEmployee(arr, emp.ID)
	
	if err != nil {
		return fmt.Errorf("Ошибка поиска: %w", err)
	}

	if findIndex != nil {
		// Если сотрудник найден, то обновляем данные
		arr[findIndex] = emp
	} else {
		// Если сотрудника с таким ID нет, то добавляем нового
		arr = append(arr, emp)
	}

	return nil

}
