# 现有最宜居住城市排行榜如下：（“宁波市”，“银川市”，“宜春市”，“宜昌市”,"咸阳市"）
-- 求各省分别拥有多少个宜居城市，降序排列，输出省名，宜居城市数量；
use china;

select ProName,count(CityID) cc from
(
    select ProID,CityID,CityName from T_City
    where CityName in ('宁波市', '银川市', '宜春市', '宜昌市', '咸阳市')
)tc join T_Province tp on tc.ProID = tp.ProID
group by tc.CityID
order by cc desc;

-- 查询哪个城市拥有最多的“旗”，取前10 ；
select CityName,count(DisName) cd from (T_City tc join T_District TD on tc.CityID = TD.CityID)
where DisName like "%旗"
group by CityName
order by cd desc ;

-- 查询省级行政区有哪几种？
select count(ProRemark) from (
    select distinct  ProRemark from T_Province
    where ProRemark is not null
) temp;

-- 查询全国有多少县级市
select * from T_District
where DisName like '%市';

-- 安徽的县级市数量
select count(DisName) from
(T_District td join T_City tc on td.CityID=tc.CityID)
where DisName like '%市' and ProID
=(select ProID from T_Province where ProName='安徽省');

-- 哪个省的县级市最多？
select ProName,temp.cd from
(
select ProID,count(DisName) cd from
(T_District td join T_City tc on td.CityID=tc.CityID)
where DisName like '%市'
group by ProID
)temp join T_Province tp on  temp.ProID=tp.ProID
order by cd desc;



