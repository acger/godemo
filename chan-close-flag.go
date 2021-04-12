//有缓存通道的第二个赋值标记位"ok"，表示的是通道已关闭并且缓存已清空

package main

import (
	"fmt"
	"os"
)

func main() {
	var c chan int

	c = make(chan int, 10)

	for i := 0; i < 9; i++ {
		c <- i
	}

	close(c)

	go func() {
		for {
			//有缓存通道的第二个赋值标记位"ok"，表示的是通道已关闭并且缓存已清空
			v, ok := <-c

			if !ok {
				fmt.Println("break")
				os.Exit(1)
				break
			}

			fmt.Println(v)
		}
	}()

	for {}
}
