package findEmployee

import (
	"fmt"
)

// Производит поиск сотрудника по ID в переданном срезе или массиве.
// array - срез для поиска,
// id - id-сотрудника.
// Возврашает индекс записи о сотруднике или error.
func findEmployee(array []byte, id int) (index int, err error) {

	for i := range array {
		if array[i].ID == id {
			index = i
			break
		}
	}

	if index != nil {
		return index
	}

	return fmt.Errorf("сотрудник не найден")

}
