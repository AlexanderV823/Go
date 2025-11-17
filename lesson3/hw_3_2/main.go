// Программа для управления небольшим складом товаров
package main

import (
	"fmt"
	"time"
)

// Глобальные константы для категорий товаров
const (
	CategoryElectronics = "Электроника"
	CategoryFood        = "Продукты"
	CategoryClothes     = "Одежда"
	MaxItems            = 100 // Максимальное количество товаров на складе
)

// Объявление Среза для хранения карт(Структур) товара
var ItemSlice []Item

// Объявление структуры товара
type Item struct {
	id        uint32
	name      string
	qty       uint32
	price     float64
	category  string
	dateAdded time.Time
}

// Глобальная переменная для подсчета товаров
var totalItems uint32

func main() {
	fmt.Println("Введите команду:\nadd - добавить товар\nd - расчет скидки\np -вывод количества товаров\nq, quit - выход")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "add" {
			val, err := addNewItem()
			if err != nil {
				fmt.Println("\nОшибка ввода значения:", err)
				fmt.Println("\nВведите команду:")
			} else {
				fmt.Printf("\nТовар добавлен. ID: %d\n", val)
				fmt.Print("\nВведите команду: ")
			}
		} else if input == "p" {
			printQtyItems()
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

// Функция для добавления нового товара и возвращение его ID
func addNewItem() (uint32, error) {

	var (
		id        uint32
		name      string
		qty       uint32
		price     float64
		category  string
		dateAdded time.Time = time.Now()
		input     string
		newItem   Item
	)

	fmt.Println("\nВведите данные о товаре:")
	fmt.Print("ID: ")
	if _, err := fmt.Scanln(&id); err != nil {
		return 0, fmt.Errorf("ожидается число в диапазоне 0 — 4 294 967 295")
	}
	newItem.id = id

	fmt.Print("Название: ")
	if _, err := fmt.Scan(&name); err != nil {
		fmt.Println("Ошибка ввода. Не правильный тип значения.", err)
		return 0, fmt.Errorf("ожидается ввод строкового значения")
	}
	newItem.name = name

	fmt.Print("Введите код категории (1 - Электроника, 2 - Продукты, 3 - Одежда): ")
	fmt.Scan(&input)
	if input == "1" {
		category = "Электроника"
	} else if input == "2" {
		category = "Продукты"
	} else if input == "3" {
		category = "Одежда"
	} else {
		return 0, fmt.Errorf("указана не существующая категория")
	}
	newItem.category = category
	fmt.Print("Количество: ")
	if _, err := fmt.Scan(&qty); err != nil {
		fmt.Println("Ошибка ввода. Не правильный тип значения.", err)
		return 0, fmt.Errorf("ожидается число в диапазоне 0 — 4 294 967 295")
	} else {
		newItem.qty = qty
	}
	fmt.Print("Цена: ")
	if _, err := fmt.Scan(&price); err != nil {
		return 0, fmt.Errorf("ожидается число с плавающей точкой")
	} else {
		newItem.price = price
	}
	newItem.dateAdded = dateAdded

	ItemSlice = append(ItemSlice, newItem)

	totalItems += 1

	return newItem.id, nil
}

// Вывод количества внесенных товаров
func printQtyItems() {
	fmt.Printf("Внесено товаров: %d\n", totalItems)
}

// Расчет скидки
func calculateDiscount() {
	var (
		itemPrice    float64
		itemDiscount float64
	)

	fmt.Println("Введите цену товара:")
	fmt.Scan(&itemPrice)
	fmt.Println("Введите размер скидки:")
	fmt.Scan(&itemDiscount)
	discount := (itemPrice * itemDiscount) / 100
	fmt.Printf("\nЦена с учетом скидки: %.2f\n", itemPrice-discount)
}
