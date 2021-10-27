/*
 * 创建数据库
 */
CREATE DATABASE IF NOT EXISTS `file_index` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

-- 目录表
CREATE TABLE `t_folder`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `parent_id` bigint NOT NULL DEFAULT 0 COMMENT '父目录ID',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '目录名称',
    `path` varchar(1000) NOT NULL DEFAULT '' COMMENT '目录绝对路径',
    `type` varchar(128) NOT NULL DEFAULT '' COMMENT '存储类型',
    `removed` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否被删除，0-未删除，1-已删除',
    `hideen` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏，0-不隐藏，1-隐藏',
    `desc` varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_parent_id` (`parent_id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '目录表';

-- 文件表
CREATE TABLE `t_file`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `folder_id` bigint NOT NULL DEFAULT 0 COMMENT '所在目录ID',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名称',
    `path` varchar(255) NOT NULL DEFAULT '' COMMENT '文件绝对路径名称',
    `size` bigint DEFAULT 0 COMMENT '文件大小，单位：B',
    `size_k` bigint DEFAULT 0 COMMENT '文件大小，单位：KB',
    `size_m` int(8) DEFAULT 0 COMMENT '文件大小，单位：MB',
    `size_g` int(8) DEFAULT 0 COMMENT '文件大小，单位：GB',
    `suffix` varchar(32) DEFAULT '' COMMENT '文件后缀名',
    `status` tinyint(2) NOT NULL DEFAULT 0 COMMENT '状态，0-插入，1-更新，2-损坏丢失，3-无法打开',
    `removed` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否被删除，0-未删除，1-已删除',
    `hideen` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏，0-不隐藏，1-隐藏',
    `file_time` datetime NOT NULL COMMENT '文件创建时间',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY `idx_parent_id` (`parent_id`) USING BTREE,
    KEY `idx_name` (`name`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '文件表';

-- minio配置表
CREATE TABLE `t_minio_config`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '配置名称',
    `value` varchar(255) NOT NULL DEFAULT '' COMMENT '配置值',
    `desc` varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = 'minio配置表';

-- 统计分析表
CREATE TABLE `t_statistics`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `type` varchar(255) NOT NULL DEFAULT '' COMMENT '文件、目录、标签类型',
    `count` int(8) NOT NULL DEFAULT 0 COMMENT '统计总数',
    `desc` varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '统计分析表';

-- 标签表
CREATE TABLE `t_tag`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '标签名称',
    `desc` varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '标签表';

-- 文件标签关系表
CREATE TABLE `t_tag_file`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `tag_id` bigint NOT NULL COMMENT '标签ID',
    `file_id` bigint NOT NULL COMMENT '文件ID',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '文件标签关系表';

-- 目录标签关系表
CREATE TABLE `t_tag_folder`(
    `id` bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `tag_id` bigint NOT NULL COMMENT '标签ID',
    `folder_id` bigint NOT NULL COMMENT '目录ID',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '目录标签关系表';