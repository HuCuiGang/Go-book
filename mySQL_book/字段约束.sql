# distinct 不重复的内容
select distinct ProID from T_City;

# 字段约束
create table person(
    id int primary key auto_increment, -- 主键约束，自增
    name varchar(20) unique not null , -- 唯一约束，非空约束
    age int default 0,-- 默认值约束
    rmb float default 0,
    sex bool,
    signature varchar(100) default '这个家伙很懒，没有留下任何个性签名',
    -- 生日，日期类型
    birthday date,
    scbz datetime
);
drop table person;

insert into person(name,sex,birthday,scbz) values ('毛泽东',true,18931226,18931226123456);
insert into person(name,sex) values ('武则天',false);


select * from person;


