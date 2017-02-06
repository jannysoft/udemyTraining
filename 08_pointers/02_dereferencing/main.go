package main

import "fmt"

func main() {

	a := 43
	fmt.Println(a) // 43
	fmt.Println(&a) // 0xc042008280

	var b *int = &a // declaring VAR b with pointer int to memory address

	fmt.Println(b) // 0xc042008280
	// dereferencing the value of b pulled from the memory address where a is stored
	fmt.Println(*b) // 43

}

/*
b is an int pointer
b points to the memory address where an int is stored
to see the value in that memory address, add a * in front of b
this is known as dereferencing
the * is an operator in this case
*/