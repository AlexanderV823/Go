# Домашнее задание к занятию "Индексы"

## Задание 1

Напишите запрос к учебной базе данных, который вернёт процентное отношение общего размера всех индексов к общему размеру всех таблиц.

SELECT
    ROUND(SUM(index_length) / SUM(data_length + index_length) * 100, 2) AS index_to_total_percentage
FROM information_schema.tables
WHERE table_schema = 'sakila' AND table_type = 'BASE TABLE';

## Задание 2

Выполните explain analyze следующего запроса:

select distinct concat(c.last_name, ' ', c.first_name), sum(p.amount) over (partition by c.customer_id, f.title)
from payment p, rental r, customer c, inventory i, film f
where date(p.payment_date) = '2005-07-30' and p.payment_date = r.rental_date and r.customer_id = c.customer_id and i.inventory_id = r.inventory_id

EXPLAIN ANALYZE
SELECT DISTINCT
    CONCAT(c.last_name, ' ', c.first_name),
    SUM(p.amount) OVER (PARTITION BY c.customer_id, f.title)
FROM payment p, rental r, customer c, inventory i, film f
WHERE date(p.payment_date) = '2005-07-30'
  AND p.payment_date = r.rental_date
  AND r.customer_id = c.customer_id
  AND i.inventory_id = r.inventory_id;

-> Table scan on <temporary>  (cost=2.5..2.5 rows=0) (actual time=13802..13802 rows=391 loops=1)
    -> Temporary table with deduplication  (cost=0..0 rows=0) (actual time=13802..13802 rows=391 loops=1)
        -> Window aggregate with buffering: sum(p.amount) OVER (PARTITION BY c.customer_id,f.title )   (actual time=4820..13576 rows=642000 loops=1)
            -> Sort: c.customer_id, f.title  (actual time=4776..4855 rows=642000 loops=1)
                -> Stream results  (cost=10.1e+6 rows=15.6e+6) (actual time=13.9..2192 rows=642000 loops=1)
                    -> Nested loop inner join  (cost=10.1e+6 rows=15.6e+6) (actual time=12..1905 rows=642000 loops=1)
                        -> Nested loop inner join  (cost=8.51e+6 rows=15.6e+6) (actual time=12..1633 rows=642000 loops=1)
                            -> Nested loop inner join  (cost=6.95e+6 rows=15.6e+6) (actual time=12..1368 rows=642000 loops=1)
                                -> Inner hash join (no condition)  (cost=1.54e+6 rows=15.4e+6) (actual time=12..60.9 rows=634000 loops=1)
                                    -> Filter: (cast(p.payment_date as date) = '2005-07-30')  (cost=1.61 rows=15400) (actual time=0.246..8.79 rows=634 loops=1)
                                        -> Table scan on p  (cost=1.61 rows=15400) (actual time=0.221..5.2 rows=16044 loops=1)
                                    -> Hash
                                        -> Covering index scan on f using idx_title  (cost=103 rows=1000) (actual time=0.0707..0.237 rows=1000 loops=1)
                                -> Covering index lookup on r using rental_date (rental_date = p.payment_date)  (cost=0.25 rows=1.01) (actual time=0.00141..0.00189 rows=1.01 loops=634000)
                            -> Single-row index lookup on c using PRIMARY (customer_id = r.customer_id)  (cost=250e-6 rows=1) (actual time=196e-6..231e-6 rows=1 loops=642000)
                        -> Single-row covering index lookup on i using PRIMARY (inventory_id = r.inventory_id)  (cost=250e-6 rows=1) (actual time=177e-6..212e-6 rows=1 loops=642000)


* Узкие места:

Декартово произведение (Cross Join) — Главная проблема!
В блоке WHERE указана таблица фильмов (film f), но не связана с остальными таблицами (отсутствует условие i.film_id = f.film_id). В результате СУБД делает Cross Join: каждая найденная строка из аренды умножается на все 1000 фильмов из таблицы film. Из нескольких сотен реальных записей за день база искусственно генерирует сотни тысяч строк.

Логическая ошибка в связи таблиц payment и rental.
Условие p.payment_date = r.rental_date некорректно. Время оплаты и время взятия фильма в аренду в реальной жизни (и в базе sakila) могут не совпадать секунда в секунду. Эти таблицы необходимо связывать по их прямому внешнему ключу: p.rental_id = r.rental_id.

Функция DATE() в условии фильтрации.
Выражение WHERE date(p.payment_date) = '2005-07-30' заставляет СУБД применять функцию к каждой строке таблицы payment (Full Table Scan). Даже если на поле payment_date будет индекс, база данных не сможет его использовать.

Неэффективное использование Оконной функции + DISTINCT.
Конструкция DISTINCT ... SUM(p.amount) OVER (PARTITION BY ...) заставляет базу данных сначала развернуть огромный массив строк, рассчитать для каждой строки окно, а затем тратить огромные ресурсы CPU и памяти на сортировку и удаление дубликатов. Здесь гораздо эффективнее работает классическая группировка GROUP BY.

* Оптимизация запроса:

SELECT
    CONCAT(c.last_name, ' ', c.first_name) AS customer_name,
    f.title AS film_title,
    SUM(p.amount) AS total_payment_amount
