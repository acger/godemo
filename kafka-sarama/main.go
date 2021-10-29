package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func consumer_test() {
	fmt.Printf("consumer_test")

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Version = sarama.V2_8_1_0

	// consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	defer consumer.Close()

	partition_consumer, err := consumer.ConsumePartition("chat-pair", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Printf("try create partition_consumer error %s\n", err.Error())
		return
	}
	defer partition_consumer.Close()

	for {
		select {
		case msg := <-partition_consumer.Messages():
			fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
				msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
		case err := <-partition_consumer.Errors():
			fmt.Printf("err :%s\n", err.Error())
		}
	}

}

func main(){
	consumer_test()
}
