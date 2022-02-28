-- Adminer 4.8.1 MySQL 8.0.28 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

CREATE DATABASE `workshop_docker` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;
USE `workshop_docker`;

DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `username` varchar(40) NOT NULL,
  `email_adress` varchar(60) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `users` (`username`, `email_adress`) VALUES
('toto',	'toto@toto.fr'),
('tata',	'tata@tata.fr');

-- 2022-02-28 12:59:29