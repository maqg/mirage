-- MySQL dump 10.13  Distrib 5.6.22, for osx10.10 (x86_64)
--
-- Host: localhost    Database: dbmirage
-- ------------------------------------------------------
-- Server version	5.6.22

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `tb_account`
--

DROP TABLE IF EXISTS `tb_account`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_account` (
  `ID` varchar(36) NOT NULL DEFAULT '',
  `U_State` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1: OK, 0: Bad',
  `U_Type` int(11) NOT NULL DEFAULT '7' COMMENT '0x1: super,0x10 admin,0x100 audit',
  `U_LastLogin` bigint(20) NOT NULL DEFAULT '0',
  `U_CreateTime` bigint(20) NOT NULL DEFAULT '0',
  `U_LastSync` bigint(20) NOT NULL DEFAULT '0',
  `U_Name` varchar(128) NOT NULL DEFAULT '',
  `U_Email` varchar(128) NOT NULL DEFAULT '',
  `U_PhoneNumber` varchar(32) NOT NULL DEFAULT '',
  `U_Password` varchar(128) NOT NULL DEFAULT '',
  `U_Description` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `tb_account_id` (`ID`),
  KEY `tb_account_state` (`U_State`),
  KEY `tb_account_name` (`U_Name`),
  KEY `tb_account_type` (`U_Type`),
  KEY `tb_account_email` (`U_Email`),
  KEY `tb_account_phonenumber` (`U_PhoneNumber`),
  KEY `tb_account_password` (`U_Password`),
  KEY `tb_account_createtime` (`U_CreateTime`),
  KEY `tb_account_lastlogin` (`U_LastLogin`),
  KEY `tb_account_lastsync` (`U_LastSync`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_account`
--

