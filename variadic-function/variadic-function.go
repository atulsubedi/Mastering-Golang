package main

import (
	"fmt"
)

func Multi(nums ...int) int {
	total := 1
	for _, num := range nums {
		// total = total * num
		total *= num
	}
	return total
}

func main() {
	fmt.Println(Multi(1, 2, 3, 4))

	// variadic function can be also used in slices

	slice := []int{2, 3, 4, 5}
	fmt.Println(Multi(slice...))
}
