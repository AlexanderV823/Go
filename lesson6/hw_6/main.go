package main

import (
	"fmt"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/addEmployee"
)

type Displayable interface {
	Display()
}

// Структура сотрудника
type Employee struct {
	ID      int    //Табельный номер
	Surname string //Фамилия
	Name    string //Имя
	Age     int    //Возраст
	Job     string //Должность
	Salary  int    //Зарплата
}

func (emp Employee) Display() { //В параметре ждать данные для вывода

	fmt.Println() //Выводить присланнные данные

}

// Срез сотрудников
var Employees = make([]Employee, 3)

func main() {
	fmt.Println("Введите команду:\nadd - добавить сотрудника\ndisp - вывод информации о сотруднике\nf - фильтрация по возрасту и зарплате\nq, quit - выход") //подсказки вывести в отдельную функцию
	for {
		var input string
		fmt.Scanln(&input)
		if input == "add" {
			var (
				id      int
				surname string
				name    string
				age     int
				job     string
				salary  int
			)
			fmt.Println("Введите таб.номер, фамилию, имя, возраст, должность и зарплату: ")
			_, err := fmt.Scanf("%d %s %s %d %s %d", &id, &surname, &name, &age, &job, &salary)
			if err != nil {
				fmt.Println("\nОшибка ввода значения:", err)
			}
			newEmp := Employee{ID: id, Surname: surname, Name: name, Age: age, Job: job, Salary: salary}

			err := addEmployee(newEmp, *Employees)
			if err != nil {
				fmt.Println("\nОшибка добавления сотрудника:", err)
			} else {
				//Сообщаем об успешной записи
				fmt.Println("Данные о сотруднике записаны")
			}
			// } else if input == "p" {
			// 	if Displayable != nil {
			// 		Displayable.Display() //Передавать необходимые данные для вывода
			// 	}
			// 	fmt.Print("\nВведите команду: ")
			// } else if input == "d" {
			// 	calculateDiscount()
			// 	fmt.Print("\nВведите команду: ")
		} else if input == "quit" || input == "q" {
			fmt.Println("\nЗакрытие программы...")
			break
		}
	}
}
