// it is a package declaration
package main

// it import the go package fmt
import (
	"fmt"
)

// (int,int) in this function signature shows that it return 2 ints
func vals()(int,int){
	return 3,6
}

func multiValue() int {
	a, b := vals()

	fmt.Println(a)
	fmt.Println(b)

	// if we only want a subset of the returned value ,we use the blank identifier also called as comma(,) ok syntax
	_, c := vals()
	fmt.Println(c)

	return 0
}
