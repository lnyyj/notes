package trabbitmq

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/lnyyj/dir_links/amqp"
)

func Test_FinoutDeadLetterMain(t *testing.T) {
	rmqconn, err := amqp.Dial(MQURL)
	checkerr("connector", err)
	defer rmqconn.Close()

	ch, err := rmqconn.Channel()
	checkerr("channel", err)
	defer ch.Close()

	err = ch.ExchangeDeclare("dlx.test", "fanout", false, false, true, false, nil)
	checkerr("dlx exchange", err)

	dlxQueue, err := ch.QueueDeclare("dlx.test", false, false, false, false, nil)
	checkerr("dlx queue", err)

	err = ch.QueueBind(dlxQueue.Name, "dlxKey", "dlx.test", false, nil)
	checkerr("dlx queue bind", err)

	args := make(map[string]interface{}, 0)
	args["x-message-ttl"] = 5000 // 单位毫秒
	args["x-dead-letter-exchange"] = "dlx.test"
	args["x-dead-letter-routing-key"] = "dlxKey"
	err = ch.ExchangeDeclare("normal.test", "fanout", false, false, false, false, nil)
	checkerr("normal exchange", err)

	normalQueue, err := ch.QueueDeclare("normal.test", false, false, false, false, args)
	checkerr("normal queue", err)

	err = ch.QueueBind(normalQueue.Name, "normalKey", "normal.test", false, nil)
	checkerr("normal queue bind", err)

	ch.Qos(
		5,     // prefetchCount：会告诉RabbitMQ不要同时给一个消费者推送多于N个消息，即一旦有N个消息还没有ack，则该consumer将block掉，直到有消息ack
		0,     // prefetchSize：最多传输的内容的大小的限制，0为不限制，但据说prefetchSize参数，rabbitmq没有实现
		false, // global：true\false 是否将上面设置应用于channel，简单点说，就是上面限制是channel级别的还是consumer级别
	)

	{ // 消费者
		normalMsgs, err := ch.Consume(normalQueue.Name, "normal1", false, false, false, false, nil)
		checkerr("normal msg", err)
		go func() {
			for msg := range normalMsgs {
				log.Printf("normal 1 msg: %s", msg.Body)
				// multiple：为了减少网络流量，手动确认可以被批处理，当该参数为 true 时，则可以一次性确认 delivery_tag 小于等于传入值的所有消息
				// msg.Ack(false) // 手动应答， 自动应答为false有效
				// msg.Nack(true, false)
			}
		}()
	}
	{ // 消费者
		normalMsgs, err := ch.Consume(normalQueue.Name, "normal2", false, false, false, false, nil)
		checkerr("normal 2 msg", err)
		go func() {
			for msg := range normalMsgs {
				log.Printf("normal 2 msg: %s", msg.Body)
				// multiple：为了减少网络流量，手动确认可以被批处理，当该参数为 true 时，则可以一次性确认 delivery_tag 小于等于传入值的所有消息
				msg.Ack(false) // 手动应答， 自动应答为false有效
				// msg.Nack(true, false)
			}
		}()
	}
	{
		dlxMsgs, err := ch.Consume(dlxQueue.Name, "dlx", false, false, false, false, nil)
		checkerr("dlx msg", err)
		go func() {
			for msg := range dlxMsgs {
				log.Printf("dead letter msg: %s", msg.Body)
				msg.Ack(false) // 手动应答， 自动应答为false有效
			}
		}()
	}

	for i := 0; ; i++ {
		ch.Publish("normal.test", "normal.test", false, false, amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(fmt.Sprintf("hi %d", i)),
		})
		time.Sleep(1 * time.Second)
	}
	select {}

}

func Test_PubSubMQ(t *testing.T) {
	rabbitmq := NewRabbitMQPubSub("test_mq_pubsub")
	qGoiot, qYKC := "goiot_ac", "ykc_ac"
	go rabbitmq.RecieveSub(qGoiot)
	// go rabbitmq.RecieveSub(qYKC)
	for i := 0; ; i++ {
		if i%2 == 0 {
			rabbitmq.PublishPub(fmt.Sprintf("push %d", i), qGoiot)
		} else {
			rabbitmq.PublishPub(fmt.Sprintf("push %d", i), qYKC)
		}
		time.Sleep(1 * time.Second)
	}
}

//订阅模式创建RabbitMQ实例
func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ("", exchangeName, "")
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//订阅模式生产
func (r *RabbitMQ) PublishPub(message string, queueName string) {
	// //1.尝试创建交换机
	var err error
	err = r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false, //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,
		nil,
	)
	r.failOnErr(err, "pub创建交换机错误")

	//2.发送消息
	err = r.channel.Publish(
		r.Exchange,
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		fmt.Println("pub error:" + err.Error())
	}
}

//订阅模式消费端代码
func (r *RabbitMQ) RecieveSub(queueName string) {
	//1.试探性创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // 交换机类型
		true,     // 是否持久化
		false,    // 自动删除
		false,    // YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
		false,    // 列是否阻塞
		nil,
	)
	r.failOnErr(err, "消费端创建交换机错误")
	//2.试探性创建队列，这里注意队列名称不要写
	q, err := r.channel.QueueDeclare(
		queueName, //随机生产队列名称
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "sub "+queueName+" queue error")

	//绑定队列到 exchange 中
	err = r.channel.QueueBind(
		q.Name,
		"", //在pub/sub模式下，这里的key要为空
		r.Exchange,
		false,
		nil)

	//消费消息
	messges, err := r.channel.Consume(
		q.Name,
		"aa",
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
