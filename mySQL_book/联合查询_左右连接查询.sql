# 联合查询
-- union
-- 中国各省的ID，省份名字
use china;
select ProID,ProName from T_Province
union
-- 广东所有地级市ID，名字
select CityID,CityName from T_City
where ProID = (
    select (select ProID from T_Province where ProName='河北省')
);


-- 查询中国有多少地级市

-- 查询河北省有多少地级市
# select count(CityID) from T_City
# where ProID = (
#     select ProID from T_Province where ProName='河北省'
#     );
select ProName,CityName from ( T_City join T_Province
    on T_City.ProID = T_Province.ProID)
where ProName = '河北省';
# 统计各省地级市的数量，输出省名、地级市数量
select T_City.ProID,ProName,count(CityID) as cc from(
T_City join T_Province
on T_City.ProID = T_Province.ProID
)
group by ProID
order by cc asc;

# 求每个省份中最大的城市ID，城市名称

# 地级市对多的省份取前10名

# 查询拥有区县最多的城市的前10名

# 查询拥有20个以上区县的城市，输出城市名、区县数量
select CityName,count(td.Id) disCount from (T_City tc join T_District td
    on tc.CityID = td.CityID)
group by CityName
having disCount >20 ;

#打出拥有20个以上区县的城市名字，子查询实现

# 北京所有区县

# 区县最多的城市(直辖市不要)是哪个省的什么市，查询结果包含省名、市名，区县数量；
select tp.ProName,tcd.CityName,tcd.ci from(
           select ProID, CityName, count(Id) ci from (
              T_City tc join T_District td on tc.CityID = td.CityID
                  )
            group by tc.CityID
            order by ci desc
           )tcd join T_Province tp on tcd.ProID=tp.ProID
where tp.ProRemark != '直辖市'
limit 1;


# 内连接 （inner join） -基于左右两表公共部分
# 左连接 （left join） -基于左右两表的公共部分 + 左表特有的部分
# 右连接  （rigth join）- 基于左右两表公共的部分 + 右表特有的部分

select ProName,CityName from(T_Province tp left join T_City tc
    on tp.ProID=tc.ProID);

-- 插入省表和市独有部分
insert into T_Province(ProName) values ("日本省");
insert into T_City(CityName) values("洛杉矶市");

select * from T_Province;






