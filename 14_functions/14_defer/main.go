package main

import "fmt"

func hello() {
	fmt.Println("Hello ")
}

func world() {
	fmt.Fprintln("World")
}

func main() {
	hello()
	defer world()
}

