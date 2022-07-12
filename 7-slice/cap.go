package main

import "fmt"

func main() {
	// slices don't store anything their own
	// they point to a backing array to store data
	x := []int{10, 20, 40}
	fmt.Printf("before append %p\n", x)
	fmt.Println(len(x), cap(x))

	x = append(x, 70) //if we don't have enough cap than append will allocate a new memory
	fmt.Println("after 1st append", len(x), cap(x))
	fmt.Printf("after 1st append %p\n", x)

	x = append(x, 80)
	fmt.Println("after 2 append", len(x), cap(x))
	fmt.Printf("after 2nd append %p\n", x)

	x = append(x, 90, 100)
	fmt.Println("after 3rd append", len(x), cap(x))
	fmt.Printf("after 3rd append %p\n", x)

	fmt.Println(len(x), cap(x))

}
