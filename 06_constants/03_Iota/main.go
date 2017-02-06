package main

import "fmt"

// IOTA IS incrementing value
// declaring and assigning IOTA value to multiple CONST
const (
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

/*

Top CONST can Also Be done like that

const (
	A = iota // 0
	B 	 // 1
	C 	 // 2
)

 */

func main() {
	// print A, B, C
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
