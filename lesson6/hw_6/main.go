package main

import (
	"fmt"
	"github.com/AlexanderV823/Go/tree/main/lesson6/hw_6/addEmployee"
)

type Displayable interface {
	Display()
}

// Структура должности
type Job struct {
	ID         int    //Код
	Name       string //Название
	Department string //Подразделение
}

// Структура сотрудника
type Employee struct {
	ID      int    `json:"id"`      //Табельный номер
	Surname string `json:"surname"` //Фамилия
	Name    string `json:"name"`    //Имя
	Age     int    `json:"age"`     //Возраст
	Job     Job    `json:"job"`     //Должность
	Salary  int    `json:"salary"`  //Зарплата
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
			// val, err := addNewItem()
			// if err != nil {
			// 	fmt.Println("\nОшибка ввода значения:", err)
			// 	fmt.Println("\nВведите команду:")
			// } else {
			// 	fmt.Printf("\nТовар добавлен. ID: %d\n", val)
			// 	fmt.Print("\nВведите команду: ")
			// }
		} else if input == "p" {
			if Displayable != nil {
				Displayable.Display() //Передавать необходимые данные для вывода
			}
			fmt.Print("\nВведите команду: ")
		} else if input == "d" {
			calculateDiscount()
			fmt.Print("\nВведите команду: ")
		} else if input == "quit" || input == "q" {
			fmt.Println("\nЗакрытие программы...")
			break
		}
	}
}
