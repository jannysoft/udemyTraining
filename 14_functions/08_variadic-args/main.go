package main

import "fmt"
// declare func main with no receiver, no parameter and no return
func main() {
	// declare var data with value of slice with type of float64
	data := []float64{43, 56, 87, 12, 45, 57}
	// declare var n with value returned from func average
	// we pass variadic args to func average with data... one at time
	n := average(data...)
	// print the value of var n which has been returned from func average
	fmt.Println(n)
}

// declare func average
// func average has no receiver
// func average has variadic parameters of type float64 that are assigned by the call from func main
// func average returns value of type float64 after calculations
func average(sf ...float64) float64 {
	// declare var total with value of zero and type float64
	total := 0.0
	// create loop in which we are not interested to assign anything to its var
	// but only interested of the value it generates
	// in this case the loop goes trough the sf variadic parameter
	// and assigns each value it finds to var v every time the loop goes trough
	for _, v := range sf {
		// then when a value is assigned to var v we add that value to var total
		// in this case
		// 0 + 43 = 43
		// 43 + 56 = 99
		// 99 + 87 = 186
		// 186 + 12 = 198
		// 198 + 45 = 243
		// 243 + 57 = 300 and this is our final value
		total += v
	}
	// return our final value of loop / by the length of func variadic parameter sf
	// in this case 300 / 6 = 50
	// now the returned value to func average of 50 can be passed to the call of func main
	return total / float64(len(sf))
}
