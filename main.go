package main

import (
	"fmt"
)

type lesser[T any] interface {
	Less(x T) bool
}

type myInt int

func (x myInt) Less(y myInt) bool {
	return x < y
}

func sort[T lesser[T]](s []T) []T {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j].Less(s[i]) {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
	return s
}

func main() {
	fib := []myInt{13, 1, 34, 2, 3, 5, 21, 8}
	sortedFib := sort(fib)
	fmt.Println(sortedFib)
}
