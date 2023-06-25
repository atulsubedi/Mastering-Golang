package main

import "fmt"

func main() {
	nums := []int{1, 4, 5, 8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 4 {
			fmt.Println("index", i)
		}
	}

	for i, v := range nums {
		fmt.Println("index:", i, "value", v)
	}

	// range on maps can be used to iterates over key/value pairs

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k, _ := range kvs {
		fmt.Printf("key: %s\n", k)
	}

}
