package main

import (
	"fmt"
	"strconv"
)

func main() {
	day := "fri"
	//if day == "mon" {
	//
	//} else if day == "tues " {
	//
	//} else {
	//
	//}

	switch day {

	case "mon":
		fmt.Println("monday")

	case "tues":
		fmt.Println("tuesday")

	default:
		fmt.Println("nothing matches")

	}

	_, err := strconv.Atoi("abc")

	switch err {
	case nil:
		fmt.Println("no error")
	default:
		fmt.Println("error happened ", err)

	}

}
