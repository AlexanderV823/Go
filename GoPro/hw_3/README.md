# Практическое задание «Создание REST API‑сервиса (CRUD задач)»

**1. Tasks API**

REST API сервис на Go для управления списком задач (to-do). Сервис работает полностью в оперативной памяти (in-memory) с поддержкой конкурентного доступа (sync.RWMutex) и обменивается данными строго в формате JSON.

**2. Архитектура проекта**
```
tasks-api/
  ├── cmd/
  │   └── server/
  │       └── main.go          # Точка входа, запуск HTTP-сервера
  ├── internal/
  │   ├── handlers/
  │   │   └── tasks.go         # HTTP-обработчики (бизнес-логика)
  │   ├── models/
  │   │   └── task.go          # Описание структуры Task
  │   └── storage/
  │       ├── storage.go       # Интерфейс Storage
  │       └── memory/
  │           └── memory.go    # Потокобезопасная in-memory реализация
  └── go.mod                   # Модуль Go
```

**3. Как запустить проект**

* Откройте терминал в корневой папке проекта (там, где находится go.mod).
* Для корректного отображения кириллицы в консоли Windows выполните:
   * В CMD: chcp 65001
   * В PowerShell: [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
* Запустите сервер: go run cmd/server/main.go
* Сервер начнет логировать входящие запросы и слушать порт :8080.

**4. Спецификация API (Таблица эндпоинтов)**


| Метод      | Эндпоинт      | Описание                           | Тело запроса (JSON)                 | Код (Успех) | Код (Ошибка) |
| :--------- | :------------ | :--------------------------------- | :---------------------------------- | :---------: | :----------: |
| **GET**    | `/health`     | Проверка работоспособности сервиса | —                                   |    `200`    |      —       |
| **GET**    | `/tasks`      | Получить список всех задач         | —                                   |    `200`    |    `405`     |
| **POST**   | `/tasks`      | Создать новую задачу               | `{"title": "string"}`               |    `201`    | `400`, `405` |
| **GET**    | `/tasks/{id}` | Получить задачу по идентификатору  | —                                   |    `200`    | `400`, `404` |
| **PUT**    | `/tasks/{id}` | Обновить задачу целиком            | `{"title": "string", "done": bool}` |    `200`    | `400`, `404` |
| **DELETE** | `/tasks/{id}` | Удалить задачу по идентификатору   | —                                   |    `204`    | `400`, `404` |

Все ответы (включая ошибки) возвращаются с заголовком Content-Type: application/json. Единый формат ошибок: {"error": "сообщение"}.

OpenAPI 3.0 Спецификация: [openapi.yaml](./openapi.yaml)

**5. Тестирование**

Тестовые примеры можно запустить в Postman: [collection.json](./collection.json)

* **Проверка работоспособности (GET /health)**

**Успешный сценарий:**
curl -i http://localhost:8080/health

**Ответ:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:10:00 GMT
Content-Length: 15

{"status":"OK"}
```

* **Создание задачи (POST /tasks)**

**Успешный сценарий:**
curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\": \"Зайти в Озон\"}"

**Ответ:**
```
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:10:15 GMT
Content-Length: 82

{"id":1,"title":"Зайти в Озон","done":false,"created_at":"2026-08-05T21:10:15+03:00"}
```

**Ошибочный сценарий (пустой title):**
curl -i -X POST http://localhost:8080/tasks -H "Content-Type: application/json" -d "{\"title\": \"   \"}"

**Ответ:**
```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:10:25 GMT
Content-Length: 39

{"error":"Поле 'title' обязательно"}
```

* **Получение списка задач (GET /tasks)**

**Успешный сценарий:**
curl -i http://localhost:8080/tasks

**Ответ:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:10:40 GMT
Content-Length: 84

[{"id":1,"title":"Зайти в Озон","done":false,"created_at":"2026-08-05T21:10:15+03:00"}]
```

**Ошибочный сценарий (недопустимый метод PATCH для коллекции):**
curl -i -X PATCH http://localhost:8080/tasks

**Ответ:**
```
HTTP/1.1 405 Method Not Allowed
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:10:48 GMT
Content-Length: 43

{"error":"Метод не поддерживается"}
```

* **Получение задачи по ID (GET /tasks/{id})**

**Успешный сценарий:**
curl -i http://localhost:8080/tasks/1

**Ответ:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:11:02 GMT
Content-Length: 82

{"id":1,"title":"Зайти в Озон","done":false,"created_at":"2026-08-05T21:10:15+03:00"}
```

**Ошибочный сценарий (ID не существует):**
curl -i http://localhost:8080/tasks/999

**Ответ:**
```
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:11:10 GMT
Content-Length: 35

{"error":"Задача не найдена"}
```

* **Обновление задачи (PUT /tasks/{id})**

**Успешный сценарий:**
curl -i -X PUT http://localhost:8080/tasks/1 -H "Content-Type: application/json" -d "{\"title\": \"Забрать посылку из Озон\", \"done\": true}"

**Ответ:**
```
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:11:35 GMT
Content-Length: 99

{"id":1,"title":"Забрать посылку из Озон","done":true,"created_at":"2026-08-05T21:10:15+03:00"}
```

**Ошибочный сценарий (невалидный ID в формате букв):**
curl -i -X PUT http://localhost:8080/tasks/abc -H "Content-Type: application/json" -d "{\"title\": \"Тест\", \"done\": true}"

**Ответ:**
```
HTTP/1.1 400 Bad Request
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:11:45 GMT
Content-Length: 53

{"error":"Неверный или отсутствующий ID задачи"}
```

* **Удаление задачи (DELETE /tasks/{id})**

**Успешный сценарий:**
curl -i -X DELETE http://localhost:8080/tasks/1

**Ответ:**
```
HTTP/1.1 204 No Content
Date: Wed, 05 Aug 2026 18:12:05 GMT
```

**Ошибочный сценарий (удаление уже несуществующей задачи):**
curl -i -X DELETE http://localhost:8080/tasks/1

**Ответ:**
```
HTTP/1.1 404 Not Found
Content-Type: application/json
Date: Wed, 05 Aug 2026 18:12:12 GMT
Content-Length: 35

{"error":"Задача не найдена"}
```

* **Логирование на стороне сервера**

При выполнении всех тестовых сценариев в консоли сервера отображаются логи в формате [МЕТОД] Путь:
```
2026/08/05 21:10:00 [GET] /health
2026/08/05 21:10:15 [POST] /tasks
2026/08/05 21:10:25 [POST] /tasks
2026/08/05 21:10:40 [GET] /tasks
2026/08/05 21:10:48 [PATCH] /tasks
2026/08/05 21:11:02 [GET] /tasks/1
2026/08/05 21:11:10 [GET] /tasks/999
2026/08/05 21:11:35 [PUT] /tasks/1
2026/08/05 21:11:45 [PUT] /tasks/abc
2026/08/05 21:12:05 [DELETE] /tasks/1
2026/08/05 21:12:12 [DELETE] /tasks/1
```