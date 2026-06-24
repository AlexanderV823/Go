package main

import (
	"database/sql"
	"log"
	"fmt"

	_ "github.com/mattn/go-sqlite3"

	"hw_1/internal/infrastructure/db"
	"hw_1/internal/infrastructure/notification"
	"hw_1/internal/service"
)

func main() {
	database, err := sql.Open("sqlite3", "orders.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// 1. Инициализация и сборка инфраструктурных адаптеров
	sqliteRepo := db.NewSQLiteRepository(database)
	if err := sqliteRepo.InitSchema(); err != nil {
		log.Fatal(err)
	}

	emailNotifier := notification.NewEmailService()
	smsNotifier := notification.NewSMSSender()


	fmt.Println("--- Заказ с Email-уведомлением ---")
	// 2. Внедрение зависимостей в слой бизнес-логики
	orderServiceWithEmail := service.NewOrderService(sqliteRepo, emailNotifier)

	err = orderServiceWithEmail.CreateOrder("Иван", []string{"onion", "potato"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\n--- Заказ с SMS-уведомлением ---")

	orderServiceWithSMS := service.NewOrderService(sqliteRepo, smsNotifier)
	// 3. Запуск бизнес-процесса
	err = orderServiceWithSMS.CreateOrder("Василий", []string{"apple"}, 5.2)
	if err != nil {
		log.Fatal(err)
	}
}
