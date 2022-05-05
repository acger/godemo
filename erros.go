package main

import (
	"errors"
	"fmt"
)

type MyError struct {
	e string
}

func (e MyError) Error() string {
	return e.e
}

func main() {

	var ErrFoo = errors.New("Error Foo")
	var ErrBar = fmt.Errorf("Error Bar \n%w", ErrFoo)
	var ErrZab = fmt.Errorf("Error Zab \n%w", ErrBar)

	fmt.Println(errors.Unwrap(ErrZab))

	e := ErrZab
	is := errors.Is(e, ErrBar)
	fmt.Println(is)
	fmt.Println(e == ErrBar)

	var err = MyError{"my error"}
	err1 := fmt.Errorf("err1\n%w", err)
	err2 := fmt.Errorf("err2\n%w", err1)

	var me MyError
	me = MyError{"my error"}
	as := errors.As(err2, &me)
	fmt.Println(as)
}
