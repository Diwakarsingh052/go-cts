package main

import "fmt"

type student struct {
	name, email string
}

func (s student) printData() {
	fmt.Println(s.name)
}

func main() {
	//var i []int
	//i := []int{10, 30, 50}
	//b := []int{60, 70, 20}
	//sum := i[1] + b[1]
	//var s []student
	//s1 := student{
	//	name:  "abc",
	//	email: "abc@email.com",
	//}
	//s = append(s, s1)
	//s = append(s, s1)
	//s = append(s, s1)
	//s = append(s, s1)
	//fmt.Println(s[1])
	////fmt.Printf("%#v", s)
	//s[0].printData()
	//fmt.Println(s[1].email)

	s2 := []student{
		{
			name:  "raj", //s2[0]
			email: "r@email.com",
		},
		{
			name:  "abc",
			email: "abc@email.com",
		},
		{}, // this {} represents one struct object
	}

	s2[0].printData()
}
