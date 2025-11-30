package main

import (
	"fmt"
	"hw_6/employee"
	"os"
)

func main() {

	for {
		var input string
		fmt.Print("Введите команду: ")
		_, err := fmt.Scan(&input)
		if err != nil {
			if err.Error() == "unexpected newline" {
				continue //если введен Enter, то выводить сообщение об ощибке не обязательно
			}
			fmt.Printf("Ошибка ввода: %s", err)
			continue
		}
		switch input {
		case "a", "add":

			var (
				id      int
				surname string
				name    string
				age     int
				job     string
				salary  float64
			)

			fmt.Print("Введите таб.номер, фамилию, имя, возраст, должность и зарплату: ")
			_, err := fmt.Scanf("%d %s %s %d %s %f", &id, &surname, &name, &age, &job, &salary)

			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s", err)
				continue
			}

			newEmp := employee.Employee{ID: id, Surname: surname, Name: name, Age: age, Job: job, Salary: salary}
			err = employee.Add(newEmp)

			if err != nil {
				fmt.Printf("Ошибка добавления сотрудника: %s", err)
			} else {
				//Сообщаем об успешной записи
				fmt.Println("Данные о сотруднике записаны")
			}
		case "d", "disp":
			var id int
			fmt.Print("Введите табельный номер: ")
			_, err := fmt.Scanf("%d", &id)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s", err)
			}
			id, err = employee.Find(employee.Employees, id)
			if err != nil {
				fmt.Printf("Ошибка поиска сотрудника: %s", err)
			}
			if id == -1 {
				fmt.Println("Сотрудник с таким табельным номером не найден")
			}
			emp := employee.Employees[id]
			emp.Display()
		case "f", "find":
			var (
				minAge    int
				minSalary float64
			)
			fmt.Print("Введите минимальный возраст: ")
			_, err := fmt.Scanf("%d", &minAge)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s", err)

			}
			err = employee.FilterAge(minAge)
			if err != nil {
				fmt.Printf("Ошибка фильтрации по возрасту: %s", err)
			}
			fmt.Print("Введите минимальную зарплату: ")
			_, err = fmt.Scanf("%d", &minSalary)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s", err)
				continue
			}
			err = employee.FilterSalary(minSalary)
			if err != nil {
				fmt.Printf("Ошибка фильтрации по зарплате: %s", err)
			}
		case "h", "help":
			fmt.Println("a, add - добавить сотрудника\nd, disp - вывод информации о сотруднике по табельному номеру\nf, find - фильтрация по возрасту и зарплате\nh, help - список команд\nq, quit - выход")
		case "q", "quit":
			fmt.Print("Закрытие программы...")
			os.Exit(0)
		default:
			fmt.Println("Неизвестная команда.")
		}
	}
}
