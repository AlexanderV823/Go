package main

import (
	"fmt"
	"hw_7/mymath"
)

func main() {

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
