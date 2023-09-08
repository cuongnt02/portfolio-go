CREATE DATABASE jforum CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE jforum;

CREATE TABLE threads (
    id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100) NOT NULL,
    created_time DATETIME NOT NULL,
    updated_time DATETIME NOT NULL
);

-- THREADS TEST DATA INSERT
INSERT INTO threads (name, created_time, updated_time) VALUES (
    'Juice',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 1 DAY)
);

INSERT INTO threads (name, created_time, updated_time) VALUES (
    'Soft Drink',
    UTC_TIMESTAMP(),
    DATE_ADD(UTC_TIMESTAMP(), INTERVAL 2 DAY)
);

INSERT INTO threads (name, created_time, updated_time) VALUES (
    'Beer',
    UTC_TIMESTAMP(),
    UTC_TIMESTAMP()
);

CREATE INDEX idx_threads_created ON threads(created)
