package findEmp

import (
	"fmt"
)

// FindEmployee производит поиск сотрудника по ID в переданном срезе или массиве.
// array - срез для поиска,
// id - id-сотрудника.
// Возврашает индекс записи о сотруднике или -1 и error.
func FindEmployee(slice []Employee, id int) (int, error) {

	for i := range slice {
		emp := slice[i]
		if emp.ID == id {
			return i, nil
		}
	}

return -1, fmt.Errorf("сотрудник не найден")

}