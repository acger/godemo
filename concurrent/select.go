package main

import (
	"fmt"
	"time"
)

func main() {

	var c chan int

	c = make(chan int)

	go func() {
		for {
			select {
			case x, ok := <-c:
				fmt.Printf("%v", x)
				fmt.Printf("%+v", x)
				fmt.Printf("%#v", x)
				println(ok)

				if ok == false {
					c = nil
				}
			}
		}
	}()

	c <- 1
	c <- 2
	c <- 3

	time.Sleep(time.Second * 1)

	close(c)

	time.Sleep(time.Second * 1)
}
