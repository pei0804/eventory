SET SESSION FOREIGN_KEY_CHECKS=0;

/* Drop Tables */

DROP TABLE IF EXISTS event_genres;
DROP TABLE IF EXISTS user_keep_statuses;
DROP TABLE IF EXISTS events;
DROP TABLE IF EXISTS user_follow_genres;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS user_follow_prefs;
DROP TABLE IF EXISTS prefs;
DROP TABLE IF EXISTS user_terminals;
DROP TABLE IF EXISTS users;

/* Create Tables */

-- イベント
CREATE TABLE events
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'イベントID',
	api_type enum('atdn','connpass','doorkeeper') NOT NULL COMMENT 'APIの種類',
	identifier varchar(10) NOT NULL COMMENT '識別子(api-event_id)',
	title varchar(200) NOT NULL COMMENT 'イベント名',
	description text NOT NULL COMMENT '説明',
	url text NOT NULL COMMENT 'イベントページURL',
	limits int NOT NULL COMMENT '参加人数上限',
	wait int NOT NULL COMMENT 'キャンセル待ち人数',
	accept int NOT NULL COMMENT '参加済み人数',
	pref_id int(2) unsigned COMMENT '都道府県ID',
	address text NOT NULL COMMENT '住所',
	start_at datetime NOT NULL COMMENT '開催日時',
	end_at datetime NOT NULL COMMENT '終了日時',
	data_hash char(64) NOT NULL COMMENT 'データ識別Hash',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id),
	UNIQUE (identifier)
) COMMENT = 'イベント' DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;


-- イベントジャンル
CREATE TABLE event_genres
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	genre_id bigint(20) unsigned NOT NULL COMMENT 'ジャンルID',
	event_id bigint(20) unsigned NOT NULL COMMENT 'イベントID',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'イベントジャンル';


-- ジャンル
CREATE TABLE genres
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ジャンルID',
	name varchar(30) NOT NULL COMMENT 'ジャンル名(表示用)',
	keyword varchar(50) NOT NULL COMMENT 'キーワード',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ジャンル';


-- 都道府県
CREATE TABLE prefs
(
	id int(2) unsigned NOT NULL AUTO_INCREMENT COMMENT '都道府県ID',
	name char(4) NOT NULL COMMENT '都道府県名',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = '都道府県';


-- ユーザー
CREATE TABLE users
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
	name varchar(30) NOT NULL COMMENT 'ユーザー名',
	email text NOT NULL COMMENT 'メールアドレス',
	password_hash char(64) NOT NULL COMMENT 'パスワードハッシュ',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザー';


-- ユーザーのキープ状態
CREATE TABLE user_keep_statuses
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	event_id bigint(20) unsigned NOT NULL COMMENT 'イベントID',
	status enum('keep','nokeep') NOT NULL COMMENT '状態',
	batch_processed tinyint(0) DEFAULT 0 NOT NULL COMMENT 'バッチ処理済み',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザーのキープ状態';


-- ユーザーのフォロージャンル
CREATE TABLE user_follow_genres
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ジャンルID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	genre_id bigint(20) unsigned NOT NULL COMMENT 'ジャンルID',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザーのフォロージャンル';


-- ユーザーのフォロー都道府県
CREATE TABLE user_follow_prefs
(
	id int(2) unsigned NOT NULL AUTO_INCREMENT COMMENT '都道府県ID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	pref_id int(2) unsigned NOT NULL COMMENT '都道府県ID',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザーのフォロー都道府県';


-- ユーザー端末情報
CREATE TABLE user_terminals
(
	id bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
	user_id bigint(20) unsigned NOT NULL COMMENT 'ユーザーID',
	platform varchar(20) NOT NULL COMMENT 'OSとバージョン',
	client_version varchar(10) NOT NULL COMMENT 'アプリのバージョン',
	token char(64) NOT NULL COMMENT 'トークン',
	identifier char(36) NOT NULL COMMENT '識別子(android:Android_ID, ios:IDFV)',
	created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日',
	updated_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '更新日',
	deleted_at datetime COMMENT '削除日',
	PRIMARY KEY (id)
) COMMENT = 'ユーザー端末情報';



/* Create Foreign Keys */

ALTER TABLE event_genres
	ADD FOREIGN KEY (event_id)
	REFERENCES events (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_keep_statuses
	ADD FOREIGN KEY (event_id)
	REFERENCES events (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE event_genres
	ADD FOREIGN KEY (genre_id)
	REFERENCES genres (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_genres
	ADD FOREIGN KEY (genre_id)
	REFERENCES genres (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE events
	ADD FOREIGN KEY (pref_id)
	REFERENCES prefs (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_prefs
	ADD FOREIGN KEY (pref_id)
	REFERENCES prefs (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_keep_statuses
	ADD FOREIGN KEY (user_id)
	REFERENCES users (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_genres
	ADD FOREIGN KEY (user_id)
	REFERENCES users (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_follow_prefs
	ADD FOREIGN KEY (user_id)
	REFERENCES users (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;


ALTER TABLE user_terminals
	ADD FOREIGN KEY (user_id)
	REFERENCES users (id)
	ON UPDATE RESTRICT
	ON DELETE RESTRICT
;
