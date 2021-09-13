CREATE TABLE `images`(
                             `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
                             `gallery_id` int(11) UNSIGNED NOT NULL,
                             `title` varchar(255) NOT NULL,
                             `image_url` varchar(255) NOT NULL,
                             `created_at` datetime NOT NULL,
                             `updated_at` datetime NOT NULL,
                             `deleted_at` datetime DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
