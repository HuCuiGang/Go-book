package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url 格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const  MQURL = "amqp://admin:admin@127.0.0.1:5672/my_vhost"

type RabbitMQ struct {
	conn *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//key
	Key string
	//链接信息
	Mqurl string
}
//创建RabbitMQ结构体实例
func NewRabbitMQ(queueName string,exchange string ,key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{QueueName: queueName,Exchange: exchange,Key: key,Mqurl: MQURL}
	var err error
	//创建rabbitmq链接
	rabbitmq.conn,err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err,"创建链接错误！")
	rabbitmq.channel,err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err,"获取channel失败")
	return rabbitmq
}

//断开channel和connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error,message string) {
	if err != nil {
		log.Fatalf("%s:%s",message,err)
		panic(fmt.Sprintf("%s:%s",message,err))
	}
}
// 简单模式step: 1.创建简单模式下的RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName,"","")
}

//简单模式step:2 2.简单模式下生产代码
func (r *RabbitMQ) PublishiSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证了队列存在，且消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外熟悉
		nil, 
		)
	if err!=nil {
		fmt.Println(err)
	}
	//2.发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，会根据exchange 类型和routkey规则，如果无法找到符合条件的队列那么会吧发送的消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，则会吧消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})

}

//简单模式step:3 3.简单模式下消费者接收消息代码
func (r *RabbitMQ) ConsumeSimple(){
	//1.申请队列，如果队列不存在会自动创建，如果存在则跳过创建
	//保证了队列存在，且消息能发送到队列中
	_,err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否为自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞
		false,
		//额外属性
		nil,
	)
	if err!=nil {
		fmt.Println(err)
	}
	//2.接收消息
	msgs,err:=r.channel.Consume(
		r.QueueName,
		//用来区分多个消费者
		"",
		//是否自动应答
		true,
		//是否具有排他性
		false,
		//如果设置为true，表示不能将同一个connection中发送的消息传递给这个connection中的消费者
		false,
		//队列消费是否阻塞
		false,
		//额外属性
		nil,
	)
	if err!=nil{
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for d:= range msgs{
			//实现我们要处理的逻辑函数
			log.Printf("Received a message:%s",d.Body)
			fmt.Println(d.Body)
		}
	}()

	log.Println("[*] Waiting for messages,To exit pressw CTRL+C ")
	<-forever

}


