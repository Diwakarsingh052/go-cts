package main

import "fmt"

type user struct {
	name string
	age  int
}

func main() {
	var usr struct{}
	var u user
	checkType("hello")
	checkType(10)
	checkType(90.44)
	checkType(true)
	checkType(usr)
	checkType(u)
}
func checkType(i interface{}) {
	// type switch // it can exec the different cases according to the type stored in the empty interface{}
	switch v := i.(type) {
	case string:
		fmt.Println(v, "it is a string")
	case bool:
		fmt.Println("it is a bool", v)
	case float64:
		fmt.Println("it is a float", v)
	case user:
		fmt.Println("it is a struct", v.age, v.name)
	default:
		fmt.Println("i don't know what is the type here")
	}

}
