-- +goose Up
CREATE TABLE `todo` (
    id int(11) UNSIGNED NOT NULL auto_increment,
    title VARCHAR(256) NOT NULL,
    content VARCHAR(512) NOT NULL DEFAULT '',
    created_at datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime(0) NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci;

-- +goose Down
DROP TABLE IF EXISTS `todo`;
