package main

import "fmt"

func main() {
	// i := 0 is the initialising part of the loop
	// i <= 100 is the condition of the loop (while this condition is true it will continue looping)
	// after a whole loop it will execute i++ which means it will add 1 to i every time the loop is executed
	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}
}
