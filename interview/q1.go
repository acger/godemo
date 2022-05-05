package main

import "sync"

func main() {
	var w sync.WaitGroup

	var (
		num  = make(chan struct{})
		char = make(chan struct{})
	)

	go func() {
		var i uint8
		for {
			select {
			case <-num:
				i++
				println(i % 10)
				i++
				println(i % 10)
				char <- struct{}{}
			}
		}
	}()

	go func() {
		w.Add(1)
		var c = 'A'
		for {
			select {
			case <-char:
				if c > int32('Z') {
					w.Done()
					return
				}

				println(string(c))
				c++

				println(string(c))
				c++
				num <- struct{}{}
			}
		}
	}()

	num <- struct{}{}

	w.Wait()
}
