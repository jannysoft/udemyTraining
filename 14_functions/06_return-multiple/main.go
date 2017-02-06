package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane ", "Doe "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}


/*
In here we have func main calling func greet and asking if func greet can assign two arguments
 of type string that has been passed to it so func main can print it

func greet then says yes I do have two parameters of type string that I can assign arguments to,
but then I need the func main to print them in the following order:

Case one, please print fname first and lname second
Case two, please print lname first and fname second

*/