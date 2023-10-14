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


--  Create a table which holds the relationship between user and event
CREATE TABLE IF NOT EXISTS `user_events` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `event_id` int(11) NOT NULL,
  `role` ENUM('ADMIN','OWNER','ATTENDEE','CONTRIBUTOR') NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
  FOREIGN KEY (`event_id`) REFERENCES `events`(`id`)
  UNIQUE KEY user_event_unique (user_id, event_id) 
);
CREATE TABLE
    IF NOT EXISTS `expenses` (
        `id` int(11) NOT NULL AUTO_INCREMENT,
        `event_id` int(11) NOT NULL,
        `item_name` varchar(255) NOT NULL,
        `cost` int(11) NOT NULL,
        `description` TEXT,
        `category` ENUM(
            'VENUE',
            'CATERING',
            'DECORATION',
        ) NOT NULL,
        `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`id`),
        FOREIGN KEY (`event_id`) REFERENCES `events`(`id`)
    );


CREATE TABLE IF NOT EXISTS `event_sessions`(
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `event_id` int(11) NOT NULL,
    `name` varchar(255) NOT NULL,
    `start_time` DATETIME NOT NULL,
    `end_time` DATETIME NOT NULL,
    `description` TEXT,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`event_id`) REFERENCES `events`(`id`)
);
    
