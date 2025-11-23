package findEmployee

import (
	"fmt"
)

// Производит поиск сотрудника по ID в переданном срезе или массиве.
// array - срез для поиска,
// id - id-сотрудника.
// Возврашает nil, error, или структуру с данными сотрудника.
func findEmployee(array []byte, id int) (Employee, error) {

	var foundEmployee *Employee // Используйте указатель для возможности вернуть nil

	for i := range array {
		if array[i].ID == id {
			foundEmployee = &array[i]
			break
		}
	}

	if foundEmployee != nil {
		return foundEmployee
	}
	return nil, fmt.Errorf("сотрудник не найден")
}
