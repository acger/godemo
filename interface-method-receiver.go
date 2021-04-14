//使用指针接收者来实现一个接口，那么只有指向那个类型的指针才能够实现对应的接口
//使用值接收者来实现一个接口，那么那个类型的值或者指针都能够实现该接口
package main

import "fmt"

type Foo struct {
	Name string
}

func main() {
	f := Foo{"foo"}

	fmt.Fprintf(&f, "\n")
}

func (f Foo) Write(b []byte) (int, error) {
	return 0, nil
}
