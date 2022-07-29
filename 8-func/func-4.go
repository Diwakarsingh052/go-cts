package main

import (
	"fmt"
)

type sum func(a, b int)

func main() {

	i := func(a, b int) {
		fmt.Println(a + b)
	}

	i(10, 30)
	calc(i, "hello")
	//18-http.HandleFunc()

}

func calc(total sum, a string) {
	total(10, 20)

}

func calc1(total func(a, b int), a string) {
	total(10, 20)
	fmt.Println(a)

}
