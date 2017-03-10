DROP TABLE IF EXISTS `tb_misc`;
CREATE TABLE `tb_misc` (
		`ID` BIGINT NOT NULL AUTO_INCREMENT,
		`M_Name` VARCHAR(64) NOT NULL DEFAULT '',
		`M_Value` VARCHAR(64) NOT NULL DEFAULT '',
		`M_Type` VARCHAR(64) NOT NULL DEFAULT '',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_misc ADD INDEX tb_misc_id (ID);
ALTER TABLE tb_misc ADD INDEX tb_misc_name (M_Name);

DROP TABLE IF EXISTS `tb_account`;
CREATE TABLE `tb_account` (
		`ID` VARCHAR(36) NOT NULL DEFAULT '',
		`U_State` TINYINT NOT NULL DEFAULT '1' COMMENT '1: OK, 0: Bad',
		`U_Type` INTEGER NOT NULL DEFAULT '7' COMMENT '0x1: super,0x10 admin,0x100 audit',
		`U_LastLogin` BIGINT NOT NULL DEFAULT '0',
		`U_CreateTime` BIGINT NOT NULL DEFAULT '0',
		`U_LastSync` BIGINT NOT NULL DEFAULT '0',
		`U_Name` VARCHAR(128) NOT NULL DEFAULT '',
		`U_Email` VARCHAR(128) NOT NULL DEFAULT '',
		`U_PhoneNumber` VARCHAR(32) NOT NULL DEFAULT '',
		`U_Password` VARCHAR(128) NOT NULL DEFAULT '',
		`U_Description` VARCHAR(1024) NOT NULL DEFAULT '',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_account ADD INDEX tb_account_id (ID);
ALTER TABLE tb_account ADD INDEX tb_account_state (U_State);
ALTER TABLE tb_account ADD INDEX tb_account_name (U_Name);
ALTER TABLE tb_account ADD INDEX tb_account_type (U_Type);
ALTER TABLE tb_account ADD INDEX tb_account_email (U_Email);
ALTER TABLE tb_account ADD INDEX tb_account_phonenumber (U_PhoneNumber);
ALTER TABLE tb_account ADD INDEX tb_account_password (U_Password);
ALTER TABLE tb_account ADD INDEX tb_account_createtime (U_CreateTime);
ALTER TABLE tb_account ADD INDEX tb_account_lastlogin (U_LastLogin);
ALTER TABLE tb_account ADD INDEX tb_account_lastsync (U_LastSync);

DROP TABLE IF EXISTS `tb_relusergroup`;
CREATE TABLE `tb_relusergroup` (
		`ID` BIGINT NOT NULL AUTO_INCREMENT,
		`RUG_GroupId` VARCHAR(36) NOT NULL DEFAULT '',
		`RUG_UserId` VARCHAR(36) NOT NULL DEFAULT '',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_relusergroup ADD INDEX tb_relusergroup_id (ID);
ALTER TABLE tb_relusergroup ADD INDEX tb_relusergroup_groupid (RUG_GroupId);
ALTER TABLE tb_relusergroup ADD INDEX tb_relusergroup_userid (RUG_UserId);

DROP TABLE IF EXISTS `tb_usergroup`;
CREATE TABLE `tb_usergroup` (
		`ID` VARCHAR(36) NOT NULL DEFAULT '',
		`UG_Name` VARCHAR(128) NOT NULL DEFAULT '',
		`UG_AccountId` VARCHAR(36) NOT NULL DEFAULT '',
		`UG_CreateTime` BIGINT NOT NULL DEFAULT '0',
		`UG_LastSync` BIGINT NOT NULL DEFAULT '0',
		`UG_Description` VARCHAR(1024) NOT NULL DEFAULT '',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_usergroup ADD INDEX tb_usergroup_id (ID);
ALTER TABLE tb_usergroup ADD INDEX tb_usergroup_name (UG_Name);
ALTER TABLE tb_usergroup ADD INDEX tb_usergroup_accountid (UG_AccountId);
ALTER TABLE tb_usergroup ADD INDEX tb_usergroup_createtime (UG_CreateTime);
ALTER TABLE tb_usergroup ADD INDEX tb_usergroup_lastsync (UG_LastSync);

DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
		`ID` VARCHAR(36) NOT NULL DEFAULT '',
		`U_Name` VARCHAR(128) NOT NULL DEFAULT '',
		`U_State` TINYINT NOT NULL DEFAULT '1' COMMENT '1: OK, 0: Bad',
		`U_Type` INTEGER NOT NULL DEFAULT '1' COMMENT 'terminal user',
		`U_LastLogin` BIGINT NOT NULL DEFAULT '0',
		`U_LastSync` BIGINT NOT NULL DEFAULT '0',
		`U_CreateTime` BIGINT NOT NULL DEFAULT '0',
		`U_Password` VARCHAR(128) NOT NULL DEFAULT '',
		`U_Email` VARCHAR(128) NOT NULL DEFAULT '',
		`U_PhoneNumber` VARCHAR(32) NOT NULL DEFAULT '',
		`U_Description` VARCHAR(1024) NOT NULL DEFAULT '',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_user ADD INDEX tb_user_id (ID);
ALTER TABLE tb_user ADD INDEX tb_user_state (U_State);
ALTER TABLE tb_user ADD INDEX tb_user_name (U_Name);
ALTER TABLE tb_user ADD INDEX tb_user_type (U_Type);
ALTER TABLE tb_user ADD INDEX tb_user_email (U_Email);
ALTER TABLE tb_user ADD INDEX tb_user_phonenumber (U_PhoneNumber);
ALTER TABLE tb_user ADD INDEX tb_user_password (U_Password);
ALTER TABLE tb_user ADD INDEX tb_user_createtime (U_CreateTime);
ALTER TABLE tb_user ADD INDEX tb_user_lastlogin (U_LastLogin);
ALTER TABLE tb_user ADD INDEX tb_user_lastsync (U_LastSync);

DROP TABLE IF EXISTS `tb_session`;
CREATE TABLE `tb_session` (
		`ID` VARCHAR(36) NOT NULL DEFAULT '',
		`S_UserId` VARCHAR(36) NOT NULL DEFAULT '',
		`S_UserName` VARCHAR(128) NOT NULL DEFAULT '',
		`S_Cookie` VARCHAR(1024) NOT NULL DEFAULT '',
		`S_CreateTime` BIGINT NOT NULL DEFAULT '0',
		`S_LastSync` BIGINT NOT NULL DEFAULT '0',
		`S_ExpireTime` BIGINT NOT NULL DEFAULT '0',
		PRIMARY KEY (`ID`)
) ENGINE=Innodb DEFAULT CHARSET=utf8;
ALTER TABLE tb_session ADD INDEX tb_session_id (ID);
ALTER TABLE tb_session ADD INDEX tb_session_userid (S_UserId);
ALTER TABLE tb_session ADD INDEX tb_session_username (S_UserName);
ALTER TABLE tb_session ADD INDEX tb_session_createtime (S_CreateTime);
ALTER TABLE tb_session ADD INDEX tb_session_lastsync (S_LastSync);
ALTER TABLE tb_session ADD INDEX tb_session_expiretime (S_ExpireTime);

