package main

import "fmt"

func main() {
	var arr = []int{200, 1, 80, 7, 91, 3, 22, 107, 24}
	fmt.Println(quicksort(arr))
}

func quicksort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}

	var left, right []int
	var pivot = arr[0]

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quicksort(left)
	right = quicksort(right)

	return append(append(left, pivot), right...)
}
