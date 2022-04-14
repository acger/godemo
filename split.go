package main

import (
	"fmt"
	"strings"
)

func main()  {
	str := "chat-pair"
	str = strings.Replace(str, "-", "", -1)
	fmt.Println(str)
}


