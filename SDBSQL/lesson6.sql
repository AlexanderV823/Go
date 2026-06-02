-- POSTGRESQL PARTITION
-- https://www.postgresql.org/docs/current/ddl-partitioning.html
-- Range Partitioning 
-- List Partitioning 
-- Hash Partitioning 
 
-- создадим базу part_pg 
 
CREATE TABLE measurement (
    city_id         int not null,
    logdate         date not null,
	peaktemp        int
) PARTITION BY RANGE (logdate);

CREATE TABLE measurement_y2023 PARTITION OF measurement
    FOR VALUES FROM ('2023-01-01') TO ('2024-01-01');
	
CREATE TABLE measurement_y2024 PARTITION OF measurement
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');

CREATE TABLE measurement_y2025 PARTITION OF measurement
    FOR VALUES FROM ('2025-01-01') TO ('2026-01-01');	

CREATE TABLE measurement_default PARTITION OF measurement DEFAULT;

INSERT INTO measurement (city_id, logdate, peaktemp)
VALUES
(1, '2022-01-15', 0),
(1, '2023-01-15', -10),
(1, '2024-01-15', -5),
(1, '2025-01-15', -19),
(1, '2026-01-15', 5),
(2, '2022-01-15', 12),
(2, '2023-01-15', 12),
(2, '2024-01-15', 5),
(2, '2025-01-15', 18),
(2, '2026-01-15', 20);

 DROP TABLE measurement_y2023;
 -- DROP TABLE measurement_y2024;
 ALTER TABLE measurement DETACH PARTITION measurement_y2024;
 ALTER TABLE measurement ATTACH PARTITION measurement_y2024
    FOR VALUES FROM ('2024-01-01') TO ('2025-01-01');
	
-- MYSQL PARTITION	
-- https://dev.mysql.com/doc/mysql-partitioning-excerpt/8.0/en/
-- Chapter 3 Partitioning Types

-- KEY
DROP TABLE IF EXISTS members;

CREATE TABLE members (
	id SERIAL PRIMARY KEY,
	joined DATE
)
PARTITION BY KEY()
PARTITIONS 6;

INSERT INTO members (joined)
WITH RECURSIVE DateRange AS (
    SELECT '1999-01-01' AS dt -- Начало диапазона
    UNION ALL
    SELECT DATE_ADD(dt, INTERVAL ROUND(rand()*100) DAY)
    FROM DateRange
    WHERE dt < '2026-01-31' -- Конец диапазона
)
SELECT dt FROM DateRange;

-- HASH
DROP TABLE IF EXISTS members;

CREATE TABLE members (
	id int,
	joined DATE
)
PARTITION BY HASH(YEAR(joined))
PARTITIONS 6;

INSERT INTO members (id, joined)
WITH RECURSIVE DateRange AS (
    SELECT '1999-01-01' AS dt -- Начало диапазона
    UNION ALL
    SELECT DATE_ADD(dt, INTERVAL ROUND(rand()*100) DAY)
    FROM DateRange
    WHERE dt < '2026-01-31' -- Конец диапазона
)
SELECT ROUND(rand()*10), dt FROM DateRange;

SELECT * FROM members PARTITION (p1);
SELECT * FROM members WHERE YEAR(joined) = 1999; 
EXPLAIN SELECT * FROM members WHERE YEAR(joined) = 1999;

-- RANGE
DROP TABLE IF EXISTS members;

-- БУДЕТ ОШИБКА
CREATE TABLE members (
	id SERIAL PRIMARY KEY,
	joined DATE
)
PARTITION BY RANGE(YEAR(joined)) (
	PARTITION y199x VALUES LESS THAN (2000),
	PARTITION y200x VALUES LESS THAN (2010),
	PARTITION y201x VALUES LESS THAN  (2020),
	PARTITION y202x VALUES LESS THAN  (2030)
);

CREATE TABLE members (
	id SERIAL PRIMARY KEY,
	joined DATE
)
PARTITION BY RANGE(id) (
	PARTITION p1 VALUES LESS THAN (100),
	PARTITION p2 VALUES LESS THAN (200),
	PARTITION p3 VALUES LESS THAN  (300),
	PARTITION pd VALUES LESS THAN  MAXVALUE
);


DROP TABLE IF EXISTS members;

CREATE TABLE members (
	id INT,
	joined DATE
)
PARTITION BY RANGE(YEAR(joined)) (
	PARTITION y199x VALUES LESS THAN (2000),
	PARTITION y200x VALUES LESS THAN (2010),
	PARTITION y201x VALUES LESS THAN  (2020),
	PARTITION y202x VALUES LESS THAN  (2030)
);

INSERT INTO members (id, joined)
WITH RECURSIVE DateRange AS (
    SELECT '1999-01-01' AS dt -- Начало диапазона
    UNION ALL
    SELECT DATE_ADD(dt, INTERVAL ROUND(rand()*100) DAY)
    FROM DateRange
    WHERE dt < '2026-01-31' -- Конец диапазона
)
SELECT ROUND(rand()*10), dt FROM DateRange;

SELECT * FROM members WHERE YEAR(joined) = 2009;
EXPLAIN SELECT * FROM members WHERE YEAR(joined) = 2009;
EXPLAIN SELECT * FROM members PARTITION (y200x)  WHERE YEAR(joined) = 2009;

