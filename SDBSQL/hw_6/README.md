# Домашнее задание к занятию "Репликация и масштабирование"

## Задание 1

Выполните конфигурацию master-slave репликации, примером можно пользоваться из лекции.

1. Подготовка тестовой среды в Docker Compose

mkdir pg-replication && cd pg-replication

* Переместить (mv docker-compose.yml pg-replication/) или отредактировать (nano docker-compose.yml) файл: [docker-compose.yml](./docker-compose.yml)

* Переместить (mv pg_hba.conf pg-replication/) или отредактировать (nano pg_hba.conf) файл: [pg_hba.conf](./pg_hba.conf)

* docker compose up -d

2. Проверка

* Проверка логов процесса клонирования и запуска

docker compose logs -f pg-slave

![Скрин 1](./hw6-1_1.jpg)

* Подключение к Master и проверка статуса отправки данных

docker exec -it pg-master psql -U postgres -c "SELECT * FROM pg_stat_replication;"

![Скрин 2](./hw6-1_2.jpg)

* Создание таблицы на Master и добавление строки

docker exec -it pg-master psql -U postgres -d mydb -c 'INSERT INTO test_sync (val) VALUES ('\''Hello, Sanek!'\'');'

![Скрин 3](./hw6-1_3.jpg)

* Тест чтения на Slave

docker exec -it pg-slave psql -U postgres -d mydb -c "SELECT * FROM test_sync;"

![Скрин 4](./hw6-1_4.jpg)

* Попытка записи на Slave

docker exec -it pg-slave psql -U postgres -d mydb -c "INSERT INTO test_sync (val) VALUES ('Error');"

![Скрин 5](./hw6-1_5.jpg)

## Задание 2

graph TD
    %% Стилизация элементов
    classDef app fill:#e1f5fe,stroke:#0288d1,stroke-width:2px;
    classDef router fill:#fff9c4,stroke:#fbc02d,stroke-width:2px;
    classDef master fill:#e8f5e9,stroke:#388e3c,stroke-width:2px;
    classDef slave fill:#ffebee,stroke:#d32f2f,stroke-width:2px;

    %% Слой приложения
    App[Бэкенд Приложения / API Gateway]:::app
    Router{Маршрутизатор Запросов<br/>Data Access Layer}:::router

    App --> Router

    %% Вертикальное шардирование
    Router -->|Запросы Книг / Магазинов| CatM[Catalog Shard: MASTER<br/>Порт: 5435<br/>Режим: Read/Write]:::master
    CatM -->|Асинхронная репликация| CatS[Catalog Shard: SLAVE<br/>Порт: 5436<br/>Режим: Read-Only]:::slave

    %% Горизонтальное шардирование пользователей
    Router -->|Пользователи: hash ID % 2 != 0| User1M[User Shard 1: MASTER<br/>Порт: 5437<br/>Режим: Read/Write]:::master
    User1M -->|Асинхронная репликация| User1S[User Shard 1: SLAVE<br/>Порт: 5438<br/>Режим: Read-Only]:::slave

    Router -->|Пользователи: hash ID % 2 == 0| User2M[User Shard 2: MASTER<br/>Порт: 5439<br/>Режим: Read/Write]:::master
    User2M -->|Асинхронная репликация| User2S[User Shard 2: SLAVE<br/>Порт: 5440<br/>Режим: Read-Only]:::slave

    %% Связи для чтения
    Router -.->|Тяжелые SELECT запросы| CatS
    Router -.->|Тяжелые SELECT запросы| User1S
    Router -.->|Тяжелые SELECT запросы| User2S

