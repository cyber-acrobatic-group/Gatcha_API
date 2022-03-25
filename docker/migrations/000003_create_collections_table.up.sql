CREATE TABLE `collections`
(
    `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'コレクションID',
    `user_id` bigint UNSIGNED NOT NULL,
    `photo_id` bigint UNSIGNED NOT NULL,
    `appear_times` bigint UNSIGNED NOT NULL COMMENT '出現回数',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
    PRIMARY KEY(`id`),
    CONSTRAINT `fk_collections_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT `fk_collections_photo` FOREIGN KEY (`photo_id`) REFERENCES `photos` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB
  DEFAULT CHARSET = `utf8mb4`
  DEFAULT COLLATE = `utf8mb4_bin` COMMENT 'コレクションテーブル';