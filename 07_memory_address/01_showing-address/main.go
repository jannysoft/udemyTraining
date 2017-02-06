package main

import "fmt"

func main() {
	//declare var a and assign value of 43
	a := 43

	// print the value of VAR a
	fmt.Println("a - ", a)
	// print the memory index of VAR a with operator "&"
	fmt.Println("a's memory address - ", &a)
	fmt.Printf("%d \n", &a)
}
