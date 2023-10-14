-- Active: 1697094688271@@127.0.0.1@3306@event_management_golang
-- Create database if not EXISTS named 'event_management_golang'
DROP DATABASE event_management;
CREATE DATABASE IF NOT EXISTS event_management;

USE event_management;

-- Create table named User if not EXISTS
CREATE TABLE IF NOT EXISTS `users` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone_number` VARCHAR (20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);


-- Create table named Event if not EXISTS
CREATE TABLE IF NOT EXISTS `events` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `start_date` DATETIME NOT NULL,
  `end_date` DATETIME NOT NULL,
  `location` varchar(255) NOT NULL,
  `description` varchar(255),
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

DROP TABLE user_events;

--  Create a table which holds the relationship between user and event
CREATE TABLE IF NOT EXISTS `user_events` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `event_id` int(11) NOT NULL,
  `role` VARCHAR(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`event_id`) REFERENCES `events`(`id`)
);


SELECT
    `id`,
    `email`,
    `phone_number`
FROM `users`
WHERE (
        `email` = 'prameshkarkiss0656s17@gmail.com'
    )

SELECT
    `users`.`id`,
    `users`.`email`,
    `users`.`phone_number`,
    `user_events`.`role`,
    `user_events`.`event_id`
FROM `users`
    INNER JOIN `user_events` ON (
        `users`.`id` = `user_events`.`user_id`
    )
WHERE (`user_events`.`event_id` = '1')