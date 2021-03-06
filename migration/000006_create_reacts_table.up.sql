CREATE TABLE `reacts` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `image_id` int(11) UNSIGNED NOT NULL,
    `user_id` int(11) UNSIGNED NOT NULL,
    `type_id` int(11) UNSIGNED NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES `users` (`id`),
    FOREIGN KEY (`image_id`) REFERENCES `images` (`id`),
    FOREIGN KEY (`type_id`) REFERENCES `react_types` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
