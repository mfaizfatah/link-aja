/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80023
 Source Host           : localhost:3308
 Source Schema         : link-aja

 Target Server Type    : MySQL
 Target Server Version : 80023
 File Encoding         : 65001

 Date: 04/04/2021 14:05:25
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for customer
-- ----------------------------
DROP TABLE IF EXISTS `customer`;
CREATE TABLE `customer`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `customer_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted` tinyint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `customer_number`(`customer_number`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of customer
-- ----------------------------
INSERT INTO `customer` VALUES (1, '1001', 'Bob Martin', '2021-04-04 06:45:27', '2021-04-04 06:45:27', 0);
INSERT INTO `customer` VALUES (2, '1002', 'Linus Torvalds', '2021-04-04 06:45:38', '2021-04-04 06:45:38', 0);

-- ----------------------------
-- Table structure for account
-- ----------------------------
DROP TABLE IF EXISTS `account`;
CREATE TABLE `account`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `account_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `customer_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `balance` float NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted` tinyint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `customer_number`(`customer_number`) USING BTREE,
  CONSTRAINT `account_ibfk_1` FOREIGN KEY (`customer_number`) REFERENCES `customer` (`customer_number`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of account
-- ----------------------------
INSERT INTO `account` VALUES (1, '555001', '1001', 10000, '2021-04-04 06:45:55', '2021-04-04 06:45:55', 0);
INSERT INTO `account` VALUES (2, '55022', '1002', 15000, '2021-04-04 06:46:04', '2021-04-04 06:46:04', 0);

-- ----------------------------
-- Table structure for transfer
-- ----------------------------
DROP TABLE IF EXISTS `transfer`;
CREATE TABLE `transfer`  (
  `id` int UNSIGNED NOT NULL AUTO_INCREMENT,
  `from_account_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `to_account_number` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `amount` float NULL DEFAULT NULL,
  `created_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `updated_at` datetime(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  `deleted` tinyint NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of transfer
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
