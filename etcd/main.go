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

	r, e := client.Put(context.TODO(), "name", "kk")

	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Println(r.Header.Revision)
}
