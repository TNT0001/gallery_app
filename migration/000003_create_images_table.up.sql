CREATE TABLE IF NOT EXISTS `images`(
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `gallery_id` int(11) UNSIGNED NOT NULL,
    `user_id` int(11) UNSIGNED NOT NULL ,
    `title` varchar(255) NOT NULL,
    `image_url` varchar(255) NOT NULL,
    `image_uuid` varchar(255) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`gallery_id`) REFERENCES `galleries`(`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
