-- +migrate Up
DROP TABLE IF EXISTS `m_event`;
CREATE TABLE `m_event` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '裏で管理する用のId',
  `event_id` varchar(10) NOT NULL COMMENT 'api_id-event_id',
  `api_id` int(1) NOT NULL COMMENT 'API ID - eventID\nATDN 0\nCONNPASS 1\nDOORKIEEPER 2',
  `title` varchar(200) NOT NULL,
  `description` text NOT NULL,
  `url` text NOT NULL,
  `limit_count` int(11) NOT NULL,
  `waitlisted` int(11) NOT NULL,
  `accepted` int(11) NOT NULL COMMENT 'connpassにはない',
  `address` text NOT NULL,
  `place` text NOT NULL,
  `start_at` datetime NOT NULL,
  `end_at` datetime NOT NULL,
  `data_hash` char(64) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
  PRIMARY KEY (`id`),
  UNIQUE (`event_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `m_event`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `event_id` (`event_id`),
  ADD KEY `search_index` (`end_at`,`updated_at`,`address`(3));


-- +migrate Down
DROP TABLE m_event;
