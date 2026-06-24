# Домашнее задание к уроку занятию Рефакторинг программы по принципам SOLID

## Исходный код

```
 package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

type Order struct {
    ID       int
    Customer string
    Products string
    Total    float64
    Status   string
}

type OrderSystem struct {
    db *sql.DB
}

func NewOrderSystem(db *sql.DB) *OrderSystem {
    return &OrderSystem{db: db}
}

func (s *OrderSystem) CreateOrder(customer string, products []string, total float64) error {
    // Создание заказа в БД
    _, err := s.db.Exec(
        "INSERT INTO orders (customer, products, total, status) VALUES (?, ?, ?, ?)",
        customer, fmt.Sprintf("%v", products), total, "pending",
    )
    if err != nil {
        return err
    }

    // Отправка уведомления
    s.sendEmailNotification(customer)

    return nil
}

func (s *OrderSystem) sendEmailNotification(customer string) {
    fmt.Printf("Уведомление отправлено клиенту %s\n", customer)
}

func main() {
    db, err := sql.Open("sqlite3", "orders.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    _, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        customer TEXT NOT NULL,
        products TEXT NOT NULL,
        total REAL NOT NULL,
        status TEXT NOT NULL
    )`)
    if err != nil {
        log.Fatal(err)
    }

    system := NewOrderSystem(db)

    err = system.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
    if err != nil {
        log.Fatal(err)
    }
}
```

Примечание: приведённый код — рабочий. Для запуска вам нужно выполнить команды:
* инициализировать модуль: go mod init example/solid
* установить зависимости: go mod tidy

После этого можно запустить: go run main.go

## Требования к рефакторингу

Примените Single Responsibility Principle:
* Разделите ответственность за работу с БД и бизнес-логику
* Выделите отправку уведомлений в отдельный компонент

Реализуйте Open/Closed Principle:
* Сделайте систему расширяемой для добавления новых типов баз данных и уведомлений

Примените Liskov Substitution Principle:
* Создайте базовый интерфейс для работы с сообщениями
* Реализуйте разные типы отправителей сообщений

Реализуйте Interface Segregation Principle:
* Разделите большие интерфейсы на мелкие, если требуется
* Создайте отдельные интерфейсы для инициализации базы данных (создание таблиц) и записи

Примените Dependency Inversion Principle:
* Внедрите зависимости через интерфейсы
* Используйте внедренные зависимости