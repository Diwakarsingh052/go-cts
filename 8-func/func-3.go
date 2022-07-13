package main

import "fmt"

func main() {
	hey()
}

func hey() {
	//anonymous func
	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20) // () this means we are calling the func

	func(a, b int) {
		fmt.Println(a + b)
	}(10, 20) // () this means we are calling the func

}
