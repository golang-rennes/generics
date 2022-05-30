package main

import (
	"fmt"
)

func sort[T ordered](s []T) []T {
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[i] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
	return s
}

type ordered interface {
	// to use <, >
	~int | ~float64 | ~rune | ~string
}

func main() {
	fib := []int{13, 1, 34, 2, 3, 5, 21, 8}
	sortedFib := sort(fib)
	fmt.Println(sortedFib)
}
