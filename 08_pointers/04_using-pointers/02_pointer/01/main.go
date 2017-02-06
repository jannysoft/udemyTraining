package main

import "fmt"

// 3. value of x is received by FUNC zero with pointer to the memory address where x is stored
func zero(z *int) {
	// 4. z is saying I need that value we have received from the pointer to be changed to 0
	// that is called dereferencing
	*z = 0
}

func main() {
	x := 5   // 1. assigning value of 5 to VAR x
	zero(&x) // 2.  passing the memory address to func zero where the value of x is stored

	// 5. print the new value of x that has been assign from FUNC zero
	fmt.Println(x) // x is 0
}
