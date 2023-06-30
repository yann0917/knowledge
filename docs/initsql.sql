/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 50616
 Source Host           : 127.0.0.1:4306
 Source Schema         : zsxq

 Target Server Type    : MySQL
 Target Server Version : 50616
 File Encoding         : 65001

 Date: 30/06/2023 22:16:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for column_topic
-- ----------------------------
DROP TABLE IF EXISTS `column_topic`;
CREATE TABLE `column_topic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `column_id` bigint(20) NOT NULL COMMENT '专栏 ID',
  `topic_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '话题 ID',
  `attached_to_column_time` timestamp NULL DEFAULT NULL COMMENT '话题加入时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `column_id` (`column_id`) USING BTREE,
  KEY `topic_id` (`topic_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='专栏话题关系表';


-- ----------------------------
-- Table structure for group_column
-- ----------------------------
DROP TABLE IF EXISTS `group_column`;
CREATE TABLE `group_column` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '群组 ID',
  `column_id` bigint(20) NOT NULL COMMENT '专栏 ID',
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '专栏标题',
  `cover_url` varchar(255) NOT NULL COMMENT 'cover',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '专栏创建时间',
  `last_topic_attach_time` timestamp NULL DEFAULT NULL COMMENT '帖子最近加入时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE,
  KEY `column_id` (`column_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='知识星球群组专栏分类';


-- ----------------------------
-- Table structure for group
-- ----------------------------
DROP TABLE IF EXISTS `group`;
CREATE TABLE `group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '群组 ID',
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '群组标题',
  `description` text NOT NULL COMMENT '群组简介',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '群组创建时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='知识星球群组';


-- ----------------------------
-- Table structure for topic
-- ----------------------------
DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `group_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '群组 ID',
  `topic_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT 'topic ID',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `content` text NOT NULL COMMENT '内容',
  `rich_content` text NOT NULL COMMENT '富文本内容',
  `images` text NOT NULL COMMENT '图片url, 多个逗号分割',
  `is_digests` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否精华帖',
  `article_id` varchar(255) NOT NULL COMMENT '文章 ID',
  `article_url` varchar(255) NOT NULL COMMENT '文章链接',
  `create_time` timestamp NULL DEFAULT NULL COMMENT '主题创建时间',
  `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE,
  KEY `topic_id` (`topic_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='知识星球主题';

SET FOREIGN_KEY_CHECKS = 1;
