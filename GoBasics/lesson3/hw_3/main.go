package main

import "fmt"

const AppName string = "UserPrifile"

var name string = "Александр"
var age uint8 = 38
var height uint8 = 192
var newsletter bool

func printProfile() {
	fmt.Println("Добро пожаловать в GoFit!")
	fmt.Println("Профиль пользователя:")
	fmt.Printf("Имя: %s\n", name)
	fmt.Printf("Возраст: %d\n", age)
	fmt.Printf("Рост: %d\n", height)
	fmt.Printf("Подписан на рассылку: %t\n", newsletter)
}

func main() {
	printProfile()
}
