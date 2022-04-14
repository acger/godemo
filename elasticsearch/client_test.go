package main

import (
	es "github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestNewClient(t *testing.T)  {
	client,err := es.NewClient(es.Config{
		Addresses: []string{"http://dev.ll:9200"},
		Username: "elastic",
		Password: "123456",
	})

	if err != nil{
		t.Fatal(err)
	}

	t.Log(client.Info())
}
