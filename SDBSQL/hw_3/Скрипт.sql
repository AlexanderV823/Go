SELECT DISTINCT district 
FROM sakila.address 
WHERE district LIKE 'K%a' 
  AND district NOT LIKE '% %';

SELECT payment_id, customer_id, staff_id, rental_id, amount, payment_date
FROM sakila.payment
WHERE payment_date >= '2005-06-15 00:00:00' 
  AND payment_date < '2005-06-19 00:00:00'
  AND amount > 10.00;

SELECT rental_id, rental_date, inventory_id, customer_id, return_date, staff_id
FROM sakila.rental
ORDER BY rental_date DESC
LIMIT 5;

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

SELECT 
    email AS original_email,
    SUBSTRING_INDEX(email, '@', 1) AS email_username,
    SUBSTRING_INDEX(email, '@', -1) AS email_domain
FROM 
    sakila.customer;

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

