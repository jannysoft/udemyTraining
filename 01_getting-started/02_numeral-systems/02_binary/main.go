package main

import "fmt" // fmt stands for format package

func main() {
	fmt.Printf("%d - %b \n", 42, 42)
	// Printf is part of fmt package
	// %d is telling go to input the number of 42 in decimal
	// %b is telling go to input the number of 42 in binary
}
