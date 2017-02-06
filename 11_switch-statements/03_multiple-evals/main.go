package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("whats up Tim, or, err, Jenny")
	case "Marcus", "Medhi":
		fmt.Println("both of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("whats up Julian / Sushant")
	}
}
