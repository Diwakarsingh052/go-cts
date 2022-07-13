package main

import "fmt"

func main() {
	//b := make([]int, 0, 100)
	var i []int
	i = make([]int, 0, 1000) // make(type,len,cap) // initialize a slice with a underlying array of specified size in the cap
	if i == nil {
		fmt.Println("it is nil")
	} else {
		fmt.Println("no it not nil")
	}

	inspectSlice("i", i)
	i = append(i, 100)
	inspectSlice("i", i)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v Length %d Cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
