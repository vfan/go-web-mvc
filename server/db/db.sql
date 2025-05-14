-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS student_management DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 使用数据库
USE student_management;

-- 用户表（用于鉴权）
CREATE TABLE IF NOT EXISTS `users` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `email` VARCHAR(100) NOT NULL COMMENT '邮箱(登录账号)',
  `password` VARCHAR(255) NOT NULL COMMENT '密码(加密存储)',
  `role` TINYINT NOT NULL DEFAULT 2 COMMENT '角色(1:管理员,2:普通用户)',
  `last_login_time` DATETIME DEFAULT NULL COMMENT '上次登录时间',
  `status` TINYINT NOT NULL DEFAULT 1 COMMENT '状态(0:禁用,1:启用)',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`),
  KEY `idx_role` (`role`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='用户表';

-- 大学表
CREATE TABLE IF NOT EXISTS `universities` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '大学ID',
  `name` VARCHAR(100) NOT NULL COMMENT '大学名称',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  `created_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '创建人ID',
  `updated_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '最后修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='大学信息表';

-- 学生表
CREATE TABLE IF NOT EXISTS `students` (
  `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '学生ID（主键）',
  `name` VARCHAR(50) NOT NULL COMMENT '学生姓名',
  `password` VARCHAR(255) NOT NULL COMMENT '登录密码(加密存储)',
  `email` VARCHAR(100) NOT NULL COMMENT '电子邮箱(登录账号，唯一标识)',
  `last_login_time` DATETIME DEFAULT NULL COMMENT '上次登录时间',
  `gender` TINYINT DEFAULT 1 COMMENT '性别(1:男,2:女,3:其他)',
  `birthday` DATE DEFAULT NULL COMMENT '出生日期',
  `phone` VARCHAR(20) DEFAULT NULL COMMENT '联系电话',
  `resume_path` VARCHAR(255) DEFAULT NULL COMMENT '简历文件相对路径',
  `university_id` BIGINT UNSIGNED DEFAULT NULL COMMENT '大学ID（关联universities表）',
  `major` VARCHAR(100) DEFAULT NULL COMMENT '专业名称',
  `education` ENUM('专科', '本科', '硕士', '博士') DEFAULT '本科' COMMENT '学历',
  `graduation_year` YEAR DEFAULT NULL COMMENT '毕业年份',
  `status` ENUM('在读', '休学', '退学', '毕业') DEFAULT '在读' COMMENT '学生状态',
  `remarks` TEXT DEFAULT NULL COMMENT '备注信息',
  `avatar` VARCHAR(255) DEFAULT NULL COMMENT '头像URL',
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` DATETIME DEFAULT NULL COMMENT '删除时间',
  `created_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '创建人ID',
  `updated_by` BIGINT UNSIGNED DEFAULT NULL COMMENT '最后修改人ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_email` (`email`),
  KEY `idx_name` (`name`),
  KEY `idx_university_id` (`university_id`),
  KEY `idx_graduation_year` (`graduation_year`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='学生信息表';
