// it is a package declaration, we can read code starting from here 
package main

// it import the go package fmt 
import (
	"fmt"

	  l "github.com/atulsubedi/go-by-example/1-loop"
)

// it is the entry point for go code , the codes which are written here are only executable
func main() {
	  fmt.Println(l.Loop())
	  fmt.Println(Ifelse())
	  fmt.Println(Slices())
	  fmt.Println(Array())
	  fmt.Println(Switch())
	  fmt.Println(Slices())
	  fmt.Println(Maps())
	  fmt.Println(Range())
	  fmt.Println(Functions())
	  fmt.Println(multiValue())
	  fmt.Println(Variadic())
	  fmt.Println(closures())
	  fmt.Println(recursion())
	  fmt.Println(userCh())
}
