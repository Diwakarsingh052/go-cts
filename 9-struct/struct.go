package main

import "fmt"

type student struct { // user is a new data type
	name  string // fields
	age   int
	marks []int
}

func main() {

	s := student{ // s is a var of the student struct
		name:  "ajay",
		age:   29,
		marks: []int{10, 203, 40, 50},
	}
	s.name = "other name" // to access the fields of the struct we need var of the struct [structVar.FieldName]
	s.marks = []int{90, 30, 0}
	fmt.Println(s)
	//fmt.Printf("%+v", s)
	fmt.Printf("%#v\n", s)
	fmt.Printf("%v\n", s.name)

	s1 := student{ // s is a var of the student struct
		name:  "raj",
		age:   29,
		marks: []int{110, 2103, 410, 510},
	}
	fmt.Println(s1.name)

	var s3 student  // setting s3 to the default
	s4 := student{} // setting the s4 to default which means each field would be set to their respective defaults
	fmt.Printf("%#v", s3)
	fmt.Println(s4)

}
