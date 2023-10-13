-- Active: 1697094688271@@127.0.0.1@3306@event_management_golang
-- Create database if not EXISTS named 'event_management_golang'
CREATE DATABASE IF NOT EXISTS event_management;

USE event_management;

ALTER TABLE event ADD COLUMN description VARCHAR(255);

-- Create table named User if not EXISTS
CREATE TABLE IF NOT EXISTS `user` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `description` varchar(255),
  `phone_number` VARCHAR (20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `email` (`email`)
);


-- Create table named Event if not EXISTS
CREATE TABLE IF NOT EXISTS `event` (
    -- In real database use UUID
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `start_date` DATETIME NOT NULL,
  `end_date` DATETIME NOT NULL,
  `location` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);