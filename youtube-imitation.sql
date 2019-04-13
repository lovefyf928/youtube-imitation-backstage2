drop database IF EXISTS youtubeImitation;
create database youtubeImitation;
use youtubeImitation;
create table User(
uid int primary key auto_increment,
userName varchar(20) not null,
email varchar(30) not null,
phoneNumber varchar(20) not null,
password varchar(50) not null,
sex tinyint(1),
birthday date,
code varchar(10)
);
create table OnePointLogin(
uid int not null,
ip varchar(20) not null,
loginTime date not null
);
create table Favorite(
favoriteId int not null key auto_increment,
name varchar(20) not null,
description varchar(256) not null
);
create table FavoriteList(
uid int not null,
favoriteName varchar(20) not null
);
