# linux操作系统笔记

## 常用命令
- find ~/desktop/ -name '*.py' //在dektop文件夹下查询.py后缀文件
- whereis python //查找可执行文件
- which python //查看当前使用的是哪个目录

- ln -s /usr/bin/python ~/desktop/  //创建快捷方式
- alias rrm='rm -rf' //给命令起别名
- history //查看命令使用历史

- 命令速查
- rm --help
- man rm
- info rm 

## 系统管理
poweroff 关机
reboot 重启
shutdown -h 12:00 定时关机
shutdown -h now 立刻关机
shutdown -c 取消定时关机

top 查看进程,显示系统所有任务

free -m 查看内存占用，以M为单位
uname -a 打印操作系统信息
uptime -p 查看系统运行时间
echo $PATH 查看环境变量
export PATH=$PATH：/home/sirouyang/Desktop/ 追加环境变量
lsof | head -n 10 查看进程打开的文件
lscpu 查看CPU信息 

time lsof 统计命令执行时间
cal/date 查看日历和时间
date +%y-%m-%d-%H-%M-%S 格式化地查看时间
runlevel 显示当前运行级别
init 6 切换运行级别
 




