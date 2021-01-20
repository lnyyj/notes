package trabbitmq

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/lnyyj/dir_links/amqp"
)

func Test_SimpleMain(t *testing.T) {
	amqp.Dial(MQURL)
}

func Test_SimpleMQ(t *testing.T) {
	rmq := NewRabbitMQSimple("test_mq")
	if rmq == nil {
		return
	}
	defer rmq.Destory()
	go rmq.ConsumeSimple("aa")

	for i := 0; ; i++ {
		rmq.PublishSimple(fmt.Sprintf("send msg %d", i))
		time.Sleep(1 * time.Second)
	}
}

//创建简单模式下RabbitMQ实例
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// 创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error
	// 获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabb"+"itmq!")
	// 获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		return nil
	}
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//直接模式队列生产
func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞处理
		nil,   //额外的属性
	)
	if err != nil {
		fmt.Printf("发送消息声明队列错误: %s\r\n", err.Error())
	}
	//调用channel 发送消息到队列中
	r.channel.Publish(
		r.Exchange,
		q.Name, // r.QueueName
		false,  //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,  //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(consumeName string) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞处理
		nil,   //额外的属性
	)
	if err != nil {
		fmt.Printf("消费者 %s 声明队列错误: %s\r\n", consumeName, err.Error())
	}

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,      // queue
		consumeName, // consumer 用来区分多个消费者
		false,       // auto-ack 是否自动应答
		false,       // exclusive 是否独有
		false,       // no-local 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,       // no-wait 列是否阻塞
		nil,         // args
	)
	if err != nil {
		fmt.Printf("消费者 %s 接收消息错误: %s\r\n", consumeName, err.Error())
	}

	forever := make(chan bool)
	//启用协程处理消息
	go func() {
		for {
			// d := range msgs
			select {
			case d := <-msgs: // 消息逻辑处理，可以自行设计逻辑
				log.Printf("consume %s Received a message: %s", consumeName, d.Body)
				d.Ack(false) // 手动应答， 自动应答为false有效
			}

		}
	}()
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
