package main

import "fmt"

func main() {

	if a, b := 20, 30; a > b {
		fmt.Println("greater", a)
	} else {
		fmt.Println("greater", b)
	}
	{
		var i int
		fmt.Println(i)
	}
	//fmt.Println(i)
	//fmt.Println(a)
}
