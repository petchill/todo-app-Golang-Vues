CREATE DATABASE IF NOT EXISTS `todo_app`;
USE `todo_app`;
CREATE TABLE IF NOT EXISTS `todos`(
    `id` INT PRIMARY KEY AUTO_INCREMENT,
    `topic` VARCHAR(255) NOT NULL,
    `describtion` VARCHAR(255) NOT NULL,
    `status` VARCHAR(255) DEFAULT 'pending',
    `deadline` DATE,
    `color` VARCHAR(255) DEFAULT '#ffffff',
    `created_by` INT,
    `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CHECK (
        `status` IN ('pending', 'inprocess', 'done', 'abandoned')
    ),
    FOREIGN KEY (created_by) REFERENCES users(id)
);