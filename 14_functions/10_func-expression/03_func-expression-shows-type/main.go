package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hello World")
	}

	greeting()

	fmt.Printf("%T \n", greeting)
}

// same as previous example but here we print the type of greeting