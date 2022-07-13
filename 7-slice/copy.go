package main

import "fmt"

func main() {

	a := []int{900, 670, 610, 950, 870, 200}
	b := make([]int, len(a)) // cap won't matter in copy
	//b := make([]int, 3, 10000) // only three emls would be copied
	fmt.Println(cap(b))
	copy(b, a)
	b[0] = 99
	fmt.Println(a)
	fmt.Println(b)
}
