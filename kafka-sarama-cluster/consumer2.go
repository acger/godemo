package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	_ "regexp"

	cluster "github.com/bsm/sarama-cluster"
)

/**
  消费者
*/
func main() {

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true

	// init consumer
	brokers := []string{"127.0.0.1:9092"}
	//可以订阅多个主题
	topics := []string{"topic_0", "topic_1", "topic_2", "topic_3", "topic_4"}
	consumer, err := cluster.NewConsumer(brokers, "my-consumer-group2", topics, config)
	if err != nil {
		panic(err)
	}
	//这里需要注意的是defer 一定要在panic 之后才能关闭连接
	defer consumer.Close()

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// consume errors
	go func() {
		for err := range consumer.Errors() {
			log.Printf("Error: %s\n", err.Error())
		}
	}()

	// consume notifications
	go func() {
		for ntf := range consumer.Notifications() {
			log.Printf("Rebalanced: %+v\n", ntf)
		}
	}()

	// 循环从通道中获取message
	//msg.Topic 消息主题
	//msg.Partition  消息分区
	//msg.Offset
	//msg.Key
	//msg.Value 消息值
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "") // 上报offset
			}
		case <-signals:
			return
		}
	}
}
