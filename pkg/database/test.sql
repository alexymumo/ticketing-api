CREATE TABLE `user`(
    `userid` int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `fullname` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL
) ENGINE=InnoDB
DEFAULT CHARSET=utf8;

drop table event if exists;
create table `event` (
    `eventid` int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `title` varchar(255) NOT NULL,
    `imageUrl` varchar(255) NOT NULL,
    `date` varchar(255) NOT NULL,
    `venue` varchar(255) NOT NULL,
    `description` varchar(255) NOT NULL,
    `time` varchar(255),
    `amount` varchar(255),
    `capacity` int
) ENGINE=InnoDB
DEFAULT CHARSET=utf8

