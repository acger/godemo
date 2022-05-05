package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

var file io.Writer

func init() {
	file, _ = os.OpenFile("./concurrent/errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func main() {
	f := func() {
		file.Write([]byte("are you ok"))
		fmt.Println("foo2")
	}

	time.Sleep(time.Second * 2)
	c := spawn(f)
	fmt.Println("bar")
	c <- 1
	fmt.Println("end")
}

func spawn(f func()) chan int {
	ci := make(chan int)

	go func(c chan<- int) {
		fmt.Println("in")
		<-ci
		f()
	}(ci)

	return ci
}
