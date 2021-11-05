-- 创建数据库
create database komablog;
-- 列出数据库名
select  datname FROM pg_database;
-- 查询所有数据信息
select * from pg_database;
-- 删除数据库
drop database komablog;

-- 创建数据表
create table person
(
    id   int primary key ,  -- 自增auto_
    name varchar,
    age int,
    sex bool,
    rmb float,
    word text
);
-- 查询表结构 desc person;

-- 插入数据
insert into person(id, name, age, sex, rmb ,word) values (1,'张三丰',90,true,12344123124,'我是大学教授');
-- 插入多行数据
insert into person(id, name, age, sex, rmb ) values
                            (3,'张无忌',25,true,1000),
                            (4,'周芷若',22,false,500);

-- 查询表中所有字段信息
select * from komablog.public."person" ;

-- 查询部分信息
select name,rmb from person;

-- 删除表
drop table person ;

-- 删除某条记录
delete from person where name ='张无忌';

-- 修改表数据
update person set name ='张无忌' where name= '张三丰';





