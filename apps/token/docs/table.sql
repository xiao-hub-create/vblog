REATE TABLE `tokens` (
    `id` bigint unsigned AUTO_INCREMENT,
    `ref_user_id` bigint unsigned,
    `issue_at` datetime(3) NULL,
    `access_token` varchar(191),
    `access_token_expire_at` datetime(3) NULL,
    `refresh_token` varchar(191),
    `refresh_token_expire_at` datetime(3) NULL,
    PRIMARY KEY (`id`),
    INDEX `idx_tokens_access_token` (`access_token`),
    INDEX `idx_tokens_refresh_token` (`refresh_token`),
    CONSTRAINT `uni_tokens_access_token` UNIQUE (`access_token`),
    CONSTRAINT `uni_tokens_refresh_token` UNIQUE (`refresh_token`)
) 
