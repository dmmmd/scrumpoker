USE `scrumpoker`;

CREATE TABLE IF NOT EXISTS `grooming_sessions`
(
  `id`    char(36)     NOT NULL DEFAULT '',
  `title` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
