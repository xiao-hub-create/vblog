CREATE TABLE `users` (
    `id` bigint unsigned AUTO_INCREMENT,
    `created_at` datetime(3) NOT NULL,
    `create_by` varchar(200),
    `updated_at` datetime(3) NULL,
    `username` varchar(191),
    `password` varchar(255),
    `avatar` varchar(255),
    `nic_name` varchar(100),
    `email` varchar(100),
    `block_at` datetime(3) NULL,
    `block_reason` text,
    PRIMARY KEY (`id`),
    INDEX `idx_users_username` (`username`),
    CONSTRAINT `uni_users_username` UNIQUE (`username`)
)
