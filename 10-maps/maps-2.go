package main

import "fmt"

type user struct {
	name  string
	email string
	marks []int
}

var student = map[int]user{
	100: {
		name:  "ajay",
		email: "ajay@email.com",
		marks: []int{10, 2, 69},
	},
	101: {
		name:  "raj",
		email: "",
		marks: nil,
	},
	102: {
		name:  "abhay",
		email: "",
		marks: nil,
	},
}

func main() {
	//u := user{
	//	name:  "ajay",
	//	email: "ajay@email.com",
	//	marks: []int{10, 2, 69},
	//}

	key := 100
	fmt.Println(student[key])
	searchRecords(101)
}

func searchRecords(id int) {
	u, ok := student[id]
	if !ok {
		fmt.Println("user not found")
		return
	}
	fmt.Println(u)

}
