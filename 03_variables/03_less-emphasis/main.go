// only package main is executable therefore is above all other packages
package main
//import fmt (format) package from the builtin library
import "fmt"

// construct function main
func main() {
	// declare variable a is integer and assign 0 value to it (0)
	var a int
	// declare variable b is string and assign 0 value to it("")
	var b string
	// declare variable c is float64 and assign 0 value to it (0)
	var c float64
	// declare variable b is boolean and assign 0 value to it (false)
	var d bool

	// PrintF (print format) the value of a in this case empty integer (0)
	// PrintF gets the value contained in variable a and prints it trough %v
	// PrintF (print format) the value of b in this case empty string ("")
	fmt.Printf("%v \n", a)
	// PrintF gets the value contained in variable a and prints it trough %v
 	// PrintF (print format) the value of c in this case empty float64 (0)
	fmt.Printf("%v \n", b)
	// PrintF gets the value contained in variable a and prints it trough %v
	fmt.Printf("%v \n", c)
	// PrintF (print format) the value of d in this case empty boolean (false)
	// PrintF gets the value contained in variable a and prints it trough %v
	fmt.Printf("%v \n", d)


	// print an empty line with Println()
	fmt.Println()
}