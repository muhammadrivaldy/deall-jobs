CREATE TABLE IF NOT EXISTS `mst_service` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`)
);

CREATE TABLE IF NOT EXISTS `mst_api` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) CHARACTER SET latin1 NOT NULL,
  `endpoint` varchar(255) CHARACTER SET latin1 NOT NULL,
  `method` varchar(15) CHARACTER SET latin1 NOT NULL,
  `service_id` int(11) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name_UNIQUE` (`name`),
  UNIQUE KEY `endpoint_method_UNIQUE` (`endpoint`,`method`)
);

CREATE TABLE IF NOT EXISTS `mst_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_type_id` int(2) NOT NULL,
  `api_id` int(3) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
);