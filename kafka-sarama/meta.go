package main

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func metadata_test() {
	fmt.Printf("metadata test\n")

	config := sarama.NewConfig()
	config.Version = sarama.V0_11_0_2

	client, err := sarama.NewClient([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("metadata_test try create client err :%s\n", err.Error())
		return
	}

	defer client.Close()

	// get topic set
	topics, err := client.Topics()
	if err != nil {
		fmt.Printf("try get topics err %s\n", err.Error())
		return
	}

	fmt.Printf("topics(%d):\n", len(topics))

	for _, topic := range topics {
		fmt.Println(topic)
	}

	// get broker set
	brokers := client.Brokers()
	fmt.Printf("broker set(%d):\n", len(brokers))
	for _, broker := range brokers {
		fmt.Printf("%s\n", broker.Addr())
	}
}

func main()  {
	metadata_test()
}