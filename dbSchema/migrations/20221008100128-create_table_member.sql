
-- +migrate Up
CREATE TABLE `member` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `account` VARCHAR(30) NOT NULL COMMENT '帳號',
    `password_hash` VARCHAR(128) NOT NULL COMMENT '密碼雜湊',
    `mail` VARCHAR(30) NOT NULL COMMENT '信箱',
    `line_id` VARCHAR(30) COMMENT 'line帳號',
    `country_code` VARCHAR(4) COMMENT '手機國碼',
    `phone` VARCHAR(10) COMMENT '手機號碼',
    `status` TINYINT(4) NOT NULL COMMENT '帳號狀態',
    `verify_status` TINYINT(4) NOT NULL COMMENT '驗證狀態',
    `group_id` BIGINT UNSIGNED NOT NULL COMMENT '組別id',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新時間',
    `deleted_at` TIMESTAMP NULL COMMENT '刪除時間',
    
    UNIQUE INDEX (`mail`),
    UNIQUE INDEX (`account`),
    FOREIGN KEY (`group_id`) REFERENCES member_group(`id`) ON DELETE CASCADE,
    PRIMARY KEY(`id`)
) AUTO_INCREMENT=1 CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='會員資料';

-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `member`;