Для обеспечения высокой доступности и отказоустойчивости каждый шард разворачивается в режиме Primary-Standby (Master-Slave) репликации:
1. Основной слой (Primary / Master):
* Сервера: User Shard 1 (Master), User Shard 2 (Master), Catalog Shard (Master).
* Режим работы: Read/Write (Чтение и запись).
* Назначение: Принимать любые транзакции (INSERT, UPDATE, DELETE) от микросервисов.
2. Репликационный слой (Standby / Slave):
* Сервера: User Replica 1, User Replica 2, Catalog Replica.
* Режим работы: Read-Only (Только чтение, асинхронное получение WAL-логов).
* Назначение: Выполнение тяжелых поисковых запросов (SELECT), чтобы не нагружать ими Master-сервер. При падении Master-сервера, реплика автоматически (через оркестратор вроде Patroni) переключается в режим Master.

## Задача 3

1. Активный master и пассивный репликационный slave (Hot Standby)

Этот подход ориентирован на отказоустойчивость (High Availability), а не на распределение нагрузки.

* Минимальный Replication Lag: Пассивный сервер не нагружен SELECT-запросами пользователей. Он тратит 100% ресурсов на применение WAL-логов, гарантируя минимальное отставание от мастера.
* Безопасное резервное копирование: Снятие бэкапов (pg_dump) или аналитические тяжелые выгрузки выполняются на Slave. Это исключает блокировки таблиц на основном рабочем сервере (Master).
* Быстрый и чистый Failover: В случае аварии пассивный Slave готов мгновенно стать новым Мастером, так как его дисковая подсистема и память не перегружены сторонними операциями.

2. Master-сервер и несколько slave-серверов (Read-Scalability)

Архитектура создана для горизонтального масштабирования производительности приложений с преобладанием операций чтения (Read-Heavy).

* Линейное масштабирование чтения: Добавление новых Slave-серверов позволяет пропорционально увеличивать количество обрабатываемых SELECT-запросов в секунду (RPS).
* Разгрузка Master-сервера: Мастер полностью освобождается от задач чтения и направляет свои CPU и диск исключительно на транзакции записи (INSERT, UPDATE, DELETE).
* Повышенная живучесть: Выход из строя одного или нескольких Slave-серверов не приводит к отказу системы. Трафик чтения автоматически балансируется между оставшимися репликами.

3. Активный сервер с механизмом репликации DRBD (RAID по сети)

DRBD (Distributed Replicated Block Device) работает на уровне ядра ОС и реплицирует данные не средствами СУБД, а поблочно через сеть.

* Абсолютная независимость от СУБД: Репликация происходит на уровне файловой системы. Архитектура защищает не только данные базы, но и любые конфигурационные файлы или системные логи.
* Исключение логических ошибок: Поскольку копируются блоки данных (нули и единицы), исключены ошибки парсинга SQL-команд или рассинхронизации логических транзакций.
* Синхронность «из коробки»: DRBD гарантирует, что блок данных физически запишется на диски обоих серверов до того, как операционная система подтвердит успешное завершение операции записи СУБД.

4. SAN-кластер (Storage Area Network)

Это архитектура с общей дисковой памятью (Shared Storage), где несколько вычислительных серверов СУБД подключены к одной дисковой полке по высокоскоростным каналам (Fibre Channel/iSCSI).

* Нулевое отставание (No Replication Lag): Сервера не тратят ресурсы на пересылку логов или данных друг другу — они смотрят на одни и те же физические файлы данных в один и тот же момент времени.
* Мгновенное восстановление (RTO = 0): Если падает вычислительный узел (сервер), запасному серверу не нужно накатывать логи. Он просто перехватывает управление дисковым массивом и продолжает работу с той же секунды.
* Аппаратная надежность промышленного уровня: Дисковые массивы SAN оснащены дублирующими контроллерами, независимыми блоками питания и аппаратными RAID, что исключает потерю данных на уровне железа.


5. Сравнительная таблица методов

| Критерий | Master + 1 Пассивный Slave | Master + Много Slaves | Активный сервер + DRBD | SAN-кластер |
| Основной упор  | Надежность (HA) | Скорость чтения | Целостность диска | Отсутствие задержек |
| Уровень работы  | Логический (СУБД) | Логический (СУБД) | Блочный (Ядро ОС) | Аппаратный (Железо) |
| Основной упор  | Минимальная | Зависит от нагрузки | Нулевая (при синхронном) | Полностью отсутствует |

