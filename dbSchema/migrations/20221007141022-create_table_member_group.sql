
-- +migrate Up
CREATE TABLE `member_group` (
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(36) NOT NULL COMMENT '名稱',
    `memo` VARCHAR(128) NULL COMMENT '備註',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '創建時間',

     UNIQUE (`name`),
     PRIMARY KEY (`id`)
) AUTO_INCREMENT=1 CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='會員組別';

-- +migrate Down
SET FOREIGN_KEY_CHECKS=0;
DROP TABLE IF EXISTS `member_group`;
