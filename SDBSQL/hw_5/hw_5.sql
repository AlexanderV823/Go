-- Задача 1

SELECT 
    ROUND(SUM(index_length) / SUM(data_length + index_length) * 100, 2) AS index_to_total_percentage
FROM information_schema.tables 
WHERE table_schema = 'sakila' AND table_type = 'BASE TABLE';

-- Задача 2

EXPLAIN ANALYZE
SELECT DISTINCT 
    CONCAT(c.last_name, ' ', c.first_name), 
    SUM(p.amount) OVER (PARTITION BY c.customer_id, f.title)
FROM payment p, rental r, customer c, inventory i, film f
WHERE date(p.payment_date) = '2005-07-30' 
  AND p.payment_date = r.rental_date 
  AND r.customer_id = c.customer_id 
  AND i.inventory_id = r.inventory_id;

-- Добавление индексов и оптимизация

CREATE INDEX idx_payment_date_amount_rental 
ON payment (payment_date, rental_id, amount);

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



