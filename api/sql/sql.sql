CREATE DATABASE IF NOT EXISTS golangsocialnetwork;
USE golangsocialnetwork;

DROP TABLE IF EXISTS users;

CREATE TABLE users (
    id int auto_increment primary key,
    name varchar(50) not null,
    nick varchar(50) not null unique,
    mail varchar(50) not null unique,
    password varchar(100) not null,
    createin timestamp default current_timestamp()
) ENGINE=INNODB;