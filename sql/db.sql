
create database video default charset utf8;
use video;


create table users(author_id bigint not null primary key auto_increment,
login_name varchar(64) ,pwd text) engine=innodb;


create table video_info(video_id varchar(64) not null primary key,author_id bigint,name text,
display_ctime text,create_time datetime) engine=innodb;



create table sessions(session_id tinytext  not null ,
ttl tinytext,login_name varchar(64)) engine=innodb;



create table  comments(comment_id varchar(64) not null primary key,video_id varchar(64),

author_id  bigint,content text,time datetime
) engine=innodb;



