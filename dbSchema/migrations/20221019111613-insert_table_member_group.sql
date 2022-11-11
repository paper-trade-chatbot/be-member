
-- +migrate Up
ALTER TABLE `member_group` AUTO_INCREMENT = 1;
INSERT INTO 
	`member_group`(`id`,`name`)
VALUES
	(1,'default'),
	(2,'system'),
	(3,'company'),
	(4,'admin'),
	(5,'agent'),
	(6,'member'),
	(7,'member_suppressed'),
	(8,'member_freezed'),
	(9,'member_disabled'),
	(10,'test');

-- +migrate Down

SET SQL_SAFE_UPDATES = 0;
DELETE FROM `member_group`;
SET SQL_SAFE_UPDATES = 1;
