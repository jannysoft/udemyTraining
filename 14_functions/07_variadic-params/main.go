package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}


/*
FUNC MAIN
In func main we are declaring var "n" and assigning the value from the call of func average
then we print that var "n"

FUNC AVERAGE
When func main calls func average, main asks average can I pass the following arguments
and you do the calculation and return the result to me.

Func average says yes I can assign the arguments passed to my by main because
I have identifier "sf" that cen accept multiple parameters passed from main's arguments.

But first I will do my calculations:
1. I will declare variable "total" with zero value and type float64
2. I will do a loop trough my sf parameter and see how many parameters I have in it "range sf"
and then for each number I find in my parameters I will assign it to v.
3. Each time I declare a new number in var "v" I will assign it to var "total" and I will add
the previous number that as assign to var "total"
4. When I am done with my loop I will then get the final value that I have calculated in my loop
which is contained in var "total" and devide it by the number of arguments that has been passed to
my parameters

func average calculation and result:

(43 + 56 + 87 + 12 + 45 + 57) = 300 a value that is contained in total

300 / by the length of parameters in sf "6" = 50

return 50 for a final value in func average and pass it back to main's var n so it can be printed

*/