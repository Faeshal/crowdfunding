# ************************************************************
# Sequel Ace SQL dump
# Version 20035
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# Host: 103.147.32.15 (MySQL 5.5.5-10.3.34-MariaDB-0ubuntu0.20.04.1)
# Database: crowdfunding
# Generation Time: 2022-08-19 8:49:37 AM +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table campaign_images
# ------------------------------------------------------------

DROP TABLE IF EXISTS `campaign_images`;

CREATE TABLE `campaign_images` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `campaign_id` bigint(20) DEFAULT NULL,
  `file_name` longtext DEFAULT NULL,
  `is_primary` bigint(20) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_campaigns_campaign_images` (`campaign_id`),
  CONSTRAINT `fk_campaigns_campaign_images` FOREIGN KEY (`campaign_id`) REFERENCES `campaigns` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table campaigns
# ------------------------------------------------------------

DROP TABLE IF EXISTS `campaigns`;

CREATE TABLE `campaigns` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `name` longtext DEFAULT NULL,
  `short_description` longtext DEFAULT NULL,
  `description` longtext DEFAULT NULL,
  `perks` longtext DEFAULT NULL,
  `backer_count` bigint(20) DEFAULT NULL,
  `goal_amount` bigint(20) DEFAULT NULL,
  `current_amount` bigint(20) DEFAULT NULL,
  `slug` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_campaigns_user` (`user_id`),
  CONSTRAINT `fk_campaigns_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

LOCK TABLES `campaigns` WRITE;
/*!40000 ALTER TABLE `campaigns` DISABLE KEYS */;

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`)
VALUES
	(1,1,'villa bogor','short desc','long desc','passive income,free lunch,free parking',1,100000,20000,'villa-bogor-1','2022-08-18 09:10:38.647','2022-08-18 09:22:32.019');

/*!40000 ALTER TABLE `campaigns` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table transactions
# ------------------------------------------------------------

DROP TABLE IF EXISTS `transactions`;

CREATE TABLE `transactions` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `campaign_id` int(11) DEFAULT NULL,
  `user_id` int(11) DEFAULT NULL,
  `amount` int(11) DEFAULT NULL,
  `status` varchar(255) DEFAULT NULL,
  `code` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `payment_url` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

LOCK TABLES `transactions` WRITE;
/*!40000 ALTER TABLE `transactions` DISABLE KEYS */;

INSERT INTO `transactions` (`id`, `campaign_id`, `user_id`, `amount`, `status`, `code`, `created_at`, `updated_at`, `payment_url`)
VALUES
	(16,1,2,20000,'paid','','2022-08-18 09:21:43','2022-08-18 09:22:31','https://app.sandbox.midtrans.com/snap/v3/redirection/ee473fb8-d812-4718-b31a-f43e9c72f286'),
	(17,1,2,20000,'pending','','2022-08-19 08:25:23','2022-08-19 08:25:26','https://app.sandbox.midtrans.com/snap/v3/redirection/92a1d0be-7016-470c-8f25-7f22eafa9d2c');

/*!40000 ALTER TABLE `transactions` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table users
# ------------------------------------------------------------

DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` longtext DEFAULT NULL,
  `occupation` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `password_hash` longtext DEFAULT NULL,
  `avatar_file_name` longtext DEFAULT NULL,
  `role` longtext DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`)
VALUES
	(1,'faeshal','jakarta','faeshal@mail.com','$2a$04$wIb0C/o4lQhTEJk3vwuOauiPMviofOQDgSq/yymJf5uyBTi0td3fW','','user','2022-08-18 04:42:26.663','2022-08-18 04:42:26.663'),
	(2,'zein','jakarta','zein@mail.com','$2a$04$k44WpHBlH6PifGrYM/O5d.d/jn92obSuz0odDEWP2RZk4VzjrFN/W','','user','2022-08-18 09:10:52.208','2022-08-18 09:10:52.208');

/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
