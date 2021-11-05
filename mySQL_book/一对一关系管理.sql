# 给班级分配班主任
update clazz set masterid = (
    -- 查询robin的id
    select id from teacher where name = 'robin'
    ) where name = '丐帮';

update clazz set masterid = (
    -- 查询robin的id
    select id from teacher where name = 'steve'
    ) where name = '斧头帮';

update clazz set masterid = (
    -- 查询robin的id
    select id from teacher where name = 'bill'
    ) where name = '天地会';

update clazz set masterid = (
    -- 查询robin的id
    select id from teacher where name = 'jackma'
    ) where name = '小刀会';

select * from clazz;
select * from teacher;

# 根据班级找班主任
select * from teacher where id = (
    select masterid from clazz where name ='天地会'
    );

# 根据班主任找班级
select * from clazz where masterid=(
    select id from teacher where name ='jackma'
    );






