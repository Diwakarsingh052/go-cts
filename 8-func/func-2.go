package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	sum, err := SumString("12", "20")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sum)
}

func SumString(s, x string) (int, error) { // err must be the last value to be returned
	a, err := strconv.Atoi(s)

	if err != nil {
		return 0, err // it stops the current func and return the values if any
	}
	b, err := strconv.Atoi(x)

	if err != nil {
		return 0, err // when err happens then set other values to their defaults
	}

	sum := a + b
	return sum, nil

}
