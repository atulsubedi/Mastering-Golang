package main

import "fmt"

func Slices() int {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("len:", len(s))

	s = append(s, "d", "e", "f")
	fmt.Println("append:", s)
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("copy'd:", c)
	l := s[2:5]
	fmt.Println("Sl1:", l)
	l = s[:5]
	fmt.Println("Sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g","h","i"}
	fmt.Println("del:",t)
	
	twoD := make([][]int,3)
	for i:= 0; i<3; i++{
		innerlen := i+1
		twoD[i] = make([]int,innerlen)
		for j:=0; j<innerlen; j++{
			twoD[i][j] = i+j
		}
	}
	fmt.Println("2d:",twoD)
	return 0
}
