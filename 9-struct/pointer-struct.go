package main

import "fmt"

type user struct {
	name string
}

func main() {

	//var p *int = &a
	var x user
	p := &x // var s *student = &x

	p.name = "ajay" // we don't need * operator to access the field
	x.name = "abc"
	fmt.Println(x)
	fmt.Println(p)

}
