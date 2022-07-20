package main

import (
	"fmt"
	"strconv"
)

//var ErrFileNotFound = errors.New("a.txt file not found in the root directory")

func main() {
	// problem statement : we need to achieve variable output for our errors like strconv package

	_, err := strconv.Atoi("efg")
	fmt.Println(err)
	_, err = strconv.Atoi("abc")
	fmt.Println(err)
	_, err = strconv.ParseInt("qwerty", 10, 64)
	fmt.Println(err)
}
