package main //Объявляем главный пакет

import (
	"fmt" //стандартный пакет для вывода текста
)

func main() {
	fmt.Println("Hello, World!") //Непосредственно вывод сообщения
	fmt.Println("Hello, Sanek!")
	a := 1
	b := 0
	if b == 0 {
		fmt.Println("Division by zero!")
		return
	}
	fmt.Println(a / b)

}
