package main

import "fmt"

func main() {
	var i []int // default value is nil which means it is not pointing to any backing/underlying array in the memory
	fmt.Println(i)
	fmt.Printf("%#v\n", i)
	if i == nil {
		fmt.Println("default value is nil")
	}

	//i[1] = 100 // update // not for adding // panic that index doesn't exist in the memory

}
