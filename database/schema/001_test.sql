CREATE DATABASE IF NOT EXISTS `todo_app` CHARACTER SET utf8mb4 COLLATE utf8mb4 \ _unicode \ _ci;
USE `todo_app`;
CREATE TABLE IF NOT EXISTS `test` (
    `id` int PRIMARY KEY AUTO \ _INCREMENT,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4 \ _thai \ _520 \ _w2,
    `owner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4 \ _thai \ _520 \ _w2,
    `created_at` TIMESTAMP DEFAULT CURRENT \ _TIMESTAMP,
    `updated_at` TIMESTAMP ON UPDATE CURRENT \ _TIMESTAMP,
    INDEX (name)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4 \ _unicode \ _ci;