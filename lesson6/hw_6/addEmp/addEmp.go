package addEmp

import (
	"fmt"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/findEmp"
)

// Добавляет структуру сотрудника в срез. Если сотрудник с таким ID уже есть в срезе, данные о нем перезаписываются.
// emp - структура с данными сотрудника.
// slc - срез для поиска и записи.
func AddEmployee(emp Employee, slc []Employee) (err error) {

	if err != nil {
		return fmt.Errorf("Ошибка вызова addEmployee: %w", err)
	}

	// Сначала, нужно проверить нет ли уже такого сотрудника
	// Ищем сотрудника по ID
	findIndex, err := findEmp.FindEmployee(slc, emp.ID)

	if err != nil {
		return fmt.Errorf("Ошибка поиска: %w", err)
	}

	if findIndex != -1 {
		// Если сотрудник найден, то обновляем данные
		slc[findIndex] = emp
	} else {
		// Если сотрудника с таким ID нет, то добавляем нового
		slc = append(slc, emp)
	}

	return nil

}
