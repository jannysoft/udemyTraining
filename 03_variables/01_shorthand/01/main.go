// only main package is executable
package main

// import fmt(format) from builtin library
import "fmt"

// declare variable g is a string
// assign string "cowboy" to variable g
var g string = "cowboy"

// constructing our main function
func main() {
	// a, b, c, d are shorthanded variables
	// shorthanded variables can only be declared in a function

	// a is variable which contains integer 10
	a := 10
	// b is variable which contains string
	b := "golang"
	// c is variable which contains float
	c := 4.17
	// d is variable which contains boolean
	d := true

	// print (initialise) variable a, \n is escape character for new line
	// %v is the value in default format in this case (10 integer)
	fmt.Printf("%v \n", a)
	// print (initialise) variable b, \n is escape character for new line
	// %v is the value in default format in this case ("golang" string)
	fmt.Printf("%v \n", b)
	// print (initialise) variable c, \n is escape character for new line
	// %v is the value in default format in this case (4.17 float64)
	fmt.Printf("%v \n", c)
	// print (initialise) variable d, \n is escape character for new line
	// %v is the value in default format in this case (boolean true)
	fmt.Printf("%v \n", d)
	// print (initialise) variable g, \n is escape character for new line
	// %v is the value in default format in this case ("cowboy" string)
	fmt.Printf("%v \n", g)
}
