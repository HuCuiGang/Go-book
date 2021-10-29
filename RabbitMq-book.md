# 定义和特征

- 1.RabbitMQ是面向消息的中间件，用于组件之间的解耦，主要体现在消息的发送者和消费者之间无强依赖关系﹔

- RabbitMQ特点： 高可用，扩展性，多语言客户端，管理界面等;

- 3.主要使用场景∶流量削峰，异步处理，应用解耦等;

# RabbitMQ服务器安装
- 安装erlang ---依赖于它
1.wget http://www.rabbitmq.com/releases/erlang/erlang-19.0.4-1.el7.centos.x86_64.rpm

- 安装RabbitMQ:
1.wget http://www.rabbitmq.com/releases/rabbitmq-server/v3.6.10/rabbitmq-server-3.6.10-1.el7.noarch.rpm

#  RabbitMQ命令介绍
rpm -ivh erlang //安装erlang
rpm -rabbitMq   //安装rabbitmq
systemctl start rabbitmq-server  //启动
rabbitmqctl stop  //停止
rabbitmq-plugins list // 查看管理插件列表
rabbitmq-plugins enable 插件名 // 启用插件
rabbitmq-plugins disable 插件名 //卸载插件

#  RabbitMQ 管理界面
- 默认端口 127.0.0.1：5672/#/

# docker install rabbitmq
- 1.拉取镜像
docker pull rabbitmq:3.7.7-management

- 2. 根据下载的镜像创建和启动容器
docker run -d --name rabbitmq3.7.7 -p 5672:5672 -p 15672:15672 -v `pwd`/data:/var/lib/rabbitmq --hostname myRabbit -e RABBITMQ_DEFAULT_VHOST=my_vhost  -e RABBITMQ_DEFAULT_USER=admin -e RABBITMQ_DEFAULT_PASS=admin df80af9ca0c9

说明：

-d 后台运行容器；

--name 指定容器名；

-p 指定服务运行的端口（5672：应用访问端口；15672：控制台Web端口号）；

-v 映射目录或文件；

--hostname  主机名（RabbitMQ的一个重要注意事项是它根据所谓的 “节点名称” 存储数据，默认为主机名）；

-e 指定环境变量；（RABBITMQ_DEFAULT_VHOST：默认虚拟机名；RABBITMQ_DEFAULT_USER：默认的用户名；RABBITMQ_DEFAULT_PASS：默认用户名的密码）

5、使用命令：docker ps 查看正在运行容器

6、可以使用浏览器打开web管理端：http://Server-IP:15672


