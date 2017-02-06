package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Hello Tim")
	case "Jenny":
		fmt.Println("Hello Jenny")
	case "Marcus":
		fmt.Println("Hello Marcus")
		fallthrough
	case "Medhi":
		fmt.Println("Hello Medhi")
		fallthrough
	case "Julian":
		fmt.Println("Hello Julian")
	case "Sushant":
		fmt.Println("Hello Sushant")
	}
}
