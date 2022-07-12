package main

import "fmt"

func main() {
	var a [5]int //fixed size
	a[0] = 100   // update the certain index
	a[3] = 200
	b := [4]int{10, 30, 50} //fixed size
	fmt.Println(a, b)

	for i, v := range a { // index,value
		fmt.Println(i, v)
	}

	for _, v := range a { // value
		fmt.Println(v)
	}

	for i := range a { // index
		fmt.Println(i)
	}

}
