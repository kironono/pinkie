
-- +migrate Up
CREATE TABLE `job_sessions` (
  `id`          bigint unsigned not null auto_increment,
  `job_id`      bigint unsigned not null,
  `start_at`    datetime not null,
  `end_at`      datetime,
  `created_at`  datetime not null,
  `updated_at`  datetime not null,
  PRIMARY KEY(`id`)
) Engine=InnoDB DEFAULT CHARSET=utf8mb4;

-- +migrate Down
DROP TABLE `job_sessions`;
