// only main package is executable
package main

// import fmt(format) from builtin library
import "fmt"

// declare variable g is a string
// assign string "cowboy" to variable g
var g = "cowboy"

// constructing our main function
func main() {
	// a, b, c, d are shorthanded variables
	// shorthanded variables can only be declared in a function

	// a is variable which contains integer 10
	a := 10
	// b is variable which contains string
	b := "golang"
	// c is variable which contains float64
	c := 4.17
	// d is variable which contains boolean
	d := true

	// PrintF (format) (initialise) variable a and show what type is it ?
	// %T displays the type contained in the variable a in this case (integer)
	fmt.Printf("%T \n", a)
	// PrintF (format) (initialise) variable a and show what type is it ?
	// %T displays the type contained in the variable a in this case (string)
	fmt.Printf("%T \n", b)
	// PrintF (format) (initialise) variable a and show what type is it ?
	// %T displays the type contained in the variable a in this case (float64)
	fmt.Printf("%T \n", c)
	// PrintF (format) (initialise) variable a and show what type is it ?
	// %T displays the type contained in the variable a in this case (boolean)
	fmt.Printf("%T \n", d)
	// PrintF (format) (initialise) variable a and show what type is it ?
	// %T displays the type contained in the variable a in this case (string)
	fmt.Printf("%T \n", g)
}
