-- https://dev.mysql.com/doc/refman/8.0/en/select.html

SELECT 1;

SELECT 1+10;

SELECT '1+10';

SELECT (SELECT 10); -- вложенный запрос

-- SELECT * FROM (SELECT 1) 
SELECT * FROM (SELECT 1) AS a;

USE sakila;

SELECT * FROM actor;

SELECT first_name FROM actor;

SELECT first_name, last_name FROM actor;

-- SELECT first_name AS имя, last_name AS фамилия  FROM actor;
SELECT first_name AS `имя`, last_name AS `фамилия`  FROM actor;  
SELECT first_name AS `name`, last_name AS `surname`  FROM actor; 

-- ORDER BY 
-- SELECT first_name AS имя  FROM actor ORDER BY имя ;
SELECT first_name, last_name
FROM actor
ORDER BY first_name; 
 
SELECT first_name, last_name
FROM actor
ORDER BY first_name ASC; 

SELECT first_name, last_name
FROM actor
ORDER BY first_name DESC;

SELECT first_name, last_name
FROM actor
ORDER BY first_name, last_name; -- first_name ASC, last_name ASC

SELECT first_name, last_name
FROM actor
ORDER BY first_name DESC, last_name;

SELECT first_name, last_name
FROM actor
ORDER BY first_name DESC, last_name DESC;
 
SELECT actor_id, first_name, last_name
FROM actor
ORDER BY first_name DESC, last_name DESC;

SELECT actor_id, first_name, last_name
FROM actor 
ORDER BY actor_id;

SELECT actor_id, first_name, last_name
FROM actor; 

SELECT actor_id, first_name, last_name
FROM actor 
ORDER BY actor_id DESC;

-- LIMIT
SELECT first_name, last_name
FROM actor
LIMIT 5;

SELECT actor_id, first_name, last_name
FROM actor
LIMIT 5 offset 2;

SELECT actor_id, first_name, last_name
FROM actor
LIMIT 2, 5;

SELECT actor_id, first_name, last_name
FROM actor
ORDER BY actor_id DESC 
LIMIT 2, 5;

-- DISTINCT
SELECT first_name
FROM actor;

SELECT DISTINCT first_name
FROM actor;

SELECT DISTINCT first_name, last_name
FROM actor;

SELECT DISTINCT actor_id, first_name
FROM actor;

-- WHERE 
SELECT * FROM payment
WHERE amount<1;

SELECT * FROM payment
WHERE amount=0;

SELECT * FROM payment
WHERE amount=0 OR payment_date<'2005-08-01 14:19:48';

SELECT * FROM payment
WHERE amount>1 OR amount<3;

SELECT * FROM payment;

SELECT * FROM payment
WHERE amount>1 AND amount<3;

SELECT * FROM payment
WHERE amount BETWEEN 1 AND 3; -- amount=>1 AND amount<=3 

-- ALIAS
-- SELECT payment_id, amount AS price 
-- FROM payment
-- WHERE price BETWEEN 1 AND 3;

SELECT payment_id, amount AS price 
FROM payment
WHERE amount BETWEEN 1 AND 3
ORDER BY price;

-- CAST 
SELECT payment_id, amount, payment_date 
FROM payment;

SELECT 
	payment_id, 
	amount, 
	CAST(payment_date AS DATE) AS 'дата' 
FROM payment;

-- операции с числами
SELECT 1+1;
SELECT 3/3;
SELECT 3*3;
SELECT 3-3;
SELECT 3/0; -- NULL

SELECT ROUND(100.576);

SELECT ROUND(100.576,2);
SELECT TRUNCATE(100.576,2);

SELECT FLOOR(100.576); -- до меньшего
SELECT CEIL(100.576); -- до большего

SELECT ABS(-100.576); 

SELECT POWER(2, 8);
SELECT POWER(5, 2);

SELECT SQRT(25);
SELECT POWER(25, 1/2);
SELECT POWER(8, 1/3); -- 1.9999
SELECT ROUND(POWER(8, 1/3)); -- 2

SELECT 1024 DIV 100;
SELECT 1024 % 100;

SELECT LEAST(17, 6, 13, 7, 1, 189);
SELECT GREATEST(17, 6, 13, 7, 1, 189);

SELECT RAND();
SELECT RAND()*100;
SELECT * FROM actor ORDER BY RAND(); 

-- работа со строками
SELECT CONCAT(first_name, last_name)  AS `user` FROM actor;
SELECT CONCAT(first_name, ' ', last_name) AS `user` FROM actor;
SELECT CONCAT('Пользователь: ', first_name, ' ', last_name)  AS `user` FROM actor;
SELECT CONCAT_WS(' ', 'Пользователь:', first_name, last_name) AS `user` FROM actor;

SELECT 
	first_name, 
	LENGTH(first_name) 
FROM actor;

SELECT 
	first_name, 
	LENGTH(first_name),
	CHAR_LENGTH(first_name) 
FROM actor;

SELECT  
	LENGTH('привет'),
	CHAR_LENGTH('привет'); 

SELECT 
	last_name,
	POSITION('D' IN last_name)
FROM actor; 

SELECT 
	last_name,
	POSITION('d' IN last_name)
FROM actor; 

SELECT 
	last_name,
	SUBSTR(last_name, 2, 3) -- нумерация символов с 1
FROM actor; 

SELECT 
	last_name,
	LEFT(last_name, 3),
	RIGHT(last_name, 3) 
FROM actor; 

SELECT 
	last_name,
	LOWER(last_name),
	UPPER(last_name) 
FROM actor;
SELECT 	UPPER('Привет');

SELECT
	INSERT(last_name, 1, 1, 'Пользователь: ')
FROM actor; 

SELECT
	REPLACE(last_name, 'A', 'X')
FROM actor; 

SELECT
	REPLACE(last_name, 'А', 'X') -- почему не меняет?
FROM actor;

SELECT '         ivan@netology.ru    ',
       TRIM('         ivan@netology.ru    ');

SELECT first_name 
FROM actor
WHERE first_name LIKE 'M%'; 

SELECT first_name 
FROM actor
WHERE first_name LIKE '_EN%'; 

SELECT first_name 
FROM actor
WHERE first_name LIKE '%EN%';

SELECT first_name 
FROM actor
WHERE first_name LIKE '______';

-- работа с датами и временем
SELECT NOW();
SELECT CURDATE(); 
SELECT CURRENT_TIMESTAMP;

SELECT NOW(), DATE_ADD(NOW(), INTERVAL 3 DAY);

SELECT NOW(), DATE_SUB(NOW(), INTERVAL 3 DAY);

SELECT NOW(), YEAR(NOW()), MONTH(NOW()), WEEK(NOW()), 
	DAY(NOW()), HOUR(NOW()), MINUTE(NOW()), SECOND(NOW());

SELECT NOW(), EXTRACT(HOUR FROM NOW()), HOUR(NOW());

SELECT DATEDIFF(last_update, payment_date)
FROM payment; 

SELECT TIMESTAMPDIFF(YEAR, payment_date, last_update)
FROM payment; 

SELECT TIMESTAMPDIFF(DAY, payment_date, last_update)
FROM payment;

SELECT TIMESTAMPDIFF(MINUTE, payment_date, last_update)
FROM payment;

SELECT DATE_FORMAT(payment_date, '%d-%a-%m-%Y') FROM payment; 

SELECT TIME_FORMAT(payment_date, '%H:%i:%s') FROM payment; 

SELECT DATE_FORMAT(payment_date, '%d-%m-%Y %H:%i:%s') FROM payment; 

