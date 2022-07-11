package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println(os.Args)
	list := os.Args[1:] // from first index till the end of the list
	fmt.Println(list)

}
