package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(1<<63 - 1)

	a := strconv.FormatInt(9223372036854775807, 36)

	fmt.Println(a)
}
