package main

import (
	"fmt"
	es "github.com/elastic/go-elasticsearch/v7"
	"log"
)

func main() {
	client, err := es.NewClient(es.Config{
		Addresses: []string{"http://dev.ll:9200"},
		Username:  "elastic",
		Password:  "123456",
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(client.Info())
}
