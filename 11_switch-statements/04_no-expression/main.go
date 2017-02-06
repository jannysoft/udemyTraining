package main

import "fmt"

func main() {
	myFriendsName := "Mar"
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("whats up my friend with name only with 2 chars")
	case myFriendsName == "Tim":
		fmt.Println("Whats up Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Whats up Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("your name is eather Marcus or Methi")
	case myFriendsName == "Julian":
		fmt.Println("Whats up Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Whats up Sushant")
	default:
		fmt.Println("Nothing matched this is the default")
	}
}
