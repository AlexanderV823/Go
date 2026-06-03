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

                      [ API Gateway / Application Layer ]
                         /                           \
         (Route by Functional Domain)          (Route by Functional Domain)
                       /                                       \
         [ User Service / Router ]                      [ Catalog Service ]
          /                     \                               |
  hash(id)%2=0              hash(id)%2=1                        |
        /                         \                             |
[ USER SHARD 1 ]           [ USER SHARD 2 ]             [ CATALOG SHARD ]
(Таблица: users A-M)       (Таблица: users N-Z)         (Таблицы: books, shops)
  Режим: Read/Write          Режим: Read/Write            Режим: Read/Write

        |                          |                            |
        v                          v                            v
[ USER REPLICA 1 ]         [ USER REPLICA 2 ]           [ CATALOG REPLICA ]
  Режим: Read-Only           Режим: Read-Only             Режим: Read-Only

Для обеспечения высокой доступности и отказоустойчивости каждый шард разворачивается в режиме Primary-Standby (Master-Slave) репликации:
1. Основной слой (Primary / Master):
* Сервера: User Shard 1 (Master), User Shard 2 (Master), Catalog Shard (Master).
* Режим работы: Read/Write (Чтение и запись).
* Назначение: Принимать любые транзакции (INSERT, UPDATE, DELETE) от микросервисов.
2. Репликационный слой (Standby / Slave):
* Сервера: User Replica 1, User Replica 2, Catalog Replica.
* Режим работы: Read-Only (Только чтение, асинхронное получение WAL-логов).
* Назначение: Выполнение тяжелых поисковых запросов (SELECT), чтобы не нагружать ими Master-сервер. При падении Master-сервера, реплика автоматически (через оркестратор вроде Patroni) переключается в режим Master.