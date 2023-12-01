CREATE DATABASE tisea;

USE tisea;

CREATE TABLE `tisea_users` (
	`id` INT unsigned NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(20) NOT NULL,
	`password_hash` TEXT NOT NULL,
	`email` VARCHAR(32) NOT NULL,
	`level` INT unsigned NOT NULL DEFAULT 0,
	`bio` VARCHAR(50) DEFAULT "",
	`nickname` VARCHAR(20) DEFAULT "",
	`group_id` INT unsigned NOT NULL DEFAULT 0,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
);