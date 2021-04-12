package main

import "fmt"

/**
 *	练习：斐波纳契闭包
 */

// 返回一个“返回int的函数”
func fibonacci() func() int {
	a, b, num, i := 0, 1, 0, 0

	return func() int {
		switch i++; i {
		case 1:
			return a
		case 2:
			return b
		default:
			num = a + b
			a = b
			b = num
			return num
		}

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
