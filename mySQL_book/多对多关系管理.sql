-- 创建【学生_课程】中间表
create table student_course(
  -- 学生id
  sid integer not null,

  -- 课程id
  cid integer not null,

  -- 主键是【学生id和课程id的联合主键】
  primary key (sid,cid)
);

-- 如果没有设置主键，可以通过修改表字段的方式来添加【联合主键】
-- alter table student_course add constraint s_c primary key (sid,cid);

select * from student_course;

-- 野间忠一郎选修了Python
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='野间忠一郎'),
  (select id from course where course.name='Python')
);

-- 野间忠一郎选修了Java
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='野间忠一郎'),
  (select id from course where course.name='Java')
);

-- 野间忠一郎选修了HTML5
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='野间忠一郎'),
  (select id from course where course.name='HTML5')
);

-- 香香八婆选修了Python
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='香香八婆'),
  (select id from course where course.name='Python')
);

-- 香香八婆选修了HTML5
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='香香八婆'),
  (select id from course where course.name='HTML5')
);

-- 二郎神选修了Python
insert into student_course(sid, cid) values (
  (SELECT id from student where student.name='二郎神'),
  (select id from course where course.name='Python')
);

# 能根据学生查询其所有的选修
select * from course where id in (
    select cid from student_course
    where sid = (
        select id from student where name = '野间忠一郎'
        )
);

# 根据课程查询其所有学员
select * from student where id in (
    select sid from student_course where cid =(
    select id from course where name= "Python"
    )
);

