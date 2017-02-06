package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Fprintln(x)
		y := "The credit belongs with the one who is in the ring"
		fmt.Fprintln(y)
	}
	// fmt.Println(y) // outside scope of y
}