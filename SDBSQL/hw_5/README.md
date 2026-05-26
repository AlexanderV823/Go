# Домашнее задание к занятию "Индексы"

## Задание 1

Напишите запрос к учебной базе данных, который вернёт процентное отношение общего размера всех индексов к общему размеру всех таблиц.

SELECT
    ROUND(SUM(index_length) / SUM(data_length + index_length) * 100, 2) AS index_to_total_percentage
FROM information_schema.tables
WHERE table_schema = 'sakila' AND table_type = 'BASE TABLE';