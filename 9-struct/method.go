package main

import "fmt"

type user struct {
	name     string
	email    string
	password string
}

func (x user) printData() { // func (receiver) methodName(param) (return values){}
	fmt.Println(x.name)
	fmt.Println(x)
}

func (x user) anyName() {
	fmt.Println(x.name)
}

func main() {
	u := user{
		name:     "ajay",
		email:    "a@email.com",
		password: "abcd",
	}
	u1 := user{
		name:     "abc",
		email:    "abc@email.com",
		password: "abcd",
	}

	u.printData() // u data would be "copied" to the x receiver/var of the struct
	u.anyName()

	u1.printData()
}
