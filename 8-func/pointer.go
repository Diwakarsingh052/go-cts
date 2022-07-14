package main

import "fmt"

func main() {

	a := 10
	var p *int // default is nil
	p = &a     // storing the address of var a in pointer p
	fmt.Println(&a)
	fmt.Println(p) // it would print the address

	*p = 20 // * is a dereference operator to access the memory stored in the pointer
	// any changes made by p would be reflected back to the var a because p points to the memory location of a var
	fmt.Println(a)
	fmt.Println(*p) //*p prints value at this address

}
