CREATE TABLE IF NOT EXISTS `recipes` (
    id integer PRIMARY KEY AUTO_INCREMENT,
    title varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    making_time varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    serves varchar(100) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    ingredients varchar(300) CHARACTER SET utf8 COLLATE utf8_unicode_ci NOT NULL,
    cost integer NOT NULL,
    created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime on update CURRENT_TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

