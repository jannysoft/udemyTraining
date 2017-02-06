package main

import "fmt"

// declare func main with no receiver no parameters and no return values
func main() {
	// declare var greeting and assign the value of anonymous function to it
	// assigning func to variable is also called func expression

	// !!! IMPORTANT the only way to have a func inside another func is to assign it to a variable

	greeting := func() {
		// print string hello world
		fmt.Println("Hello, World !!")
	}
	// call the var greeting which now has to be called as func
	greeting()
}

