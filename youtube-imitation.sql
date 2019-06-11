drop database IF EXISTS youtubeImitation;
SET NAMES 'utf8';
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
) default charset=utf8;
create table OnePointLogin(
uid int not null,
ip varchar(20) not null,
loginTime date not null
) default charset=utf8;
create table Favorite(
favoriteId int not null key auto_increment,
name varchar(20) not null,
description varchar(256) not null
) default charset=utf8;
create table FavoriteList(
uid int not null,
favoriteName varchar(20) not null
) default charset=utf8;
create table Channel(
id int not null primary key auto_increment,
uid int not null,
name varchar(20) not null,
subscriber int,
classification varchar(20),
introduction varchar(30),
registrationTime date not null,
channelDescription varchar(500),
position varchar(20),
link varchar(100),
foreign key (uid) references user(uid)
) default charset=utf8;
create table Video(
id int not null primary key auto_increment,
channelId int not null,
channelClassification varchar(20),
category varchar(20) not null,
videoIntroduction varchar(500),
name varchar(50) not null,
viewCount int not null,
good int not null,
bad int not null,
releaseTime varchar(100) not null,
videoPath varchar(100) not null,
videoImg varchar(100) not null,
foreign key (channelId) references Channel(id)
) default charset=utf8;
create table Comment(
videoId int not null,
userId int not null,
content varchar(500) not null,
commentTime date,
good int,
bad int
) default charset=utf8;
create table subscribe(
userId int not null,
Subscribed varchar(20) not null,
foreign key (userId) references User(uid),
foreign key (Subscribed) references Channel(id)
) default charset=utf8;
