// 

package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("em:",a)

	// assign 100 to an array at fourth index
	a[4] = 100
	// prints array a
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	fmt.Println("len:",len(a))

	b := [5]int



}