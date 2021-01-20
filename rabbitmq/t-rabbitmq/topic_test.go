package trabbitmq

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/lnyyj/dir_links/amqp"
)

//话题模式
//创建RabbitMQ实例
func NewRabbitMQTopic(exchangeName, routingKey, queueKey string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueKey, exchangeName, routingKey)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//话题模式发送消息
func (r *RabbitMQ) PublishTopic(message string) {
	//1.尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic", //要改成topic
		true,
		false,
		false,
		false,
		nil,
	)

	r.failOnErr(err, "Failed to declare an exchange")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		r.Key, //要设置
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//话题模式接受消息
//要注意key,规则
//其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
//匹配 qlf.* 表示匹配 qlf.hello, 但是qlf.hello.one需要用qlf.#才能匹配到
func (r *RabbitMQ) RecieveTopic() {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic", //交换机类型
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exch"+"ange")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		r.QueueName, //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare a queue")

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		r.Key, //在pub/sub模式下，这里的key要为空
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range messges {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("退出请按 CTRL+C")
	<-forever
}

func Test_TopicMQ(t *testing.T) {
	rmq1 := NewRabbitMQTopic("card_pay", "occupy_order_cardpay", "occupy_order_card_pay")
	rmq2 := NewRabbitMQTopic("card_pay", "order_cardpay", "order_card_pay")
	defer rmq1.Destory()
	defer rmq2.Destory()

	go rmq1.RecieveTopic()
	go rmq2.RecieveTopic()

	for i := 0; i < 1000; i++ {
		if i%2 == 1 {
			rmq1.PublishTopic("hello " + strconv.Itoa(i))
		} else {
			rmq2.PublishTopic("hi " + strconv.Itoa(i))
		}
		time.Sleep(1 * time.Second)
	}
}
