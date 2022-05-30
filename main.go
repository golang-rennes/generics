package main

import (
	"fmt"
)

func filter[T any](s []T, pred func(T) bool) []T {
	res := make([]T, 0, len(s))
	for _, x := range s {
		if pred(x) {
			res = append(res, x)
		}
	}
	return res
}

func isEven(x int) bool {
	return x%2 == 0
}

func main() {
	fib := []int{1, 2, 3, 5, 8, 13, 21, 34}
	evenFib := filter(fib, isEven)
	fmt.Println(evenFib)
	fmt.Println()
}
