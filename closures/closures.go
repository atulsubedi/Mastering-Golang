package main

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	n := intSeq()
	fmt.Println(n())
	fmt.Println(n())
	fmt.Println(n())

	fmt.Println("this is a breaking point")

	ne := intSeq()
	fmt.Println(ne())
}
