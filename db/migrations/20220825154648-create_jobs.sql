
-- +migrate Up
CREATE TABLE `jobs` (
  `id` bigint unsigned not null auto_increment,
  `name` varchar(64) not null,
  `slug` varchar(64) not null,
  `created_at` datetime not null,
  `updated_at` datetime not null,
  PRIMARY KEY(`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `jobs`;
