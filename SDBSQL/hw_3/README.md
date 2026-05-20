# Домашнее задание к занятию "SQL. Часть 1"

## Задание 1

Получите уникальные названия районов из таблицы с адресами, которые начинаются на “K” и заканчиваются на “a” и не содержат пробелов.

SELECT DISTINCT district
FROM sakila.address
WHERE district LIKE 'K%a'
  AND district NOT LIKE '% %';


## Задание 2

Получите из таблицы платежей за прокат фильмов информацию по платежам, которые выполнялись в промежуток с 15 июня 2005 года по 18 июня 2005 года включительно и стоимость которых превышает 10.00.

SELECT payment_id, customer_id, staff_id, rental_id, amount, payment_date
FROM sakila.payment
WHERE payment_date >= '2005-06-15 00:00:00'
  AND payment_date < '2005-06-19 00:00:00'
  AND amount > 10.00;

## Задание 3

Получите последние пять аренд фильмов.

SELECT rental_id, rental_date, inventory_id, customer_id, return_date, staff_id
FROM sakila.rental
ORDER BY rental_date DESC
LIMIT 5;

## Задание 4

Одним запросом получите активных покупателей, имена которых Kelly или Willie.

Сформируйте вывод в результат таким образом:

все буквы в фамилии и имени из верхнего регистра переведите в нижний регистр,
замените буквы 'll' в именах на 'pp'.

SELECT
    customer_id,
    REPLACE(LOWER(first_name), 'll', 'pp') AS processed_first_name,
    LOWER(last_name) AS processed_last_name,
    email,
    active
FROM
    sakila.customer
WHERE
    active = 1
    AND (first_name = 'KELLY' OR first_name = 'WILLIE');

## Задание 5

Выведите Email каждого покупателя, разделив значение Email на две отдельных колонки: в первой колонке должно быть значение, указанное до @, во второй — значение, указанное после @.

SELECT
    email AS original_email,
    SUBSTRING_INDEX(email, '@', 1) AS email_username,
    SUBSTRING_INDEX(email, '@', -1) AS email_domain
FROM
    sakila.customer;

## Задание 6

Доработайте запрос из предыдущего задания, скорректируйте значения в новых колонках: первая буква должна быть заглавной, остальные — строчными.

SELECT
    email AS original_email,

    CONCAT(
        UPPER(LEFT(SUBSTRING_INDEX(email, '@', 1), 1)),
        LOWER(SUBSTRING(SUBSTRING_INDEX(email, '@', 1), 2))
    ) AS formatted_username,

    CONCAT(
        UPPER(LEFT(SUBSTRING_INDEX(email, '@', -1), 1)),
        LOWER(SUBSTRING(SUBSTRING_INDEX(email, '@', -1), 2))
    ) AS formatted_domain

FROM
    sakila.customer;