package main

import "fmt"

func main() {
	var i int64
	var s string
	var ok bool
	var f float64
	fmt.Println("hello", i, ok, f)
	fmt.Println(s)
	fmt.Printf("%q\n", s)
	var x int8 = 127
	x += 2 // x = x+2
	fmt.Println(x)

}
