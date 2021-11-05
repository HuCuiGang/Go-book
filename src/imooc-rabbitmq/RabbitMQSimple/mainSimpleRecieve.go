package RabbitMQSimple

import "imooc-rabbitmq/RabbitMQ"

func main() {
	rabbitmq:=RabbitMQ.NewRabbitMQSimple("imoocSimeple")
	rabbitmq.ConsumeSimple()
}
