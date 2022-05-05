package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	client, err := clientv3.NewFromURL("http://192.168.93.133:2379")

	if err != nil {
		fmt.Println(err.Error())
	}

	c := client.Watch(context.Background(), "name")

	for r := range c {
		for _, ev := range r.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
