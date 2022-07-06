/*
 Navicat MySQL Data Transfer

 Source Server         : t
 Source Server Type    : MySQL
 Source Server Version : 80029
 Source Host           : localhost:3306
 Source Schema         : youthcamp

 Target Server Type    : MySQL
 Target Server Version : 80029
 File Encoding         : 65001

 Date: 12/06/2022 21:53:57
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `content` varchar(50) DEFAULT NULL,
  `create_date` date NOT NULL,
  `token` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=163033780 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Table structure for favorite
-- ----------------------------
DROP TABLE IF EXISTS `favorite`;
CREATE TABLE `favorite` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `video_id` int NOT NULL,
  `token` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=162756974 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户的视频点赞表';

-- ----------------------------
-- Records of favorite
-- ----------------------------
BEGIN;
INSERT INTO `favorite` VALUES (156300010, 154347356, 156083940, '154347356');
COMMIT;

-- ----------------------------
-- Table structure for relations
-- ----------------------------
DROP TABLE IF EXISTS `relations`;
CREATE TABLE `relations` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `to_user_id` int NOT NULL,
  `token` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=163206162 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of relations
-- ----------------------------
BEGIN;
INSERT INTO `relations` VALUES (156373620, 163493274, 160238754, '163493274');
INSERT INTO `relations` VALUES (156724370, 154347356, 153969272, '154347356');
INSERT INTO `relations` VALUES (157446564, 160238754, 158153231, '160238754');
INSERT INTO `relations` VALUES (161853510, 153969272, 158153231, '153969272');
INSERT INTO `relations` VALUES (162294051, 153969272, 160238754, '153969272');
COMMIT;

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `password` varchar(20) NOT NULL,
  `follow_count` int DEFAULT '0',
  `follower_count` int DEFAULT '0',
  `is_follow` tinyint(1) DEFAULT '0',
  `token` varchar(255) DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `users_name_uindex` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=163515682 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of users
-- ----------------------------
BEGIN;
INSERT INTO `users` VALUES (153969272, 'china', '123456', 0, 1, 1, '153969272');
INSERT INTO `users` VALUES (154347356, 'zxy', '123456', 1, 0, 1, '154347356');
INSERT INTO `users` VALUES (158153231, 'hjh', '123456', 0, 0, 0, '158153231');
INSERT INTO `users` VALUES (160238754, 'bytedance', '123456', 0, 1, 1, '160238754');
INSERT INTO `users` VALUES (162649382, 'ych', '123456', 0, 0, 0, '162649382');
INSERT INTO `users` VALUES (163493274, 'pony', '123456', 1, 0, 1, '163493274');
COMMIT;

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos` (
  `id` int NOT NULL AUTO_INCREMENT,
  `author_id` int NOT NULL,
  `play_url` varchar(255) NOT NULL DEFAULT '0',
  `cover_url` varchar(255) NOT NULL DEFAULT '0',
  `favorite_count` int NOT NULL DEFAULT '0',
  `comment_count` int NOT NULL DEFAULT '0',
  `is_favorite` tinyint(1) NOT NULL DEFAULT '0',
  `title` varchar(255) NOT NULL DEFAULT '0',
  `create_time` datetime NOT NULL,
  `token` varchar(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=162811647 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of videos
-- ----------------------------
BEGIN;
INSERT INTO `videos` VALUES (156083940, 153969272, 'hello.mp4', 'hello.jpg', 1, 0, 1, 'hello', '2022-06-12 21:38:59', '');
INSERT INTO `videos` VALUES (156719599, 163493274, 'gu.mp4', 'gu.jpg', 0, 0, 0, 'gu', '2022-06-12 21:45:55', '');
INSERT INTO `videos` VALUES (157400039, 160238754, '遵义.mp4', '遵义.jpg', 0, 0, 0, '遵义', '2022-06-12 21:34:21', '');
INSERT INTO `videos` VALUES (158504986, 154347356, '你的名字.mp4', '你的名字.jpg', 0, 0, 0, '你的名字', '2022-06-12 21:22:54', '');
INSERT INTO `videos` VALUES (159145641, 158153231, 'sheep.mp4', 'sheep.jpg', 0, 0, 0, 'sheep', '2022-06-12 21:30:40', '');
INSERT INTO `videos` VALUES (159321251, 158153231, 'swim.mp4', 'swim.jpg', 0, 0, 0, 'swim', '2022-06-12 21:26:29', '');
INSERT INTO `videos` VALUES (159636478, 162649382, '柴犬.mp4', '柴犬.jpg', 0, 0, 0, '柴犬', '2022-06-12 21:23:54', '');
INSERT INTO `videos` VALUES (160623400, 162649382, 'dogs.mp4', 'dogs.jpg', 0, 0, 0, 'dogs', '2022-06-12 21:24:15', '');
INSERT INTO `videos` VALUES (162723602, 154347356, 'cat.mp4', 'cat.jpg', 0, 0, 0, 'cat', '2022-06-12 21:23:22', '');
INSERT INTO `videos` VALUES (162811646, 160238754, 'messi.mp4', 'messi.jpg', 0, 0, 0, 'messi', '2022-06-12 21:34:09', '');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
