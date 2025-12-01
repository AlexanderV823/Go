package main

import (
	"fmt"
	"hw_7/mymath"
)

func main() {

	var i string

	for {
		fmt.Println("Выполнить расчет? (д/н): ")
		_, err := fmt.Scan(&i)

		if err != nil {
			fmt.Println("Ошибка: %w", err)
			return
		}

		if i == "д" {
			calc()
		} else if i == "н" {
			fmt.Println("Выход из программы...")
			break
		} else {
			fmt.Println("Неизвестная команда. Повторите ввод")
		}
	}
}

func calc() {

	var (
		a int
		b int
	)

	fmt.Println("Введите число a: ")
	_, err := fmt.Scan(&a)

	if err != nil {
		fmt.Println("Ошибка: %w", err)
		return
	}
	fmt.Println("Введите число b: ")
	_, err = fmt.Scan(&b)

	if err != nil {
		fmt.Println("Ошибка: %w", err)
		return
	}
	c := mymath.Add(a, b)

	fmt.Printf("Результат сложения: %d \n", c)

	c = mymath.Multiply(a, b)

	fmt.Printf("Результат умножения: %d \n", c)

}
