package main

import "fmt"

// call func greet and pass two arguments to it
// separate the arguments we will pass with comma
func main() {
	greet("Jane", "Doe")
}

// you can declare func with two or more parameters
// in this case fname and lname which are of type string
// we separate the two parameters in the func greet with comma
func greet(fname string, lname string ) {
	// print the passed arguments to parameters fname and l name
	fmt.Println(fname, lname)
}
