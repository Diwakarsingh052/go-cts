package main

import (
	"fmt"
)

func main() {
	var a []int
	a = append(a, 10, 20, 30, 40)
	a = append(a, 100, 200, 300, 400)
	a = append(a, 10, 20, 30, 40)
	a[2] = 999
	fmt.Println(a)
	fmt.Println(len(a))

}
