package main

import "fmt"

func main() {
	// calling func greet and passing in two arguments
	fmt.Println(greet("Jane", " Doe"))
}

func greet(fname, lname string) string {
	// Sprint is a string print func from fmt
	// in this case we use Sprint to return a value to the func greet
	return fmt.Sprint(fname, lname)
}
