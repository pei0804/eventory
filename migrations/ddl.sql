-- phpMyAdmin SQL Dump
-- version 4.5.2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: 2017 年 2 月 15 日 14:51
-- サーバのバージョン： 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+09:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `eventory`
--

-- --------------------------------------------------------

--
-- テーブルの構造 `m_event`
--

CREATE TABLE `m_event` (
  `id` int(11) NOT NULL COMMENT 'Id',
  `event_id` varchar(10) NOT NULL COMMENT 'api_id-event_id',
  `api_id` int(1) NOT NULL COMMENT 'API ID - eventID\nATDN 0\nCONNPASS 1\nDOORKIEEPER 2',
  `title` varchar(200) NOT NULL,
  `description` text NOT NULL,
  `url` text NOT NULL,
  `limit_count` int(11) NOT NULL,
  `waitlisted` int(11) NOT NULL,
  `accepted` int(11) NOT NULL COMMENT 'connpass only',
  `address` text NOT NULL,
  `place` text NOT NULL,
  `start_at` datetime NOT NULL,
  `end_at` datetime NOT NULL,
  `data_hash` char(64) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- テーブルのデータのダンプ `m_event`
--

--
-- Indexes for table `m_event`
--
ALTER TABLE `m_event`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `event_id` (`event_id`),
  ADD KEY `search_index` (`end_at`,`updated_at`,`address`(3));

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `m_event`
--
ALTER TABLE `m_event`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'è£ã§ç®¡ç†ã™ã‚‹ç”¨ã®Id', AUTO_INCREMENT=665;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
