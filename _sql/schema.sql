CREATE TABLE IF NOT EXISTS `rooms` (
  `room_id` int PRIMARY KEY AUTO_INCREMENT,
  `host` varchar(100) NOT NULL,
  `room_code` varchar(15) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS `questions` (
  `question_id` int PRIMARY KEY AUTO_INCREMENT,
  `question` varchar(2000) NOT NULL,
  `questioner_name` varchar(45) NOT NULL DEFAULT 'Anonymous',
  `total_likes` int NOT NULL DEFAULT 0,
  `fk_room_code` varchar(15) NOT NULL
);

ALTER TABLE `questions` ADD FOREIGN KEY (`fk_room_code`) REFERENCES `rooms` (`room_code`) ON DELETE CASCADE;