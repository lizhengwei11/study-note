package main

import (
	"github.com/streadway/amqp"
	"log"
)

func failOneError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s :%s", msg, err)
	}
}

func main() {
	//1.尝试连接RabbitMQ,建立连接
	coon, err := amqp.Dial("amqp://coco:kk123123123@127.0.0.1:5672/")
	failOneError(err, "Failed to open a channel")
	//关闭连接
	defer coon.Close()
	//2.创建一个通道，大多数API都是用过通道操作
	channel, err := coon.Channel()
	if err != nil {
		return
	}
	//关闭通道
	defer channel.Close()
	//3.声明要发送到的队列
	declare, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOneError(err, "Failed to declare a queue")

	body := "Hello World!"
	//4.将消息发布声明的队列
	err = channel.Publish(
		"",           // exchange
		declare.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOneError(err, "Failed to publish a message")

}
