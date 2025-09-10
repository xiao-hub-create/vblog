REATE TABLE `blogs` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NOT NULL,
    `create_by` varchar(200),
    `updated_at` datetime(3) NULL,
    `title` varchar(200),
    `summary` text,
    `content` text,
    `category` varchar(200),
    `tags` longtext,
    PRIMARY KEY (`id`),
    INDEX `idx_blogs_category` (`category`)
)
