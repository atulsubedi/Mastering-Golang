package loop

import (
	"fmt"
)
// var Name = "hello world"
func Loop() int {
	i := 2
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 5; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}
	return 0
}
