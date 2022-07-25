package main

import (
	"fmt"
)

type homes interface {
	property() int
}

type SingleHome struct {
	rooms   int
	cost    int
	intrest float64
}

func (s SingleHome) property() int {
	fmt.Println(s)
	return s.cost
}

type TownHome struct {
	rooms    int
	location string
}

func (t TownHome) property() int {
	fmt.Println(t)
	return t.rooms
}

func realEstate(h homes) {
	fmt.Println("test")
	list := h.property()
	fmt.Println(list)

}

func main() {
	s1 := SingleHome{
		rooms:   4,
		cost:    100000,
		intrest: 5.2,
	}
	t1 := TownHome{
		rooms:    10,
		location: "Ind",
	}
	//var h homes
	//fmt.Println(h.property())
	_ = s1
	//fmt.Println(s1)
	//fmt.Println(t1)
	realEstate(s1)

	realEstate(t1)
}
