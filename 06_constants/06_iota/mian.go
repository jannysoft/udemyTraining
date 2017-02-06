package main

import "fmt"

// another case of declaring CONST with value iota
const (
	// trow away the value of first CONST because we dont need it
	_ = iota //0
	// multiply IOTA times 10
	B = iota * 10 // 1 * 10
	// multiply IOTA times 10
	C = iota * 10 // 2 * 10
)

func main() {
	fmt.Println(B)
	fmt.Println(C)
}