## Задача 4

* Переносим или редактируем Docker Compose файл: [docker-compose.yml](./docker-compose_2.yml)

nano docker-compose.yml

* Разворачиваем и проверяем работу

docker compose up -d

docker compose ps

* Используемые порты

| Контейнер | Назначение | Порт | Разрешенные операции |
| pg-catalog-master | Книги и магазины | 5435 | Запись + Чтение |
| pg-catalog-slave | Книги и магазины (Реплика) | 5436 | Только SELECT |
| pg-user1-master | Пользователи (Нечетные ID) | 5437 | Запись + Чтение |
| pg-user1-slave | Пользователи (Нечетные ID, Реплика) | 5438 | Только SELECT |

* Проверка кластера Каталога (Книги/Магазины)

docker exec -it pg-catalog-master psql -U postgres -d catalog_db -c "SELECT application_name, state FROM pg_stat_replication;"

* Проверка кластера Пользователей (Шард 1 — Нечетные)

docker exec -it pg-user1-master psql -U postgres -d user_db -c "SELECT application_name, state FROM pg_stat_replication;"

* Проверка кластера Пользователей (Шард 2 — Четные)

docker exec -it pg-user2-master psql -U postgres -d user_db -c "SELECT application_name, state FROM pg_stat_replication;"

![Скрин 6](./hw6-4_1.jpg)

* Тестирование вертикального шардинга (Каталог)

* Запись на Master

docker exec -it pg-catalog-master psql -U postgres -d catalog_db -c "
CREATE TABLE IF NOT EXISTS shops (shop_id INT PRIMARY KEY, title TEXT);
CREATE TABLE IF NOT EXISTS books (book_id INT PRIMARY KEY, title TEXT, shop_id INT);
INSERT INTO shops VALUES (1, 'Буквоед');
INSERT INTO books VALUES (10, 'Капитанская дочка', 1);
"

* Чтение со Slave

docker exec -it pg-catalog-slave psql -U postgres -d catalog_db -c "SELECT * FROM books;"

![Скрин 7](./hw6-4_2.jpg)

* Тестирование горизонтального шардинга (Пользователи)

* Запись структуры на Шард 1

docker exec -it pg-user1-master psql -U postgres -d user_db -c "
CREATE TABLE IF NOT EXISTS users (user_id INT PRIMARY KEY, name TEXT, CONSTRAINT chk_odd CHECK (user_id % 2 != 0));
"

* Запись структуры на Шард 2

docker exec -it pg-user2-master psql -U postgres -d user_db -c "
CREATE TABLE IF NOT EXISTS users (user_id INT PRIMARY KEY, name TEXT, CONSTRAINT chk_even CHECK (user_id % 2 = 0));
"

* Проверка правильности записи на Шард 1 пользователя Иван (ID=1)

docker exec -it pg-user1-master psql -U postgres -d user_db -c "INSERT INTO users VALUES (1, 'Иван');"

* Проверка правильности записи на Шард 2 пользователя Анна (ID=2)

docker exec -it pg-user2-master psql -U postgres -d user_db -c "INSERT INTO users VALUES (2, 'Анна');"

* Проверка Реплики 1 (д.б. Иван)

docker exec -it pg-user1-slave psql -U postgres -d user_db -c "SELECT * FROM users;"

* Проверка Реплики 2 (д.б. выведена Анна)

docker exec -it pg-user2-slave psql -U postgres -d user_db -c "SELECT * FROM users;"

![Скрин 7](./hw6-4_3.jpg)

* Тест проверки ограничений шардинга

Отправка пользователя с четным ID на нечетный шард

docker exec -it pg-user1-master psql -U postgres -d user_db -c "INSERT INTO users VALUES (4, 'Петр');"

![Скрин 8](./hw6-4_4.jpg)