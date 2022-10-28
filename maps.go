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
	fmt.Println("v1:",v1)

	fmt.Println("len:",len(m))

	// Deleting slices using builtin delete

	delete(m,"k2")
	fmt.Println("map:",m)
	
	_,prs := m["k2"]
	fmt.Println("prs:",prs)

	n := map[string]int {"boo" : 1 , "foo" : 2 }
	fmt.Println("map:",n)

	return 0
}
