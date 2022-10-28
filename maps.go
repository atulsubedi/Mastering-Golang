package main

import "fmt"

func Maps() int {

	// to create a empty map use the builtin make:
	//  make(map[key-type]val-type)

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("maps:", m)
	

	v1 := m["k1"]
	fmt.Println("v1:",v1             )
	return 0
}
