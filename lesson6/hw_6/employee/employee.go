package employee

import "fmt"

type Displayable interface {
	Display()
}

// Employee - структура сотрудника
type Employee struct {
	ID      int     //Табельный номер
	Surname string  //Фамилия
	Name    string  //Имя
	Age     int     //Возраст
	Job     string  //Должность
	Salary  float64 //Зарплата
}

func (emp Employee) Display() { //В параметре ждать данные для вывода
	fmt.Printf("Таб.номер: %d\nФамилия: %s\nИмя: %s\nВозраст: %d\nДолжность: %s\nОклад: %.2f руб.\n", emp.ID, emp.Surname, emp.Name, emp.Age, emp.Job, emp.Salary)
}

// Employees - срез сотрудников
var Employees = make([]Employee, 3)

// Add добавляет структуру сотрудника в срез. Если сотрудник с таким ID уже есть в срезе, данные о нем перезаписываются.
// emp - структура с данными сотрудника.
func Add(emp Employee) (err error) {
	// Сначала, нужно проверить нет ли уже такого сотрудника
	// Ищем сотрудника по ID
	findIndex, err := Find(Employees, emp.ID)

	if err != nil {
		return fmt.Errorf("ошибка поиска: %w", err)
	}

	if findIndex != -1 {
		// Если сотрудник найден, то обновляем данные
		Employees[findIndex] = emp
	} else {
		// Если сотрудника с таким ID нет, то добавляем нового
		Employees = append(Employees, emp)
	}
	return nil
}

// Find производит поиск сотрудника по ID в переданном срезе или массиве.
// slice - срез для поиска,
// id - табельный номер сотрудника.
// Возврашает индекс записи о сотруднике или -1 если такой записи нет.
func Find(slice []Employee, id int) (int, error) {

	for i := range slice {
		emp := slice[i]
		if emp.ID == id {
			return i, nil
		}
	}

	return -1, nil

}

// FilterAge выводит список сотрудников старше или равных, указанному возрасту
// minAge - возраст для фильтрации
func FilterAge(minAge int) (err error) {
	var filteredEmp []Employee
	for _, emp := range Employees {
		if emp.Age >= minAge {
			filteredEmp = append(filteredEmp, emp)
		}
	}
	for _, emp := range filteredEmp {
		emp.Display()
	}
	return nil
}

// FilterSalary выводит список сотрудников с зарплатой выше или равной, указанной
// minSalary - зарплата для фильтрации
func FilterSalary(minSalary float64) (err error) {
	var filteredEmp []Employee
	for _, emp := range Employees {
		if emp.Salary >= minSalary {
			filteredEmp = append(filteredEmp, emp)
		}
	}
	for _, emp := range filteredEmp {
		emp.Display()
	}
	return nil
}
