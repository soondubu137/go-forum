CREATE TABLE `post` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `post_id` bigint(20) NOT NULL,
    `title` varchar(255) COLLATE utf8mb4_general_ci NOT NULL,
    `content` text COLLATE utf8mb4_general_ci NOT NULL,
    `author_id` bigint(20) NOT NULL,
    `community_id` int(11) NOT NULL,
    `status` tinyint(4) NOT NULL DEFAULT '1',
    `created_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `post_id` (`post_id`),
    KEY `idx_author_id` (`author_id`),
    KEY `idx_community_id` (`community_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
