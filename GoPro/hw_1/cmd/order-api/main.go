package main

import (
	"database/sql"
	"log"

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

	notifier := notification.NewEmailService()

	// 2. Внедрение зависимостей в слой бизнес-логики
	orderService := service.NewOrderService(sqliteRepo, notifier)

	// 3. Запуск бизнес-процесса
	err = orderService.CreateOrder("Иван", []string{"onion", "potato"}, 10.5)
	if err != nil {
		log.Fatal(err)
	}
}
