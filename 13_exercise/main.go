package main

import (
	"fmt"
)

//func exe1() {
//	fmt.Println("Hello World")
//}
//
//func exe2() {
//	myName := "Krasmir"
//	fmt.Println("Hello", myName)
//}
//
//func exe3() {
//	var myName string
//	fmt.Print("Please enter your name: ")
//	fmt.Scan(&myName)
//
//	fmt.Println("Hello ", myName)
//
//}
//
//func exe4() {
//	var num1 float64
//	var num2 float64
//
//	fmt.Print("Please enter your first number: ")
//	fmt.Scan(&num1)
//	fmt.Print("Please enter your second number: ")
//	fmt.Scan(&num2)
//
//	outcome := num1 / num2
//
//	fmt.Println(outcome)
//}
//
//func exe5() {
//	for i := 1; i <= 100; i++ {
//		if i%2 == 0 {
//			fmt.Println(i)
//		}
//	}
//}
//
//func exe6() {
//	for i := 0; i <= 100; i++ {
//		if i%3 == 0 && i%5 == 0 {
//			fmt.Println(i, "\t is Fizzbuzz")
//		} else if  i%3 == 0 {
//			fmt.Println(i, "\t is Buzz")
//		} else if i%5 == 0 {
//			println(i, "\t is Buzz")
//		}
//	}
//}

func exe7() {
	counter := 0
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			counter += i
		} else if i%5 == 0 {
			counter += i
		}
	}
	fmt.Println(counter)
}

func main() {
	//exe1()
	//exe2()
	//exe3()
	//exe4()
	//exe5()
	//exe6()
	exe7()
}
