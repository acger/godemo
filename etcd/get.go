package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	client, err := clientv3.NewFromURL("http://192.168.93.133:2379")

	if err != nil {
		fmt.Println(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.
		Second)
	resp, e := client.Get(ctx, "name", clientv3.WithRev(536))

	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(string(resp.Kvs[0].Value))
	fmt.Println(resp.Kvs[0].Version)

	fmt.Printf("%v", resp)
	cancel()
}
