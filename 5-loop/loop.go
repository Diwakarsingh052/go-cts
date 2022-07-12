package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {

	}
	x := 0
	for ; x <= 10; x += 2 { // x = x+2

	}
	z := 0
	for z <= 10 {
		//do some work
		//inc
		z++
	}

	ok := true
	a := 1
	for ok {
		if a == 10 {
			fmt.Println(a)
			break // break stop the loop
		}
		a++
	}

}
