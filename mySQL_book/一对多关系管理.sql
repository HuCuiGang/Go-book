-- 修改学生表信息，为每个学生指定classid
update student set classid=(select id from clazz where clazz.name = '丐帮') where student.name='野间忠一郎';
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='二郎神';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='张三丰';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='郭小四';
update student set classid=(select id from clazz where clazz.name = '丐帮') where student.name='隔壁老王';
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='练过的六爷';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='洪七公';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='香香八婆';
update student set classid=(select id from clazz where clazz.name = '丐帮') where student.name='马英九';
update student set classid=(select id from clazz where clazz.name = '小刀会') where student.name='十三姨';
update student set classid=(select id from clazz where clazz.name = '斧头帮') where student.name='山本五十六';
update student set classid=(select id from clazz where clazz.name = '天地会') where student.name='包租婆';

-- 通过查询学生表中所有classid指向“小刀会”的记录，来查询班级“小刀会”的所有学生
select * from student where classid = (
    select id from clazz where name = '丐帮'
    );

# 根据学生查询其班级
select * from clazz where id = (
    select classid from student where name = '十三姨'
    );




