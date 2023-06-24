package main

import "fmt"

func main() {

	numbers := []int{2, 3, 4}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range numbers {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}
	dog := map[string]int{"a": 1, "b": 2}
	for k, v := range dog {
		fmt.Printf("%s --> %d\n", k, v)
	}
	for k := range dog {
		fmt.Println("k:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
