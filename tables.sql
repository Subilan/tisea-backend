CREATE DATABASE tisea;

USE tisea;

CREATE TABLE `tisea_users` (
	`username` VARCHAR(20) NOT NULL,
	`email` VARCHAR(32) NOT NULL,
	`password_hash` TEXT NOT NULL,
	`nickname` VARCHAR(20) DEFAULT "",
	`bio` VARCHAR(50) DEFAULT "",
	`id` INT unsigned NOT NULL AUTO_INCREMENT,
	`created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
	`group_id` INT unsigned NOT NULL DEFAULT 0,
	`level` INT unsigned NOT NULL DEFAULT 0,
	PRIMARY KEY (`id`)
);