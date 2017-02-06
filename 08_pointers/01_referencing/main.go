package main

import "fmt"

func main() {

	// declare VAR and assign value of 43 to it
	a := 43

	// print var a
	fmt.Println(a)
	// print where in memory VAR a is stored
	fmt.Println(&a)

	// declare VAR b and assign the info of where about VAR a is stored in the memory
	// VAR b is a pointer to a int type
	// that means * is pointing to memory address where int(integer is stored)
	var b *int = &a


	//print VAR b
	fmt.Println(b)

	// the above code makes b a pointer to the memory address where an int is stored
	// b is of type "int pointer"
	// *int -- the * is part of the type -- b is of type *int
}
