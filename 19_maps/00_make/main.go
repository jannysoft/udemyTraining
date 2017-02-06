package main

import "fmt"

func main() {
	var myG = make(map[string]string)

	myG["Tim"] = "Good morning"
	myG["Jimmy"] = "Bonjour"
	myG["Chris"] = "Privet"

	fmt.Println(myG)
}
