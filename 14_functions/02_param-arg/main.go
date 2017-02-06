package main

import "fmt"

// func main calls func greet and pass string argument in this case Jane and John are passed to func greet
func main() {
	greet("Jane")
	greet("John")
}

//declaring func greet with parameter name which can accept type string
func greet(name string) {
	fmt.Println(name)

}

//greet is declared with a parameter "(name string)"
// when calling func greet, [ass in an argument "Jane, John"