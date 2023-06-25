package main

import (
	"fmt"
)

func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0

	for _, v := range nums {
		total += v
	}
	fmt.Println(total)

}

func main() {
	sum(1, 4, 8)
	nums := []int{1, 9, 8}
	sum(nums...)

}
