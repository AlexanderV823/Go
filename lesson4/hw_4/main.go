package main

import "fmt"

func main() {
	// Инициализация переменной для ввода
	var age int

	// Цикл для 5 пользователей
	for i := 0; i < 5; i++ {
		// Запрос возраста у пользователя
		fmt.Print("Введите возраст: ")
		fmt.Scan(&age)

		// Определение статуса в зависимости от возраста
		switch {
		case age < 18:
			fmt.Println("Ребёнок")
		case age < 65:
			fmt.Println("Взрослый")
		default:
			fmt.Println("Пенсионер")
		}
	}
}
