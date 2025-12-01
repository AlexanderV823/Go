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
		_, err := fmt.Scanln(&input)
		if err != nil {
			if err.Error() == "unexpected newline" {
				continue //если введен Enter, то выводить сообщение об ощибке не обязательно
			}
			fmt.Printf("Ошибка ввода: %s\n", err)
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
			_, err := fmt.Scanln(&id, &surname, &name, &age, &job, &salary)

			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s\n", err)
				continue
			}

			newEmp := employee.Employee{ID: id, Surname: surname, Name: name, Age: age, Job: job, Salary: salary}
			err = employee.Add(newEmp)

			if err != nil {
				fmt.Printf("Ошибка добавления сотрудника: %s\n", err)
				continue
			} else {
				//Сообщаем об успешной записи
				fmt.Println("Данные о сотруднике записаны")
			}
		case "d", "disp":
			var id int
			fmt.Print("Введите табельный номер: ")
			_, err := fmt.Scanln(&id)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s\n", err)
				continue
			}
			id, err = employee.Find(employee.Employees, id)
			if err != nil {
				fmt.Printf("Ошибка поиска сотрудника: %s\n", err)
				continue
			}
			if id == -1 {
				fmt.Println("Сотрудник с таким табельным номером не найден")
				continue
			}
			emp := employee.Employees[id]
			emp.Display()
		case "f", "find":
			var (
				minAge    int
				minSalary float64
			)
			fmt.Print("Введите минимальный возраст: ")
			_, err := fmt.Scanln(&minAge)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s\n", err)
				continue
			}
			fmt.Print("Введите минимальную зарплату: ")
			_, err = fmt.Scanln(&minSalary)
			if err != nil {
				fmt.Printf("Ошибка ввода значения: %s", err)
				continue
			}
			emp, err := employee.FilterAgeSalary(minAge, minSalary)
			if err != nil {
				fmt.Printf("Ошибка фильтрации по возрасту и зарплате: %s\n", err)
				continue
			}
			for _, e := range emp {
				e.Display()
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
