package main

import "fmt"

func main() {
	i := []int{10, 20, 30, 40, 50, 60, 70}

	b := i[2:5:5] // cap = max-index // to make cap == len
	//b[0] = 999 // this will not save from the side effect of the updating the value before the append
	b = append(b, 200)
	inspectSlice("i", i)
	inspectSlice("b", b)
}

func inspectSlice(name string, slice []int) {
	fmt.Printf("name %v len %d cap %d \n", name, len(slice), cap(slice))
	fmt.Println(slice)
	fmt.Println()
}
