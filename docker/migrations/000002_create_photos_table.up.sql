CREATE TABLE `photos`
(
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '写真ID',
    `text` varchar(255) NOT NULL COMMENT 'テキスト',
    `rarity` tinyint UNSIGNED NOT NULL COMMENT 'レアリティ',
    `image_path` text NOT NULL COMMENT '画像パス',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(`id`),
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  DEFAULT COLLATE = `utf8mb4_bin` COMMENT '写真テーブル';