package main

import "fmt"
// declaring multi CONST with assigned IOTA to it
const (
	A = iota // 0
	B        // 1
	C        // 2
)

// declaring another multi CONST with assigned IOTA to it
const (
	D = iota // 0
	E        // 1
	F        // 2
)

func main() {
	// print the multi CONSTs
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
}
