
-- +migrate Up
INSERT INTO 
	`member`(`account`,`password_hash`,`mail`,`status`,`verify_status`,`group_id`)
VALUES
	('system','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','system@ptbot.com.tw',1,3,2),
	('company','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','company@ptbot.com.tw',1,3,3),
	('admin01','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','admin01@ptbot.com.tw',1,3,4),
	('admin02','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','admin02@ptbot.com.tw',1,3,4),
	('admin03','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','admin03@ptbot.com.tw',1,3,4),
	('admin04','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','admin04@ptbot.com.tw',1,3,4),
	('admin05','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','admin05@ptbot.com.tw',1,3,4),
	('josh01','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','josh01@ptbot.com.tw',1,3,6),
	('josh02','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','josh02@ptbot.com.tw',1,3,7),
	('josh03','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','josh03@ptbot.com.tw',1,3,8),
	('josh04','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','josh04@ptbot.com.tw',1,3,9),
	('test01','$2a$04$6DmLuDQY.CxZ/POa74AIBe6u3MBz9.0W1E7oyDxVcf50ZBkhM5zFK','test01@ptbot.com.tw',1,3,10);


-- +migrate Down

DELETE FROM `member` WHERE `account` = 'system';
DELETE FROM `member` WHERE `account` = 'company';
DELETE FROM `member` WHERE `account` = 'admin01';
DELETE FROM `member` WHERE `account` = 'admin02';
DELETE FROM `member` WHERE `account` = 'admin03';
DELETE FROM `member` WHERE `account` = 'admin04';
DELETE FROM `member` WHERE `account` = 'admin05';
DELETE FROM `member` WHERE `account` = 'josh01';
DELETE FROM `member` WHERE `account` = 'josh02';
DELETE FROM `member` WHERE `account` = 'josh03';
DELETE FROM `member` WHERE `account` = 'josh04';
DELETE FROM `member` WHERE `account` = 'test01';
