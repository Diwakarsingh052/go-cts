package main

import "fmt"

func main() {

	a, b := 10, "hello" // creates and assign
	//a = "hey"
	a, ok := 20, false // a gets updated and ok var is created
	a, b = 88, "hey"
	fmt.Println(a, b, ok)
}
