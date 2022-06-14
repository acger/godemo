package main

import (
	"math/rand"
	"time"
)

func main() {
	c := make(chan int, 5)

	go func() {
		rand.Seed(time.Now().Unix())

		for i := 0; i < 5; i++ {
			c <- rand.Intn(10)
		}

		close(c)
	}()

	for v := range c {
		println(v)
	}
}
