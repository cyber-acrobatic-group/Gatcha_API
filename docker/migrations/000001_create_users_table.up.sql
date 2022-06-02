CREATE TABLE `users`
(
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
    `name` varchar(100) NOT NULL COMMENT 'ユーザー名',
    `password` varchar(100) NOT NULL COMMENT 'パスワード',
    `count` bigint UNSIGNED NOT NULL COMMENT 'ガチャ回数',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  DEFAULT   COLLATE = utf8mb4_unicode_ci COMMENT 'ユーザーテーブル';