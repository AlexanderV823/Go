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
	itemID    uint32
	itemName  string
	itemQty   uint32
	itemPrice float64
	dateAdded time.Time
}

// Глобальная переменная для подсчета товаров
var totalItems int

func main() {
	fmt.Println("Введите команду:\nadd - добавить товар\nd - расчет скидки\np -вывод количества товаров\nq, quit - выход")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "add" {
			fmt.Printf("Товар добавлен. ID: %d\n", addNewItem())
			fmt.Println("Введите команду:")
		} else if input == "p" {
			printQtyItems()
			fmt.Println("Введите команду:")
		} else if input == "d" {
			calculateDiscount()
			fmt.Println("Введите команду:")
		} else if input == "quit" || input == "q" {
			break
		}
	}
}

// Функция для добавления нового товара и возвращение его ID
func addNewItem() uint32 {

	var (
		itemID    uint32
		itemName  string
		itemQty   uint32
		itemPrice float64
		dateAdded time.Time = time.Now()
	)

	fmt.Println("\nВведите данные о товаре:")
	fmt.Println("ID: ")
	fmt.Scan(&itemID)
	fmt.Println("Название: ")
	fmt.Scan(&itemName)
	//fmt.Println("Категория: "); fmt.Scan(&category)
	fmt.Println("Количество: ")
	fmt.Scan(&itemQty)
	fmt.Println("Цена: ")
	fmt.Scan(&itemPrice)

	var newItem Item
	newItem.itemID = itemID
	newItem.itemName = itemName
	newItem.itemQty = itemQty
	newItem.itemPrice = itemPrice
	newItem.dateAdded = dateAdded

	ItemSlice = append(ItemSlice, newItem)

	return newItem.itemID
}

// Вывод количества внесенных товаров
func printQtyItems() {
	fmt.Println("Внесено товаров: ", len(ItemSlice))
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
