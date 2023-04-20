/*
 * 创建数据库
 */
CREATE
DATABASE IF NOT EXISTS `file_index` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_unicode_ci;

-- 文件表
CREATE TABLE `t_file`
(
    `id`          bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`        varchar(255) NOT NULL DEFAULT '' COMMENT '文件名称',
    `path`        varchar(512) NOT NULL DEFAULT '' COMMENT '文件绝对路径名称',
    `size`        bigint(20) NOT NULL DEFAULT 0 COMMENT '文件大小',
    `suffix`      varchar(32)  NOT NULL DEFAULT '' COMMENT '文件后缀名',
    `status`      tinyint(2) NOT NULL DEFAULT 0 COMMENT '状态，0-插入，1-更新，2-损坏丢失，3-无法打开',
    `removed`     tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否被删除，0-未删除，1-已删除',
    `hidden`      tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否隐藏，0-不隐藏，1-隐藏',
    `create_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint(20) NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE,
    KEY           `idx_name` (`name`) USING BTREE,
    KEY           `idx_path` (`path`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '文件表';

-- minio配置表
CREATE TABLE `t_minio_config`
(
    `id`          bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`        varchar(255)  NOT NULL DEFAULT '' COMMENT '配置名称',
    `value`       varchar(255)  NOT NULL DEFAULT '' COMMENT '配置值',
    `desc`        varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint        NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint        NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = 'minio配置表';

-- 统计分析表
CREATE TABLE `t_statistics`
(
    `id`          bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `type`        varchar(255)  NOT NULL DEFAULT '' COMMENT '文件、目录、标签类型',
    `count`       int(8) NOT NULL DEFAULT 0 COMMENT '统计总数',
    `desc`        varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint        NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint        NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '统计分析表';

-- 标签表
CREATE TABLE `t_tag`
(
    `id`          bigint        NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `name`        varchar(255)  NOT NULL DEFAULT '' COMMENT '标签名称',
    `desc`        varchar(1000) NOT NULL DEFAULT '' COMMENT '说明',
    `create_time` bigint        NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint        NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '标签表';

-- 文件标签关系表
CREATE TABLE `t_tag_file`
(
    `id`          bigint NOT NULL AUTO_INCREMENT COMMENT 'ID',
    `tag_id`      bigint NOT NULL COMMENT '标签ID',
    `file_id`     bigint NOT NULL COMMENT '文件ID',
    `create_time` bigint NOT NULL DEFAULT 0 COMMENT '创建时间',
    `update_time` bigint NOT NULL DEFAULT 0 COMMENT '最后修改时间',
    PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci COMMENT = '文件标签关系表';
