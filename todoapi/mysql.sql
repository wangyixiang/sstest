CREATE DATABASE `tododb` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `todoapi`;
CREATE TABLE `todo_webapi_todoitem` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `content` varchar(255) CHARACTER SET latin1 NOT NULL DEFAULT '',
  `created_time` datetime NOT NULL,
  `finished_time` datetime NOT NULL,
  `finished` tinyint(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
