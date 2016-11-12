
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE IF NOT EXISTS `word` (
    `id`         BIGINT(20) UNSIGNED NOT NULL PRIMARY KEY,
    `name`       VARCHAR(40) NOT NULL,
    `word_type`  TINYINT(3), 
    `meaning`    TEXT NOT NULL,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL
) ENGINE=InnodB DEFAULT CHARSET=utf8mb4; 

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS `word`;

