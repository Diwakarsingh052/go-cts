package main

import "fmt"

func main() {
	//a := 10
	//x := a // copy
	//x = 100
	//fmt.Println(a, x)
	i := []int{10, 20, 30, 40, 50, 60, 70}
	b := i[1:4]                                // this is not copy it is referencing the same backing array of the base slice
	fmt.Println("len and cap", len(b), cap(b)) // cap is calc for it from the base index where b starts and till the end of the base slice
	fmt.Printf("i= %p\n", i)
	fmt.Printf("b= %p\n", b)
	b[0] = 999 // it will not allocate a new memory so it will change in the base slice as well
	fmt.Println(i)
	fmt.Println(b)
}
