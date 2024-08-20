CREATE TABLE `users`(
    `userid` int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `fullname` varchar(255) NOT NULL,
    `email` varchar(255) NOT NULL UNIQUE,
    `password` varchar(255) NOT NULL,
    `role` ENUM('Seller','User') DEFAULT 'User'  
) ENGINE=InnoDB
DEFAULT CHARSET=utf8;

create table `location`(
    `locationid` int AUTO_INCREMENT PRIMARY KEY NOT NULL,
    `latitude` float 
    `longitude` float
)