package main

import (
	"github.com/headzoo/surf"
	"fmt"
)

func main() {
	bow := surf.NewBrowser()
	err := bow.Open("http://udemycoupon.discountsglobal.com/")
	if err != nil {
		panic(err)
	}


	s := bow.Find("a.coupon-code-link")

	fmt.Println(s[1])


	fmt.Println(bow.Title())
}

