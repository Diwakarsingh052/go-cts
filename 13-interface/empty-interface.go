package main

import "fmt"

func main() {

	var i interface{} = "hello"
	i = 21
	i = true
	i = 78.66
	i = "some string"
	fmt.Printf("%T\n", i)

	s, ok := i.(string) // type assertion is used to take the value out of the interface
	if !ok {

		fmt.Println("it is not of the string type")
		return
	}
	fmt.Println(s)

	acceptAnything(10, "hello", true, 10.99)
}

func acceptAnything(a ...interface{}) {
	fmt.Printf("%#v\n", a)
	for _, v := range a {
		fmt.Printf("%T\n", v)
	}

}
