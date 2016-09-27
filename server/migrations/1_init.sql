-- +migrate Up
DROP TABLE IF EXISTS `m_event`;
CREATE TABLE `m_event` (
		  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '裏で管理する用のId',
		  `eventId` varchar(10) NOT NULL COMMENT 'API ID - eventID\nATDN 0\nCONNPASS 1\nDOORKIEEPER 2',
		  `title` varchar(30) NOT NULL,
		  `description` text NOT NULL,
		  `url` text NOT NULL,
		  `limitCount` int(11) NOT NULL,
		  `waitlisted` int(11) NOT NULL,
		  `accepted` int(11) NOT NULL COMMENT 'connpassにはない',
		  `address` text NOT NULL,
		  `place` text NOT NULL,
		  `startAt` datetime NOT NULL,
		  `endAt` datetime NOT NULL,
		  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
		  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
		  PRIMARY KEY (`id`),
		  UNIQUE KEY `eventId` (`eventId`)
) ENGINE=InnoDB AUTO_INCREMENT=1227 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE m_event;
