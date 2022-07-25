package main

import (
	"fmt"
)

type Runner interface {
	Run()
}
type Walker interface {
	Walk()
}

// to implement WalkRunner we need to define all the methods of the embedded interface
type WalkRunner interface {
	Runner
	Walker
}

type Human struct{ name string }

func (h Human) Walk() {
}
func (h Human) Run() {
}

type Robo struct{ name string }

func (r Robo) Run() {}

func main() {
	h := Human{name: "ajay"}
	robo := Robo{name: "r1"}
	_, _ = robo, h
	//var wr WalkRunner = h // we can only assign human struct value as robo doesn't iml all the methods of the WalkRunner interface
	//anything(robo)
	var r Runner = h

	x, ok := r.(Human) // type assertion // check whether a type exists in the interface or not // if it is there than that struct would be returned
	if !ok {           // ok == false
		fmt.Println("r is not storing the instance of the human struct")
		fmt.Println("you are not human I can't hire you")
		return
	}
	fmt.Println("we can use human for some  work", x.name)
}

func anything(wr WalkRunner) {

}
