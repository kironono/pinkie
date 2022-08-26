
-- +migrate Up
CREATE TABLE `users` (
  `id`         bigint unsigned not null auto_increment,
  `email`      varchar(128) not null,
  `password`   varchar(128) not null,
  `created_at` datetime not null,
  `updated_at` datetime not null,
  PRIMARY KEY(`id`),
  UNIQUE KEY `uniq_email` (`email`) USING BTREE
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `users`;
