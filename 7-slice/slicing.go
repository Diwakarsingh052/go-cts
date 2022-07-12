package main

import "fmt"

func main() {
	a := []int{10, 20, 30, 40, 50}
	b := a[1:4] // a[index:len]
	b = a[:]    // take the whole slice
	b = a[:3]   // start till the len provided
	b = a[2:]   // from the 2nd index till the end
	fmt.Println(b)
}
