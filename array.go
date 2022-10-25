package main

import "fmt"

func Array() int {
	var a [5]int
	fmt.Println("em:",a)

	// assign 100 to an array at fourth index
	a[4] = 100
	// prints array a
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])
	fmt.Println("len:",len(a))

	 b := [5]int{1,2,3,4,5}
	 fmt.Println("dcl:",b)

	 var twoD [2][3]int 
	 for i:= 0; i<2; i++{
		for j := 0; j<3; j++{
			twoD [i][j] = i + j 
		}
	 }
	 fmt.Println("2d:",twoD)
	return 0

}