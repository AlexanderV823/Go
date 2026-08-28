# URL Shortener (Сервис сокращения URL-адресов)

Быстрый и потокобезопасный HTTP-сервис на Go для сокращения URL с защитой от Race Conditions и полным покрытием тестами.

## 🛠️ Стек технологий
* **Go** (≥ 1.18), встроенный `net/http`
* In-memory хранилище на `sync.RWMutex` и `map` с дедупликацией
* Генерация ID через `crypto/rand`, тестирование стандартными средствами

## 📁 Структура проекта
* `main.go` — Запуск сервера и JSON-ошибки.
* `shortener.go` — Бизнес-логика и защита от дубликатов.
* `shortener_test.go` — Юнит и стресс-тесты.
* `handlers_test.go` — Интеграционные HTTP-тесты.

## 🚀 Запуск и тесты
Запуск сервера: `go run main.go shortener.go`
Тесты с покрытием и детектором гонок: `go test -v -race`

## 📡 Спецификация API

### 1. Сокращение URL (POST /shorten)

Принимает длинный URL в формате JSON и возвращает уникальный короткий ID.

```bash
curl -X POST http://localhost:8080/shorten -H "Content-Type: application/json" -d '{"url": "https://google.com"}'
```

### 2. Редирект на оригинальный URL (GET /{short_url})

Принимает короткий идентификатор и возвращает статус 302 Found.

```bash
curl -I http://localhost:8080/{short_url}
```
*(Где `{short_url}` — это ID, полученный из предыдущего запроса, например: 4a_RzA)*
