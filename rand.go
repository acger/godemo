package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int63n(1 << 10))
	fmt.Println(rand.Int63n(1 << 10))
	fmt.Println(rand.Int63n(1 << 10))
}
