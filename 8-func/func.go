package main

import "fmt"

func main() {
	hello("any name", 48, 888)
	a := sum(10, 20)
	fmt.Println(a)
	ok := checkIfStringIsEmpty("a")
	fmt.Println(ok)
}

func hello(name string, age, marks int) {
	fmt.Println(name, age, marks)

}

func sum(a, b int) int {
	s := a + b
	return s
}

func checkIfStringIsEmpty(s string) bool {
	if s == "" {
		return true
	}

	return false

}
