package main

import (
	"fmt"
)

func sum(n1, n2 int) int {
	return n1 + n2
}

func plus(a int, b int, c int) int {
	total := a + b + c
	fmt.Println("total:", total)
	return 0
}

func main() {
	res := sum(4, 8)
	fmt.Println("res:", res)

	plus(3, 8, 0)

}
