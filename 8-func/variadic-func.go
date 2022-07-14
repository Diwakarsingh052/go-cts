package main

import "fmt"

func main() {
	//s := []int{10, 203, 304, 50}
	show("10", 230, 40, 59, 90)
	show("hello")
	//add()
	fmt.Println()
}
func add(a int) {

}

//func show(a []int) {
//
//}
// show have v as a variadic param which means we can pass any number of values to v
func show(s string, v ...int) { // variadic param should be the last value in the function param
	fmt.Printf("%T\n", v)
	fmt.Printf("%#v", v)

	for i, v := range v {
		fmt.Println(i, v)
	}
}
