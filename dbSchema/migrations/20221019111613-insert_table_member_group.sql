
-- +migrate Up
ALTER TABLE `member_group` AUTO_INCREMENT = 1;
INSERT INTO 
	`member_group`(`name`)
VALUES
	('default'),
	('system');

-- +migrate Down

SET SQL_SAFE_UPDATES = 0;
DELETE FROM `member_group`;
SET SQL_SAFE_UPDATES = 1;
