package RabbitMQSimple

import (
	"fmt"
	"imooc-rabbitmq/RabbitMQ"
)

func main(){
	rabbitmq:= RabbitMQ.NewRabbitMQSimple("imoocSimeple")
	rabbitmq.PublishiSimple("hello imooc")
	fmt.Println("发送成功")
}
