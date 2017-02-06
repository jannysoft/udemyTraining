package main

import "fmt"

//switch on types
// -- normally we switch on value of variables
// -- go allows you to switch on type of variables
type Contact struct {
	greeting string
	name     string
}

//we'll learn more about interfaces later

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
