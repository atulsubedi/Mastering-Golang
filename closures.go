package main

import "fmt"
// creating an anonymous function
func return_anoyn_func() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}
func int_seq () func()int {
	i := 0
	return func()int{
		i++
		return i
	}
}
 func closures()int {
	// anonymous function declared
	print_func := return_anoyn_func()
	print_func("hello i am anonymous function")

	next_int := int_seq()
	println(next_int())	
	println(next_int())
	println(next_int())
	
	next_int2 := int_seq()
	println(next_int2())
	return 0
}