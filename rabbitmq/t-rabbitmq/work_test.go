package trabbitmq

import (
	"strconv"
	"testing"
	"time"
)

func Test_WorkMQ(t *testing.T) {
	rmq := NewRabbitMQSimple("test_mq")
	defer rmq.Destory()
	go rmq.ConsumeSimple("aa")
	go rmq.ConsumeSimple("bb")

	for i := 0; ; i++ {
		rmq.PublishSimple("hello " + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
	}
}
