package main

import "fmt"

const (
	_  = iota             // assign 0 and trow it away
	// move 1 in front of ten 0 digits
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	// move 1 in front of twenty  0 digits
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	// move 1 in front of thirty  0 digits
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {
	fmt.Println("binaty\t\tdecimal")
	// print CONST KB in binary a tab
	fmt.Printf("%b\t", KB)
	// print CONST KB in decimal and add new line
	fmt.Printf("%d\n", KB)
	// print CONST MB in binary and a tab
	fmt.Printf("%b\t", MB)
	// print CONST MB in decimal and add new line
	fmt.Printf("%d\n", MB)
	// print CONST GB in binary and a tab
	fmt.Printf("%b\t", GB)
	// print CONST GB in decimal and add new line
	fmt.Printf("%d\n", GB)

}