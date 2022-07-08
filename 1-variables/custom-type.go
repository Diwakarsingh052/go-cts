package main

import (
	"fmt"
)

type money int

func main() {

	//s := "10"
	//x := 10.0
	//z := 10
	var rupee money = 100
	//var o int
	i := 1000
	rupee = money(i)
	fmt.Println(rupee)
	//time.Duration()
	hello()
}

func hello() {
	var rupee money
	rupee = 10
	fmt.Println(rupee)
}
