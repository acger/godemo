package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(1<<63 - 1)

	a := strconv.FormatInt(9223372036854775807, 36)

	fmt.Println(a)

	fmt.Println(1 << 27)
	fmt.Println(1024 * 1024 * 128)
}
