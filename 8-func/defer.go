package main

import "fmt"

func main() {
	//defer maintains a stack. value added first would be exec at last
	defer fmt.Println(1) // when your function is returning defer statements will exec
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	panic("i need panic") // defer guarantee exec // if the statements are registered before the panic or return
	fmt.Println(5)
}
