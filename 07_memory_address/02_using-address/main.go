package main

import "fmt"

// declare float64 CONST
const metersToYards float64 = 1.09361

func main() {
	// declare VAR with zero value of type FLOAT64
	var meters float64
	// as the user a question
	fmt.Print("Enter meters swam: ")
	// scan what did the user typed in and store it into memory
	fmt.Scan(&meters)
	// declare VAR yards
	// assign to it the returned value of the user input multiplied by the CONST
	yards := meters * metersToYards
	// print results
	fmt.Println(meters, "meters is ", yards, " yards.")
}
