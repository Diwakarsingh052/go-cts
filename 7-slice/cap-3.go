package main

import (
	"fmt"
)

func main() {
	i := []int{10, 20, 30, 40, 50, 60, 70}
	b := i[2:5]

	//b = append(b, 999, 888) //It is going to change the base slice because they are  sharing the same memory
	//b = append(b, 999, 888, 777) // if we don't have enough capacity it is going to allocate a new memory

	inspectSlice("i", i)
	inspectSlice("b", b)
}

//var name string
//var slice []int
func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v len %d cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