LOCK TABLES `tb_account` WRITE;
/*!40000 ALTER TABLE `tb_account` DISABLE KEYS */;
INSERT INTO `tb_account` VALUES ('00000000000000000000000000000000',1,7,0,0,0,'root','','','26c501d0a825f48acc76ef6c784cdacc',''),('42fa1e66ff5411e6b2dc60334b213917',1,7,0,0,0,'admin','','','292f137f691469948acd0e72b141e373','');
/*!40000 ALTER TABLE `tb_account` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_misc`
--

DROP TABLE IF EXISTS `tb_misc`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_misc` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `M_Name` varchar(64) NOT NULL DEFAULT '',
  `M_Value` varchar(64) NOT NULL DEFAULT '',
  `M_Type` varchar(64) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `tb_misc_id` (`ID`),
  KEY `tb_misc_name` (`M_Name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_misc`
--

LOCK TABLES `tb_misc` WRITE;
/*!40000 ALTER TABLE `tb_misc` DISABLE KEYS */;
INSERT INTO `tb_misc` VALUES (1,'hostname','Mirage','common');
/*!40000 ALTER TABLE `tb_misc` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_relusergroup`
--

DROP TABLE IF EXISTS `tb_relusergroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_relusergroup` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `RUG_GroupId` varchar(36) NOT NULL DEFAULT '',
  `RUG_UserId` varchar(36) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `tb_relusergroup_id` (`ID`),
  KEY `tb_relusergroup_groupid` (`RUG_GroupId`),
  KEY `tb_relusergroup_userid` (`RUG_UserId`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_relusergroup`
--

LOCK TABLES `tb_relusergroup` WRITE;
/*!40000 ALTER TABLE `tb_relusergroup` DISABLE KEYS */;
INSERT INTO `tb_relusergroup` VALUES (1,'00000000000000000000000000000000','e1132175ff5411e6913460334b213917'),(2,'00000000000000000000000000000000','e1132175ff5411e6913460334b213918'),(3,'00000000000000000000000000000000','e1132175ff5411e6913460334b213919');
/*!40000 ALTER TABLE `tb_relusergroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_session`
--

DROP TABLE IF EXISTS `tb_session`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_session` (
  `ID` varchar(36) NOT NULL DEFAULT '',
  `S_UserId` varchar(36) NOT NULL DEFAULT '',
  `S_UserName` varchar(128) NOT NULL DEFAULT '',
  `S_Cookie` varchar(1024) NOT NULL DEFAULT '',
  `S_CreateTime` bigint(20) NOT NULL DEFAULT '0',
  `S_LastSync` bigint(20) NOT NULL DEFAULT '0',
  `S_ExpireTime` bigint(20) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`),
  KEY `tb_session_id` (`ID`),
  KEY `tb_session_userid` (`S_UserId`),
  KEY `tb_session_username` (`S_UserName`),
  KEY `tb_session_createtime` (`S_CreateTime`),
  KEY `tb_session_lastsync` (`S_LastSync`),
  KEY `tb_session_expiretime` (`S_ExpireTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_session`
--

LOCK TABLES `tb_session` WRITE;
/*!40000 ALTER TABLE `tb_session` DISABLE KEYS */;
INSERT INTO `tb_session` VALUES ('00000000000000000000000000000000','00000000000000000000000000000000','system','{}',1483921186,1483921186,1799281186);
/*!40000 ALTER TABLE `tb_session` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user` (
  `ID` varchar(36) NOT NULL DEFAULT '',
  `U_Name` varchar(128) NOT NULL DEFAULT '',
  `U_State` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1: OK, 0: Bad',
  `U_Type` int(11) NOT NULL DEFAULT '1' COMMENT 'terminal user',
  `U_LastLogin` bigint(20) NOT NULL DEFAULT '0',
  `U_LastSync` bigint(20) NOT NULL DEFAULT '0',
  `U_CreateTime` bigint(20) NOT NULL DEFAULT '0',
  `U_Password` varchar(128) NOT NULL DEFAULT '',
  `U_Email` varchar(128) NOT NULL DEFAULT '',
  `U_PhoneNumber` varchar(32) NOT NULL DEFAULT '',
  `U_Description` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `tb_user_id` (`ID`),
  KEY `tb_user_state` (`U_State`),
  KEY `tb_user_name` (`U_Name`),
  KEY `tb_user_type` (`U_Type`),
  KEY `tb_user_email` (`U_Email`),
  KEY `tb_user_phonenumber` (`U_PhoneNumber`),
  KEY `tb_user_password` (`U_Password`),
  KEY `tb_user_createtime` (`U_CreateTime`),
  KEY `tb_user_lastlogin` (`U_LastLogin`),
  KEY `tb_user_lastsync` (`U_LastSync`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user`
--

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;
INSERT INTO `tb_user` VALUES ('e1132175ff5411e6913460334b213917','test',1,1,0,0,0,'050e7255a41627d5f9141cd9bea94357','','',''),('e1132175ff5411e6913460334b213918','test1',1,1,0,0,0,'8555f41950a8fe627b1fb13a50bcc26a','','',''),('e1132175ff5411e6913460334b213919','test2',1,1,0,0,0,'6c274b6b947fb3c65de6eb1a548e65cf','','','');
/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_usergroup`
--

DROP TABLE IF EXISTS `tb_usergroup`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_usergroup` (
  `ID` varchar(36) NOT NULL DEFAULT '',
  `UG_Name` varchar(128) NOT NULL DEFAULT '',
  `UG_AccountId` varchar(36) NOT NULL DEFAULT '',
  `UG_CreateTime` bigint(20) NOT NULL DEFAULT '0',
  `UG_LastSync` bigint(20) NOT NULL DEFAULT '0',
  `UG_Description` varchar(1024) NOT NULL DEFAULT '',
  PRIMARY KEY (`ID`),
  KEY `tb_usergroup_id` (`ID`),
  KEY `tb_usergroup_name` (`UG_Name`),
  KEY `tb_usergroup_accountid` (`UG_AccountId`),
  KEY `tb_usergroup_createtime` (`UG_CreateTime`),
  KEY `tb_usergroup_lastsync` (`UG_LastSync`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_usergroup`
--

LOCK TABLES `tb_usergroup` WRITE;
/*!40000 ALTER TABLE `tb_usergroup` DISABLE KEYS */;
INSERT INTO `tb_usergroup` VALUES ('00000000000000000000000000000000','default','00000000000000000000000000000000',0,0,'');
/*!40000 ALTER TABLE `tb_usergroup` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Temporary view structure for view `v_account`
--

DROP TABLE IF EXISTS `v_account`;
/*!50001 DROP VIEW IF EXISTS `v_account`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `v_account` AS SELECT 
 1 AS `Id`,
 1 AS `Name`,
 1 AS `State`,
 1 AS `Type`,
 1 AS `LastLogin`,
 1 AS `LastSync`,
 1 AS `CreateTime`,
 1 AS `Password`,
 1 AS `Email`,
 1 AS `PhoneNumber`,
 1 AS `Description`*/;
