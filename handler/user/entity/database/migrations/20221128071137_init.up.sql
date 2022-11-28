CREATE TABLE `mst_user` (
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(100) NOT NULL,
  `phone` VARCHAR(14) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `status` INT(2) NOT NULL,
  `password` BLOB NOT NULL,
  `user_type_id` INT(2) NOT NULL,
  `created_by` BIGINT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_by` BIGINT NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `mst_user_status` (
  `id` BIGINT NOT NULL,
  `key` VARCHAR(50) NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `mst_user_type` (
  `id` BIGINT NOT NULL,
  `name` VARCHAR(50) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO `mst_user_status` VALUES (0, "non-active", "Non Active", now(), now()), (1, "active", "Active", now(), now());

INSERT INTO `mst_user_type` VALUES (1, "Admin", now(), now()), (2, "Regular", now(), now());

INSERT INTO `mst_user` VALUES (1, "Rivaldy", "087751231234", "admin@example.com", 1, "$2a$04$w8SjLzG5jcPG4r8b8NKayeqlbTErP8iTpnvd6kQjR5C/SBlklR/wW", 1, 1, now(), 1, now());