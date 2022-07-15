package main

import "fmt"

type user struct {
	age int
}

//func (x user) update(age int) {
//	x.age = age
//	fmt.Println("inside the update method", x.age)
//}
func (x *user) updatePtr(age int) {
	x.age = age
	fmt.Println("inside the update methodPtr", x.age)
}

func main() {
	var u user
	//u.update(30) // u data would be copied to method var // it means any changes done in the method would not be reflected back
	u.updatePtr(30)
	// u address instead of the data would be copied
	// which means that x var of update ptr have access to the same memory loc as of u var of the struct any changes made there would be reflected back
	fmt.Println("main", u.age)
}
