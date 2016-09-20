SET SESSION FOREIGN_KEY_CHECKS=0;

/* Drop Tables */

DROP TABLE IF EXISTS m_event;




/* Create Tables */

CREATE TABLE m_event
(
	-- 裏で管理する用のId
	id int NOT NULL AUTO_INCREMENT COMMENT '裏で管理する用のId',
	-- API ID - eventID
	-- ATDN 0
	-- CONNPASS 1
	-- DOORKIEEPER 2
	eventId varchar(10) NOT NULL COMMENT 'API ID - eventID
ATDN 0
CONNPASS 1
DOORKIEEPER 2',
	title varchar(30) NOT NULL,
	description text NOT NULL,
	url text NOT NULL,
	limitCount int NOT NULL,
	waitlisted int NOT NULL,
	-- connpassにはない
	accepted int NOT NULL COMMENT 'connpassにはない',
	address text NOT NULL,
	place text NOT NULL,
	startAt datetime NOT NULL,
	endAt datetime NOT NULL,
	created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
	PRIMARY KEY (id),
	UNIQUE (eventId)
);