FROM payment p
INNER JOIN rental r ON p.rental_id = r.rental_id              -- Правильная связь платежа и аренды
INNER JOIN customer c ON r.customer_id = c.customer_id        -- Связь с клиентом
INNER JOIN inventory i ON r.inventory_id = i.inventory_id    -- Связь с инвентарем
INNER JOIN film f ON i.film_id = f.film_id                    -- ИСПРАВЛЕНО: связь с фильмом (убран Cross Join)
WHERE p.payment_date >= '2005-07-30 00:00:00'                 -- ИСПРАВЛЕНО: SARGable условие (индексы работают)
  AND p.payment_date < '2005-07-31 00:00:00'
GROUP BY c.customer_id, c.last_name, c.first_name, f.title;   -- ИСПРАВЛЕНО: эффективная группировка вместо оконной функции

* Добавление индексов:

CREATE INDEX idx_payment_date_amount_rental
ON payment (payment_date, rental_id, amount);

* Результат оптимизации:

-> Table scan on <temporary>  (actual time=9.56..9.72 rows=634 loops=1)
    -> Aggregate using temporary table  (actual time=9.56..9.56 rows=634 loops=1)
        -> Nested loop inner join  (cost=1016 rows=634) (actual time=0.144..7.38 rows=634 loops=1)
            -> Nested loop inner join  (cost=795 rows=634) (actual time=0.135..5.5 rows=634 loops=1)
                -> Nested loop inner join  (cost=573 rows=634) (actual time=0.126..3.83 rows=634 loops=1)
                    -> Nested loop inner join  (cost=351 rows=634) (actual time=0.115..2.22 rows=634 loops=1)
                        -> Filter: ((p.payment_date >= TIMESTAMP'2005-07-30 00:00:00') and (p.payment_date < TIMESTAMP'2005-07-31 00:00:00') and (p.rental_id is not null))  (cost=129 rows=634) (actual time=0.0993..0.582 rows=634 loops=1)
                            -> Covering index range scan on p using idx_payment_date_amount_rental over ('2005-07-30 00:00:00' <= payment_date < '2005-07-31 00:00:00')  (cost=129 rows=634) (actual time=0.0944..0.374 rows=634 loops=1)
                        -> Single-row index lookup on r using PRIMARY (rental_id = p.rental_id)  (cost=0.25 rows=1) (actual time=0.00226..0.00231 rows=1 loops=634)
                    -> Single-row index lookup on c using PRIMARY (customer_id = r.customer_id)  (cost=0.25 rows=1) (actual time=0.00224..0.00229 rows=1 loops=634)
                -> Single-row index lookup on i using PRIMARY (inventory_id = r.inventory_id)  (cost=0.25 rows=1) (actual time=0.00228..0.00233 rows=1 loops=634)
            -> Single-row index lookup on f using PRIMARY (film_id = i.film_id)  (cost=0.25 rows=1) (actual time=0.00262..0.00267 rows=1 loops=634)

## Задание 3

### Типы индексов, уникальные для PostgreSQL (нет в MySQL)

1. GIN (Generalized Inverted Index — Инвертированный индекс)
   * Для чего нужен: Незаменим для работы со сложными составными типами данных, такими как JSONB, хэши (hstore) и массивы.
   * Как работает: Вместо индексации всей строки целиком, GIN индексирует отдельные элементы (ключи и значения внутри JSON, отдельные элементы массива).
   * Аналог в MySQL: Отсутствует. В MySQL для поиска по JSON приходится создавать виртуальные генерируемые колонки и строить B-Tree индекс поверх них.
2. GiST (Generalized Search Tree — Обобщенное дерево поиска)
   * Для чего нужен: Используется для индексации геоданных (PostGIS), диапазонов дат/чисел (типы daterange, int4range) и полнотекстового поиска.
   * Как работает: Позволяет строить структуры для поиска объектов, которые перекрываются, пересекаются или находятся рядом.
   * Аналог в MySQL: В MySQL есть Spatial индекс (R-Tree) только для геометрии, но он не умеет работать с диапазонами данных или кастомными типами, как универсальный GiST.
3. BRIN (Block Range Index — Индекс по диапазонам блоков)
   * Для чего нужен: Разработан специально для огромных таблиц (VLDB — Very Large Databases) с миллиардами строк, где данные упорядочены естественным образом (например, логи логов, даты заказов, таймстампы).
   * Как работает: Вместо индексации каждой строки, BRIN хранит только минимальное и максимальное значение для группы физических блоков диска (например, каждые 128 страниц). Он занимает в сотни раз меньше места, чем B-Tree.
   * Аналог в MySQL: Полностью отсутствует.
4. Hash (Хэш-индексы)
   * Для чего нужен: Оптимизирован исключительно для поиска по точному совпадению (оператор =).
   * Как работает: Превращает значение колонки в хэш-код. Начиная с PostgreSQL 10, эти индексы стали транзакционными (WAL-logged) и полностью безопасными.
   * Аналог в MySQL: Движок InnoDB в MySQL строит скрытые (адаптивные) хэш-индексы в памяти автоматически, но пользователь не может создать хэш-индекс вручную (синтаксис USING HASH в InnoDB молча превращается в обычный B-Tree).
5. SP-GiST (Space-Partitioned GiST — Пространственно-распределенный GiST)
   * Для чего нужен: Эффективен для несимметричных данных, телефонных номеров, IP-адресов, префиксных деревьев (Trie) и текстовых подстрок.
   * Как работает: Разбивает пространство поиска на непересекающиеся области.
   * Аналог в MySQL: Отсутствует.