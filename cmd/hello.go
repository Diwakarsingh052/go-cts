package main

import "fmt"

func hello() {
	fmt.Println("this is the hello func")
}

// we can't have two func with the same name inside a package
func main() {

	fmt.Println("this is the main func from the hello.go file")

}
