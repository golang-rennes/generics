package main

import (
	"fmt"
)

func index[T comparable](s []T, x T) int {
	for i, y := range s {
		if x == y {
			return i
		}
	}
	return -1
}

func main() {
	fib := []int{1, 2, 3, 5, 8, 13, 21, 34}
	i := index(fib, 8)
	fmt.Println(i)
}
