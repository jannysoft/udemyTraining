package main

import "fmt"

func main() {
	z := 2

	for x := 0; x < 100; x++ {
		if x%z == 1 {
			fmt.Println("odd")
		} else {
			fmt.Println("even")
		}
	}

}
