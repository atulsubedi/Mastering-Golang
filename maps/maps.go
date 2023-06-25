package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 44
	m["k2"] = 545

	fmt.Println("maps:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	b := map[string]int{"foo": 1, "bar": 4}

	fmt.Println(b)

	delete(b, "foo")
	fmt.Println(b)
}
