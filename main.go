package main

import (
	"fmt"
	"strconv"
)

type myInt int

func (i myInt) String() string {
	return strconv.Itoa(int(i))
}

type myString string

func (s myString) String() string {
	return string(s)
}

func print[T fmt.Stringer](x T) {
	fmt.Printf("%s\n", x.String())
}

func main() {
	print(myInt(42))
	print(myString("foo"))
}
