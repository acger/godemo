//增加切片元素时有两种情况
//1、底层数组容量足够容纳新元素，则可以直接添加，append后返回的切片与原切片使用相同数组
//2、底层数组容量不足时，则会创建新的数组，append返回的切片指向新数组，原切片依旧使用原数组

package main

import "fmt"

func main() {

	a := [4]int{1, 2, 3, 4}

	b := a[1:3]

	//增加切片容量时，会创建并使用新的底层数组
	b = append(b, 14, 8, 9, 10)

	fmt.Println(a, b)

	c := make([]int, 2, 2)

	c[0] = 10
	c[1] = 10

	c = append(c, 20, 30)

	fmt.Println(c)
}
