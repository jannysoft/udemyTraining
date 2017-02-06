package main

import "fmt"

func main() {

	// call in func greet and pass two arguments to it
	fmt.Println(greet("Jane", " D0e"))
}

// declare func with named return value in this case "s"
func greet(fname, lname string) (s string) {

	// assign string to s
	s = fmt.Sprint(fname, lname)
	// return the outcome to s
	return
}
