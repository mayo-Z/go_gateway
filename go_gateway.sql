-- MySQL dump 10.13  Distrib 8.0.21, for Win64 (x86_64)
--
-- Host: localhost    Database: go_gateway
-- ------------------------------------------------------
-- Server version	8.0.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `gateway_admin`
--

DROP TABLE IF EXISTS `gateway_admin`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_admin` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='管理员表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_admin`
--

LOCK TABLES `gateway_admin` WRITE;
/*!40000 ALTER TABLE `gateway_admin` DISABLE KEYS */;
INSERT INTO `gateway_admin` VALUES (1,'admin','admin','2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3','2020-04-10 16:42:05','2021-09-30 16:37:49',0);
/*!40000 ALTER TABLE `gateway_admin` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_app`
--

DROP TABLE IF EXISTS `gateway_app`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_app` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `app_id` varchar(255) NOT NULL DEFAULT '' COMMENT '租户id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '租户名称',
  `secret` varchar(255) NOT NULL DEFAULT '' COMMENT '密钥',
  `white_ips` varchar(1000) NOT NULL DEFAULT '' COMMENT 'ip白名单，支持前缀匹配',
  `qpd` bigint NOT NULL DEFAULT '0' COMMENT '日请求量限制',
  `qps` bigint NOT NULL DEFAULT '0' COMMENT '每秒请求量限制',
  `create_at` datetime NOT NULL COMMENT '添加时间',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除 1=删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COMMENT='网关租户表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_app`
--

LOCK TABLES `gateway_app` WRITE;
/*!40000 ALTER TABLE `gateway_app` DISABLE KEYS */;
INSERT INTO `gateway_app` VALUES (31,'app_id_a','租户A','449441eb5e72dca9c42a12f3924ea3a2','white_ips',100000,100,'2020-04-15 20:55:02','2020-04-21 07:23:34',0),(32,'app_id_b','租户B','8d7b11ec9be0e59a36b52f32366c09cb','',100,0,'2020-04-15 21:40:52','2022-05-16 18:03:48',0),(33,'app_id','租户名称','','',0,0,'2020-04-15 22:02:23','2020-04-15 22:06:51',1),(34,'app_id45','名称','07d980f8a49347523ee1d5c1c41aec02','',0,0,'2020-04-15 22:06:38','2020-04-15 22:06:49',1),(35,'app_id_1','app_id_35','8272b17ee5dd781e1c39eea075db9c9e','',5000,500,'2021-09-26 17:01:49','2021-09-26 17:06:02',1),(36,'test_app_2','test_app_2','4df7f28b881cca3bc48f043f872fcc39','',0,0,'2021-10-08 15:38:21','2021-10-08 15:38:40',1);
/*!40000 ALTER TABLE `gateway_app` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_access_control`
--

DROP TABLE IF EXISTS `gateway_service_access_control`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_access_control` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务id',
  `open_auth` tinyint NOT NULL DEFAULT '0' COMMENT '是否开启权限 1=开启',
  `black_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '黑名单ip',
  `white_list` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单ip',
  `white_host_name` varchar(1000) NOT NULL DEFAULT '' COMMENT '白名单主机',
  `clientip_flow_limit` int NOT NULL DEFAULT '0' COMMENT '客户端ip限流',
  `service_flow_limit` int NOT NULL DEFAULT '0' COMMENT '服务端限流',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=205 DEFAULT CHARSET=utf8 COMMENT='网关权限控制表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_access_control`
--

LOCK TABLES `gateway_service_access_control` WRITE;
/*!40000 ALTER TABLE `gateway_service_access_control` DISABLE KEYS */;
INSERT INTO `gateway_service_access_control` VALUES (162,35,1,'','','',0,0),(165,34,0,'','','',0,0),(167,36,0,'','','',0,0),(168,38,1,'111.11','22.33','11.11',12,12),(169,41,1,'111.11','22.33','11.11',12,12),(170,42,1,'111.11','22.33','11.11',12,12),(171,43,0,'111.11','22.33','11.11',12,12),(172,44,0,'','','',0,0),(173,45,0,'','','',0,0),(174,46,0,'','','',0,0),(175,47,0,'','','',0,0),(176,48,0,'','','',0,0),(177,49,0,'','','',0,0),(178,50,0,'','','',0,0),(179,51,0,'','','',0,0),(180,52,0,'','','',0,0),(181,53,0,'','','',0,0),(182,54,1,'127.0.0.3','127.0.0.2','',11,12),(183,55,1,'127.0.0.2','127.0.0.1','',45,34),(184,56,0,'192.168.1.0','','',0,0),(185,57,0,'','','',0,0),(186,58,1,'','','',0,0),(187,59,1,'127.0.0.1','','',2,3),(188,60,1,'','','',0,0),(189,61,0,'','','',0,0),(190,62,0,'','','',0,0),(191,63,0,'','','',0,0),(192,64,0,'','','',0,0),(193,65,0,'','','',0,0),(194,66,0,'','','',0,0),(195,67,1,'','127.0.0.1,127.0.0.2','',0,0),(196,68,0,'','127.0.0.1','',100,1000),(197,69,1,'127.0.0.2','127.0.0.1','',100,1000),(198,70,1,'127.0.0.2','127.0.0.1','',100,1000),(199,71,1,'127.0.0.2','127.0.0.1','',100,1000),(200,72,1,'127.0.0.32','127.0.0.1','',100,1000),(201,73,1,'127.0.0.2','127.0.0.1','',100,1000),(202,74,0,'','','',0,0),(203,75,0,'','','',0,0),(204,76,1,'','','',0,0);
/*!40000 ALTER TABLE `gateway_service_access_control` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_grpc_rule`
--

DROP TABLE IF EXISTS `gateway_service_grpc_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_grpc_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务id',
  `port` int NOT NULL DEFAULT '0' COMMENT '端口',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=177 DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_grpc_rule`
--

LOCK TABLES `gateway_service_grpc_rule` WRITE;
/*!40000 ALTER TABLE `gateway_service_grpc_rule` DISABLE KEYS */;
INSERT INTO `gateway_service_grpc_rule` VALUES (171,53,8009,''),(172,54,8002,'add metadata1 datavalue,edit metadata2 datavalue2'),(173,58,8012,'add meta_name meta_value'),(174,63,8714,'add h1 v1'),(175,64,8715,''),(176,73,8095,'add k1 u1');
/*!40000 ALTER TABLE `gateway_service_grpc_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_http_rule`
--

DROP TABLE IF EXISTS `gateway_service_http_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_http_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL COMMENT '服务id',
  `rule_type` tinyint NOT NULL DEFAULT '0' COMMENT '匹配类型 0=url前缀url_prefix 1=域名domain ',
  `rule` varchar(255) NOT NULL DEFAULT '' COMMENT 'type=domain表示域名，type=url_prefix时表示url前缀',
  `need_https` tinyint NOT NULL DEFAULT '0' COMMENT '支持https 1=支持',
  `need_strip_uri` tinyint NOT NULL DEFAULT '0' COMMENT '启用strip_uri 1=启用',
  `need_websocket` tinyint NOT NULL DEFAULT '0' COMMENT '是否支持websocket 1=支持',
  `url_rewrite` varchar(5000) NOT NULL DEFAULT '' COMMENT 'url重写功能 格式：^/gatekeeper/test_service(.*) $1 多个逗号间隔',
  `header_transfor` varchar(5000) NOT NULL DEFAULT '' COMMENT 'header转换支持增加(add)、删除(del)、修改(edit) 格式: add headname headvalue 多个逗号间隔',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=187 DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_http_rule`
--

LOCK TABLES `gateway_service_http_rule` WRITE;
/*!40000 ALTER TABLE `gateway_service_http_rule` DISABLE KEYS */;
INSERT INTO `gateway_service_http_rule` VALUES (165,35,1,'',0,0,0,'',''),(168,34,0,'',0,0,0,'',''),(170,36,0,'',0,0,0,'',''),(171,38,0,'/abc',1,0,1,'^/abc $1','add head1 value1'),(172,43,0,'/usr',1,1,0,'^/afsaasf $1,^/afsaasf $1',''),(173,44,1,'www.test.com',1,1,1,'',''),(174,47,1,'www.test.com',1,1,1,'',''),(175,48,1,'www.test.com',1,1,1,'',''),(176,49,1,'www.test.com',1,1,1,'',''),(177,56,0,'/test_http_service',1,1,1,'^/test_http_service/abb/(.*) /test_http_service/bba/$1','add header_name header_value'),(178,59,1,'test.com',0,1,1,'','add headername headervalue'),(179,60,0,'/test_strip_uri',0,1,0,'^/aaa/(.*) /bbb/$1',''),(180,61,0,'/test_https_server',1,1,0,'',''),(181,62,0,'/test_http_service_indb',0,0,0,'',''),(182,66,0,'/testtest2',0,0,0,'^/test_http_service/abb/(.*) ,/test_http_service/bba/$1','add h1 v1'),(183,67,0,'/testhttp3',1,1,1,'^/testhttp3/abb/(.*) /testhttp3/bba/$1','add h1 v1'),(184,74,0,'test_http_string',0,1,0,'',''),(185,75,0,'/test_http_string',0,0,0,'',''),(186,76,0,'/test_http_string3',0,0,0,'^/test_http_string3/abbb(.*) test_http_string3/baaa$1','add go_gateway 1.0.0,del Accept 111');
/*!40000 ALTER TABLE `gateway_service_http_rule` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_info`
--

DROP TABLE IF EXISTS `gateway_service_info`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_info` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `load_type` tinyint NOT NULL DEFAULT '0' COMMENT '负载类型 0=http 1=tcp 2=grpc',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称 6-128 数字字母下划线',
  `service_desc` varchar(255) NOT NULL DEFAULT '' COMMENT '服务描述',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '添加时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint DEFAULT '0' COMMENT '是否删除 1=删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=77 DEFAULT CHARSET=utf8 COMMENT='网关基本信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_info`
--

LOCK TABLES `gateway_service_info` WRITE;
/*!40000 ALTER TABLE `gateway_service_info` DISABLE KEYS */;
INSERT INTO `gateway_service_info` VALUES (34,0,'websocket_test','websocket_test','2020-04-13 01:31:47','1971-01-01 00:00:00',1),(35,1,'test_grpc','test_grpc','2020-04-13 01:34:32','1971-01-01 00:00:00',1),(36,2,'test_httpe','test_httpe','2020-04-11 21:12:48','1971-01-01 00:00:00',1),(38,0,'service_name','11111','2020-04-15 07:49:45','2020-04-11 23:59:39',1),(41,0,'service_name_tcp','11111','2020-04-13 01:38:01','2020-04-12 01:06:09',1),(42,0,'service_name_tcp2','11111','2020-04-13 01:38:06','2020-04-12 01:13:24',1),(43,1,'service_name_tcp4','service_name_tcp4','2020-04-15 07:49:44','2020-04-12 01:13:50',1),(44,0,'websocket_service','websocket_service','2020-04-15 07:49:43','2020-04-13 01:20:08',1),(45,1,'tcp_service','tcp_desc','2020-04-15 07:49:41','2020-04-13 01:46:27',1),(46,1,'grpc_service','grpc_desc','2020-04-13 01:54:12','2020-04-13 01:53:14',1),(47,0,'testsefsafs','werrqrr','2020-04-13 01:59:14','2020-04-13 01:57:49',1),(48,0,'testsefsafs1','werrqrr','2020-04-13 01:59:11','2020-04-13 01:58:14',1),(49,0,'testsefsafs1222','werrqrr','2020-04-13 01:59:08','2020-04-13 01:58:23',1),(50,2,'grpc_service_name','grpc_service_desc','2020-04-15 07:49:40','2020-04-13 02:01:00',1),(51,2,'gresafsf','wesfsf','2020-04-15 07:49:39','2020-04-13 02:01:57',1),(52,2,'gresafsf11','wesfsf','2020-04-13 02:03:41','2020-04-13 02:02:55',1),(53,2,'tewrqrw111','123313','2020-04-13 02:03:38','2020-04-13 02:03:20',1),(54,2,'test_grpc_service1','test_grpc_service1','2020-04-15 07:49:37','2020-04-15 07:38:43',1),(55,1,'test_tcp_service1','redis服务代理','2020-04-15 07:49:35','2020-04-15 07:46:35',1),(56,0,'test_http_service','测试HTTP代理','2020-04-16 00:54:45','2022-05-18 08:13:44',0),(57,1,'test_tcp_service','测试TCP代理','2020-04-19 14:03:09','2022-05-09 20:52:49',0),(58,2,'test_grpc_service','测试GRPC服务','2020-04-21 07:20:16','2022-05-16 17:42:38',0),(59,0,'test.com:8080','测试域名接入','2020-04-18 22:54:14','2020-04-18 20:29:13',0),(60,0,'test_strip_uri','测试路径接入','2020-04-21 06:55:26','2022-05-18 08:23:03',0),(61,0,'test_https_server','测试https服务','2020-04-19 12:22:33','2021-09-17 16:33:51',1),(62,0,'test_http_service_indb','test_http_service_indb','2021-09-23 03:21:36','2021-09-23 03:21:36',0),(63,2,'test_grpc_service_3','test_grpc_service_3','2021-09-25 15:22:03','2021-09-25 15:40:42',0),(64,2,'test_grpc_service_4','test_grpc_service_3','2021-09-25 15:22:47','2021-10-01 19:37:11',1),(65,1,'test_tcp_service_3','test_tcp_service_003','2021-09-25 15:48:23','2022-05-09 20:27:35',1),(66,0,'testtest2','testtest2','2021-10-03 17:04:36','2021-10-07 13:09:31',0),(67,0,'testhttp3','testhttp3','2021-10-07 11:05:18','2022-04-30 17:21:08',0),(68,1,'tcp_test_service_1','tcp_test_service_1','2021-10-08 09:05:16','2021-10-08 09:07:45',1),(69,1,'tcp_test_service_2','tcp_test_service_2','2021-10-08 09:09:36','2021-10-08 09:34:06',1),(70,1,'tcp_test_service_4','tcp_test_service_1','2021-10-08 09:35:37','2022-05-09 20:27:23',1),(71,1,'tcp_test_service_5','tcp_test_service_1','2021-10-08 09:40:11','2022-05-09 20:27:12',1),(72,1,'grpc_test_service_1','grpc_test_service_1','2021-10-08 09:47:10','2022-05-09 20:27:18',1),(73,2,'grpc_test_service_2','grpc_test_service_2','2021-10-08 09:57:30','2021-10-08 09:59:57',0),(76,0,'test_http_string3','test_http_string3','2021-11-14 14:57:15','2022-05-05 15:47:08',0);
/*!40000 ALTER TABLE `gateway_service_info` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_load_balance`
--

DROP TABLE IF EXISTS `gateway_service_load_balance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_load_balance` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL DEFAULT '0' COMMENT '服务id',
  `check_method` tinyint NOT NULL DEFAULT '0' COMMENT '检查方法 0=tcpchk,检测端口是否握手成功',
  `check_timeout` int NOT NULL DEFAULT '0' COMMENT 'check超时时间,单位s',
  `check_interval` int NOT NULL DEFAULT '0' COMMENT '检查间隔, 单位s',
  `round_type` tinyint NOT NULL DEFAULT '2' COMMENT '轮询方式 0=random 1=round-robin 2=weight_round-robin 3=ip_hash',
  `ip_list` varchar(2000) NOT NULL DEFAULT '' COMMENT 'ip列表',
  `weight_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '权重列表',
  `forbid_list` varchar(2000) NOT NULL DEFAULT '' COMMENT '禁用ip列表',
  `upstream_connect_timeout` int NOT NULL DEFAULT '0' COMMENT '建立连接超时, 单位s',
  `upstream_header_timeout` int NOT NULL DEFAULT '0' COMMENT '获取header超时, 单位s',
  `upstream_idle_timeout` int NOT NULL DEFAULT '0' COMMENT '链接最大空闲时间, 单位s',
  `upstream_max_idle` int NOT NULL DEFAULT '0' COMMENT '最大空闲链接数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=205 DEFAULT CHARSET=utf8 COMMENT='网关负载表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_load_balance`
--

LOCK TABLES `gateway_service_load_balance` WRITE;
/*!40000 ALTER TABLE `gateway_service_load_balance` DISABLE KEYS */;
INSERT INTO `gateway_service_load_balance` VALUES (162,35,0,2000,5000,2,'127.0.0.1:50051','100','',10000,0,0,0),(165,34,0,2000,5000,2,'100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072','50,50,50,80','',20000,20000,10000,100),(167,36,0,2000,5000,2,'100.90.164.31:8072,100.90.163.51:8072,100.90.163.52:8072,100.90.165.32:8072','50,50,50,80','100.90.164.31:8072,100.90.163.51:8072',10000,10000,10000,100),(168,38,0,0,0,1,'111:111,22:111','11,11','111',1111,111,222,333),(169,41,0,0,0,1,'111:111,22:111','11,11','111',0,0,0,0),(170,42,0,0,0,1,'111:111,22:111','11,11','111',0,0,0,0),(171,43,0,2,5,1,'111:111,22:111','11,11','',1111,2222,333,444),(172,44,0,2,5,2,'127.0.0.1:8076','50','',0,0,0,0),(173,45,0,2,5,2,'127.0.0.1:88','50','',0,0,0,0),(174,46,0,2,5,2,'127.0.0.1:8002','50','',0,0,0,0),(175,47,0,2,5,2,'12777:11','11','',0,0,0,0),(176,48,0,2,5,2,'12777:11','11','',0,0,0,0),(177,49,0,2,5,2,'12777:11','11','',0,0,0,0),(178,50,0,2,5,2,'127.0.0.1:8001','50','',0,0,0,0),(179,51,0,2,5,2,'1212:11','50','',0,0,0,0),(180,52,0,2,5,2,'1212:11','50','',0,0,0,0),(181,53,0,2,5,2,'1111:11','111','',0,0,0,0),(182,54,0,2,5,1,'127.0.0.1:80','50','',0,0,0,0),(183,55,0,2,5,3,'127.0.0.1:81','50','',0,0,0,0),(184,56,0,2,5,2,'127.0.0.1:20003,127.0.0.1:20004','50,50','',0,0,0,0),(185,57,0,2,5,2,'127.0.0.1:7002','50','',0,0,0,0),(186,58,0,2,5,2,'127.0.0.1:50055','50','',0,0,0,0),(187,59,0,2,5,2,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0),(188,60,0,2,5,2,'127.0.0.1:20003,127.0.0.1:20004','50,50','',0,0,0,0),(189,61,0,2,5,2,'127.0.0.1:3003,127.0.0.1:3004','50,50','',0,0,0,0),(190,62,0,0,0,0,'127.0.0.1:80','50','',0,0,0,0),(191,63,0,0,0,0,'127.0.0.1:2003','51','',0,0,0,0),(192,64,0,0,0,0,'127.0.0.1:2003','51','',0,0,0,0),(193,65,0,0,0,0,'127.0.0.1:2003','51','',0,0,0,0),(194,66,0,0,0,0,'127.0.0.1:8881','50','',0,0,0,0),(195,67,0,0,0,3,'127.0.0.1:20003,127.0.0.1:20004','50,50','',0,0,0,0),(196,68,0,0,0,2,'127.0.0.1:1003,127.0.0.1:1004','40,50','',0,0,0,0),(197,69,0,0,0,2,'127.0.0.1:1003,127.0.0.1:1004','50,60','',0,0,0,0),(198,70,0,0,0,2,'127.0.0.1:1003,127.0.0.1:1004','50,60','',0,0,0,0),(199,71,0,0,0,2,'127.0.0.1:1003,127.0.0.1:1004','50,60','',0,0,0,0),(200,72,0,0,0,2,'127.0.0.1:1003,127.0.0.1:1004','50,60','',0,0,0,0),(201,73,0,0,0,0,'127.0.0.1:1003,127.0.0.1:1004','50,60','',0,0,0,0),(202,74,0,0,0,0,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0),(203,75,0,0,0,0,'127.0.0.1:2003,127.0.0.1:2004','50,50','',0,0,0,0),(204,76,0,0,0,3,'127.0.0.1:20003,127.0.0.1:20004','50,50','',10,5,0,0);
/*!40000 ALTER TABLE `gateway_service_load_balance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `gateway_service_tcp_rule`
--

DROP TABLE IF EXISTS `gateway_service_tcp_rule`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `gateway_service_tcp_rule` (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `service_id` bigint NOT NULL COMMENT '服务id',
  `port` int NOT NULL DEFAULT '0' COMMENT '端口号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=188 DEFAULT CHARSET=utf8 COMMENT='网关路由匹配表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `gateway_service_tcp_rule`
--

LOCK TABLES `gateway_service_tcp_rule` WRITE;
/*!40000 ALTER TABLE `gateway_service_tcp_rule` DISABLE KEYS */;
INSERT INTO `gateway_service_tcp_rule` VALUES (171,41,8002),(172,42,8003),(173,43,8004),(174,38,8004),(175,45,8001),(176,46,8005),(177,50,8006),(178,51,8007),(179,52,8008),(180,55,8010),(181,57,8011),(182,65,8724),(183,68,8013),(184,69,8099),(185,70,8098),(186,71,8097),(187,72,8096);
/*!40000 ALTER TABLE `gateway_service_tcp_rule` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-05-20 18:33:05
