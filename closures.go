package main

import "fmt"

func return_anoyn_func() func(string) {
	return func(msg string) {
		fmt.Println(msg)
	}
}
 func closures()int {
	print_func := return_anoyn_func()
	print_func("hello i am anonymous function")

	return 0
}