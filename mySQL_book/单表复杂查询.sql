# 找出不叫区县的区县
select * from T_District
where (DisName not like '%区') and (DisName not like '%县');
# 找出内蒙和新疆的省份ID
select * from T_Province
where ProName like '%新疆%' or ProName like '%内蒙%';
# 修改新疆维吾尔自治区  的名字
update T_Province set ProName ='新疆维吾尔自治区'
where ProName = '新 疆维吾尔自治区';
# 找出内蒙和新疆的所有地级市
select * from T_City
where (ProID=5) or (ProID=24);
# 查询河北所有的地区市，按CityID降序
select * from T_City where ProID=3
order by CityID desc;
# 升序
select * from T_City where ProID=3
order by CityID asc;
# 升序并且取前5名
select * from T_City where ProID=3
order by CityID asc
limit 5 ;
#in between
-- 查询广东省宜居城市
select * from T_City
where ProID =( select ProID from T_Province
where ProName = '广东省')
and CityName in ('青岛市','昆明市','三亚市','大连市');
# 查询省份ID在10,20 之间的身份
select * from T_Province
where ProID between 10 and 20 ;

# 查询中国各省份分别有多少地级市
select ProID,count(CityID) from T_City
group by ProID
order by ProID asc;
# 查询中国各身份分别有多少地级市，降序取前5名
select ProID,COUNT(CityID) AS cities from T_City
group by ProID
order by count(CityID) desc
limit 5 ;
# 查询中国各身份分别有多少地级市，取地级市大于10的省份
select ProID,COUNT(CityID) AS cities from T_City
group by ProID
having  cities<10
order by cities desc;
# 查询区县最多的城市
select CityID,count(Id) as disCount from T_District
group by CityID -- 分组
order by disCount desc -- 降序
limit 10 ; -- 显示十行
# 查询中国区县最多的10个城市，打印出城市名字
select * from T_City where CityID in (
    select temp.CityID from
    (
        select CityID,count(Id) disCont from T_District
        group by CityID
        ORDER BY disCont desc
        limit 10
    )temp
);







