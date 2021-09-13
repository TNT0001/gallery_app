CREATE TABLE IF NOT EXISTS `friends` (
    `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL,
    `friend_id` int(11) NOT NULL,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL,
    `deleted_at` datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
