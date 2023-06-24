// it is a package declaration we can read the code from here
package main

// this statement import the go package fmt
import "fmt"

// it is a function for simple addition of two digits
func plus(a int, b int) int {
	return a + b
}

// it is a function for addition of three digits
func plusPlus(a, b, c int) int {
	return a + b + c
}

// the func Functions is the entry point for the go code
// code written here are only executable
func main() {

	// the func plus is called here
	res := plus(1, 2)
	// it gives the result of the called func plus
	fmt.Printf("result is:%d\n", res)

	// the func plusPlus is called here
	res2 := plusPlus(2, 3, 4)

	// it gives the result of the called func plusPlus
	fmt.Printf("result is:%d\n", res2)

}
