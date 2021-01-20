package trabbitmq

import (
	"fmt"
	"log"

	"github.com/lnyyj/dir_links/amqp"
)

//MQURL 连接信息
//url格式 amqp://账号：密码@rabbitmq服务器地址：端口号/vhost
const MQURL = "amqp://guest:guest@rabbitmq-1:5672/"

//rabbitMQ结构体
type RabbitMQ struct {
	conn      *amqp.Connection
	channel   *amqp.Channel
	QueueName string //队列名称
	Exchange  string //交换机名称
	Key       string //bind Key 名称
	Mqurl     string //连接信息
}

//创建结构体实例
func NewRabbitMQ(queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

func checkerr(message string, err error) {
	if err != nil {
		log.Fatalf("%s error: %s", message, err.Error())
		panic(fmt.Sprintf("%s:%s", message, err.Error()))
	}
}
