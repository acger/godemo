package main

import (
	"fmt"
	"github.com/VictoriaMetrics/fastcache"
)

func main() {
	c := fastcache.New(1 << 27)
	c.Set([]byte("name"), []byte("illya"))
	r := c.Get(nil, []byte("name"))
	fmt.Println(string(r))
	fmt.Println("ok")
}
