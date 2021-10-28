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




