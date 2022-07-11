package main

import "fmt"

func main() {

	a := 10
	if a > 20 && a < 30 {

	} else if a > 20 {

	} else {

	}

	ok := false

	if !ok { // ok == true // !ok is ok == false
		fmt.Println("this condition is true")
	} else {
		fmt.Println("this condition is false")
	}

}