SET character_set_client = @saved_cs_client;

--
-- Temporary view structure for view `v_user`
--

DROP TABLE IF EXISTS `v_user`;
/*!50001 DROP VIEW IF EXISTS `v_user`*/;
SET @saved_cs_client     = @@character_set_client;
SET character_set_client = utf8;
/*!50001 CREATE VIEW `v_user` AS SELECT 
 1 AS `Id`,
 1 AS `Name`,
 1 AS `State`,
 1 AS `Type`,
 1 AS `LastLogin`,
 1 AS `LastSync`,
 1 AS `CreateTime`,
 1 AS `Password`,
 1 AS `Email`,
 1 AS `PhoneNumber`,
 1 AS `Description`*/;
SET character_set_client = @saved_cs_client;

--
-- Dumping routines for database 'dbmirage'
--

--
-- Final view structure for view `v_account`
--

/*!50001 DROP VIEW IF EXISTS `v_account`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_account` AS select `tb_account`.`ID` AS `Id`,`tb_account`.`U_Name` AS `Name`,`tb_account`.`U_State` AS `State`,`tb_account`.`U_Type` AS `Type`,`tb_account`.`U_LastLogin` AS `LastLogin`,`tb_account`.`U_LastSync` AS `LastSync`,`tb_account`.`U_CreateTime` AS `CreateTime`,`tb_account`.`U_Password` AS `Password`,`tb_account`.`U_Email` AS `Email`,`tb_account`.`U_PhoneNumber` AS `PhoneNumber`,`tb_account`.`U_Description` AS `Description` from `tb_account` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;

--
-- Final view structure for view `v_user`
--

/*!50001 DROP VIEW IF EXISTS `v_user`*/;
/*!50001 SET @saved_cs_client          = @@character_set_client */;
/*!50001 SET @saved_cs_results         = @@character_set_results */;
/*!50001 SET @saved_col_connection     = @@collation_connection */;
/*!50001 SET character_set_client      = utf8 */;
/*!50001 SET character_set_results     = utf8 */;
/*!50001 SET collation_connection      = utf8_general_ci */;
/*!50001 CREATE ALGORITHM=UNDEFINED */
/*!50013 DEFINER=`root`@`localhost` SQL SECURITY DEFINER */
/*!50001 VIEW `v_user` AS select `tb_user`.`ID` AS `Id`,`tb_user`.`U_Name` AS `Name`,`tb_user`.`U_State` AS `State`,`tb_user`.`U_Type` AS `Type`,`tb_user`.`U_LastLogin` AS `LastLogin`,`tb_user`.`U_LastSync` AS `LastSync`,`tb_user`.`U_CreateTime` AS `CreateTime`,`tb_user`.`U_Password` AS `Password`,`tb_user`.`U_Email` AS `Email`,`tb_user`.`U_PhoneNumber` AS `PhoneNumber`,`tb_user`.`U_Description` AS `Description` from `tb_user` */;
/*!50001 SET character_set_client      = @saved_cs_client */;
/*!50001 SET character_set_results     = @saved_cs_results */;
/*!50001 SET collation_connection      = @saved_col_connection */;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-03-11 23:14:23
